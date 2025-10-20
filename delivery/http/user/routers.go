package user

import "github.com/go-chi/chi/v5"

// NewRouter creates a new router for the user handler.
func NewRouter(user_handler *UserHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/users/{id}", user_handler.GetUserByID)
	return r
}
