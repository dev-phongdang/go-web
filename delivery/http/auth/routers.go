package auth

import (
	"github.com/go-chi/chi/v5"
)

// NewRouter creates a new router for the user handler.
func NewRouter(handler *AuthHandler) *chi.Mux {

	r := chi.NewRouter()
	r.Post("/login", handler.Login)

	return r
}
