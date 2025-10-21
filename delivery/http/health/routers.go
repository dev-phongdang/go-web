package health

import (
	"net/http"
	"web-api/delivery/http/response"

	"github.com/go-chi/chi/v5"
)

// NewRouter creates a new router for the user handler.
func NewRouter() *chi.Mux {

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		response.Success(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	return r
}
