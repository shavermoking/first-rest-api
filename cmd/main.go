package main

import (
	todo "first-rest-api"
	"first-rest-api/pkg/handler"
	"first-rest-api/pkg/repository"
	"first-rest-api/pkg/service"
	"log"

	"github.com/spf13/viper"
)

func main() {
	if err := initConfigs(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfigs() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
