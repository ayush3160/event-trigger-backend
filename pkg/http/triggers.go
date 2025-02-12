package http

import (
	"context"
	"encoding/json"
	"eventtrigger-backend/pkg/models"
	"eventtrigger-backend/pkg/services/cronjobs"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2/bson"
)

type TriggerService struct {
	logger              *zap.Logger
	triggersCollection  *mongo.Collection
	cronJobSchedulerSvc *cronjobs.CronScheduler
	eventsSvc           *EventsService
}

func NewTriggerService(logger *zap.Logger, triggersCollection *mongo.Collection, cronJobSchedulerSvc *cronjobs.CronScheduler, eventsSvc *EventsService) *TriggerService {
	return &TriggerService{
		logger:              logger,
		triggersCollection:  triggersCollection,
		cronJobSchedulerSvc: cronJobSchedulerSvc,
		eventsSvc:           eventsSvc,
	}
}

func (es *TriggerService) handleError(w http.ResponseWriter, err error, message string, statusCode int) {
	es.logger.Error(message, zap.Error(err))
	http.Error(w, err.Error(), statusCode)
}

func (es *TriggerService) TestTrigger(w http.ResponseWriter, r *http.Request) {
	var trigger models.Trigger
	if err := json.NewDecoder(r.Body).Decode(&trigger); err != nil {
		es.handleError(w, err, "Invalid trigger payload", http.StatusBadRequest)
		return
	}

	callBackFunc, err := es.createHttpJob(trigger)
	if err != nil {
		es.handleError(w, err, "Error creating job", http.StatusInternalServerError)
		return
	}

	es.cronJobSchedulerSvc.AddJob(trigger.Name, trigger.Schedule, func() {
		callBackFunc()
		es.cronJobSchedulerSvc.RemoveJob(trigger.Name)
	})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Trigger registered for testing"))
}

func (es *TriggerService) CreateTrigger(w http.ResponseWriter, r *http.Request) {
	var trigger models.Trigger
	if err := json.NewDecoder(r.Body).Decode(&trigger); err != nil {
		es.handleError(w, err, "Invalid trigger payload", http.StatusBadRequest)
		return
	}

	if _, err := es.triggersCollection.InsertOne(r.Context(), trigger); err != nil {
		es.handleError(w, err, "Error inserting trigger into database", http.StatusInternalServerError)
		return
	}

	callBackFunc, err := es.createHttpJob(trigger)
	if err != nil {
		es.handleError(w, err, "Error creating job", http.StatusInternalServerError)
		return
	}

	es.cronJobSchedulerSvc.AddJob(trigger.Name, trigger.Schedule, callBackFunc)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Trigger created"))
}

func (es *TriggerService) DeleteTrigger(w http.ResponseWriter, r *http.Request) {
	eventName := chi.URLParam(r, "name")
	if _, err := es.triggersCollection.DeleteOne(r.Context(), bson.M{"name": eventName}); err != nil {
		es.handleError(w, err, "Error deleting trigger", http.StatusInternalServerError)
		return
	}

	es.cronJobSchedulerSvc.RemoveJob(eventName)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Trigger deleted"))
}

func (es *TriggerService) ListTriggers(w http.ResponseWriter, r *http.Request) {
	var triggers []models.Trigger
	cursor, err := es.triggersCollection.Find(r.Context(), bson.M{})
	if err != nil {
		es.handleError(w, err, "Error listing triggers", http.StatusInternalServerError)
		return
	}

defer cursor.Close(r.Context())

	if err := cursor.All(r.Context(), &triggers); err != nil {
		es.handleError(w, err, "Error parsing triggers", http.StatusInternalServerError)
		return
	}

	es.writeJSONResponse(w, triggers)
}

func (es *TriggerService) GetTrigger(w http.ResponseWriter, r *http.Request) {
	eventName := chi.URLParam(r, "name")
	var trigger models.Trigger
	if err := es.triggersCollection.FindOne(r.Context(), bson.M{"name": eventName}).Decode(&trigger); err != nil {
		es.handleError(w, err, "Error fetching trigger", http.StatusInternalServerError)
		return
	}

	es.writeJSONResponse(w, trigger)
}

func (es *TriggerService) createHttpJob(trigger models.Trigger) (func(), error) {
	parsedURL, err := url.Parse(trigger.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %v", err)
	}

	return func() {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		req, err := http.NewRequestWithContext(ctx, trigger.MethodType, parsedURL.String(), nil)
		if err != nil {
			es.logger.Error("Failed to create request", zap.Error(err), zap.String("endpoint", trigger.Endpoint))
			return
		}

		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			es.logger.Error("Error executing scheduled request", zap.Error(err), zap.String("endpoint", trigger.Endpoint))
			return
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		es.logger.Debug("Scheduled job response", zap.String("endpoint", trigger.Endpoint), zap.Int("status_code", resp.StatusCode), zap.String("response_body", string(body)))
	}, nil
}

func (es *TriggerService) writeJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		es.logger.Error("Failed to encode JSON response", zap.Error(err))
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
