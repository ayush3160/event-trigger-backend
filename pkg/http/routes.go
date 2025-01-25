package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func New(r chi.Router, eventsCollection *mongo.Collection, logger *zap.Logger) {

	eventService := NewEventService(logger, eventsCollection)

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

	r.Route("/events", func(r chi.Router) {
		r.Post("/create", eventService.CreateEvent)
	})
}
