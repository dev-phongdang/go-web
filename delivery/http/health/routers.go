package health

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// NewRouter creates a new router for the user handler.
func NewRouter() *chi.Mux {

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	return r
}
