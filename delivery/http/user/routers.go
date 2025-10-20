package user

import (
	"github.com/go-chi/chi/v5"
)

// NewRouter creates a new router for the user handler.
func NewRouter(handler *UserHandler) *chi.Mux {

	r := chi.NewRouter()
	r.Get("/{id}", handler.GetUserByID)

	return r
}
