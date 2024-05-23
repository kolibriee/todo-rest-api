package controller

import (
	"net/http"

	v1 "github.com/kostylevdev/todo-rest-api/internal/controller/http/v1"
	"github.com/kostylevdev/todo-rest-api/internal/service"
)

type RouterInitializer interface {
	InitRouter() http.Handler
}

type Controller struct {
	RouterInitializer
}

func NewController(services *service.Service) *Controller {
	return &Controller{
		RouterInitializer: v1.NewHandler(services),
	}
}
