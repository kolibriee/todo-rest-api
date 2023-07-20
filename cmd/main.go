package main

import (
	"log"

	tryrest "github.com/kolibri7557/try-rest-api"
	handler "github.com/kolibri7557/try-rest-api/pkg/handler"
	"github.com/kolibri7557/try-rest-api/pkg/repository"
	"github.com/kolibri7557/try-rest-api/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(tryrest.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}
