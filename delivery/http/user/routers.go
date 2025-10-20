package user

import (
	"log/slog"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog/v3"
)

// NewRouter creates a new router for the user handler.
func NewRouter(user_handler *UserHandler, logger *slog.Logger) *chi.Mux {

	r := chi.NewRouter()
	r.Use(httplog.RequestLogger(logger, nil))
	r.Get("/users/{id}", user_handler.GetUserByID)
	return r
}
