package http

import (
	"encoding/json"
	models "eventtrigger-backend/pkg/models"
	"io"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type EventService struct {
	logger           *zap.Logger
	eventsCollection *mongo.Collection
}

func NewEventService(logger *zap.Logger, eventsCollection *mongo.Collection) *EventService {
	return &EventService{
		logger:           logger,
		eventsCollection: eventsCollection,
	}
}

func (es *EventService) CreateEvent(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		es.logger.Error("error reading body", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var trigger models.Event

	err = json.Unmarshal(body, &trigger)
	if err != nil {
		es.logger.Error("error unmarshalling JSON", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = es.eventsCollection.InsertOne(r.Context(), trigger)

	if err != nil {
		es.logger.Error("error creating event", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("event created"))
}
