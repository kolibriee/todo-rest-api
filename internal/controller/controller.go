package controller

import (
	v1 "github.com/kostylevdev/todo-rest-api/internal/controller/http/v1"
	"github.com/kostylevdev/todo-rest-api/internal/service"
)

type Controller struct {
	Handler v1.RouterInitializer
}

func NewController(services *service.Service) *Controller {
	return &Controller{
		Handler: v1.NewHandler(services),
	}
}
