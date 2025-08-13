package router

import "github.com/go-chi/chi/v5"

type Router interface {
	Register(r chi.Router)
}

func SetupRouter(routes Router) *chi.Mux {
	chiRouter := chi.NewRouter()
	routes.Register(chiRouter)
	return chiRouter
}
