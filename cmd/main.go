package main

import (
	"log"

	tryrest "github.com/kolibri7557/try-rest-api"
	handler "github.com/kolibri7557/try-rest-api/pkg/handler"
	"github.com/kolibri7557/try-rest-api/pkg/repository"
	"github.com/kolibri7557/try-rest-api/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(tryrest.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
