package router

import (
	"medium/controller"

	"github.com/go-chi/chi/v5"
)

type BlogRouter struct {
	BlogController controller.BlogController
}

func NewBlogRouter(controller controller.BlogController) Router {
	return &BlogRouter{
		BlogController: controller,
	}
}

func (br *BlogRouter) Register(r chi.Router) {
	r.Post("/create", br.BlogController.CreateBlog)
}
