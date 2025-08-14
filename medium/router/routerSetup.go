package router

import "github.com/go-chi/chi/v5"

type Router interface {
	Register(r chi.Router)
}

func SetupRouter(routes, blogRouter Router) *chi.Mux {
	chiRouter := chi.NewRouter()

	chiRouter.Route("/api/users", func(r chi.Router) {
		routes.Register(r)
	})

	chiRouter.Route("/api/blog/", func(r chi.Router) {
		blogRouter.Register(r)
	})
	return chiRouter
}
