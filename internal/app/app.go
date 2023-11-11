package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	v1 "github.com/kostylevdev/todo-rest-api/internal/controller/http/v1"
	"github.com/kostylevdev/todo-rest-api/internal/repository"
	server "github.com/kostylevdev/todo-rest-api/internal/server"
	"github.com/kostylevdev/todo-rest-api/internal/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Run(configPath string) {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(configPath); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(&repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}
	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	controller := v1.NewHandler(service)
	var srv server.Server
	go func() {
		if err := srv.Run(viper.GetString("port"), controller.InitRouter()); err != nil {
			logrus.Fatalf("error occured while runnirest server: %s", err.Error())
		}
	}()
	logrus.Info("todo app started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logrus.Info("shutting down server and database")
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

func initConfig(configPath string) error {
	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
