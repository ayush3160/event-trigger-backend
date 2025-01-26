package http

import (
	"eventtrigger-backend/pkg/services/cronjobs"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func New(r chi.Router, triggersCollection, eventsCollection *mongo.Collection, redisClient *redis.Client, cronSchedulerSvc *cronjobs.CronScheduler, logger *zap.Logger) {

	eventsService := NewEventsService(logger, eventsCollection, redisClient)
	triggerService := NewTriggerService(logger, triggersCollection, cronSchedulerSvc, eventsService)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("login"))
		})
		r.Post("/register", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("register"))
		})
	})

	r.Route("/triggers", func(r chi.Router) {
		r.Post("/test", triggerService.TestTrigger)
		r.Post("/create", triggerService.CreateTrigger)
		r.Post("/delete", triggerService.DeleteTrigger)
		r.Get("/list", triggerService.ListTriggers)
		r.Get("/get/{name}", triggerService.GetTrigger)
	})
}
