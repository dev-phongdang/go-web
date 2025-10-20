package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Routers struct {
	Root *chi.Mux
	Sub  *chi.Mux
}

func NewRouters(pattern string, r *chi.Mux) *Routers {
	root := chi.NewRouter()
	root.Use(middleware.RequestID)
	root.Use(middleware.RealIP)
	root.Use(middleware.Logger)
	root.Use(middleware.Recoverer)
	root.Mount(pattern, r)
	return &Routers{Root: root, Sub: r}
}
func (r *Routers) Register(pattern string, sub *chi.Mux) error {
	r.Sub.Mount(pattern, sub)
	return nil
}
