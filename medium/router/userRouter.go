package router

import (
	"medium/controller"

	"github.com/go-chi/chi/v5"
)

type UserRouter struct {
	UserController controller.UserController
}

func NewRouter(_controller controller.UserController) Router {
	return &UserRouter{
		UserController: _controller,
	}
}

func (ur *UserRouter) Register(r chi.Router) {
	r.Post("/signup", ur.UserController.CreateUser)
}
