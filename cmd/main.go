package main

import (
	"log"

	webServer "github.com/Cadeusept/notes-app"
	"github.com/Cadeusept/notes-app/pkg/handler"
	"github.com/Cadeusept/notes-app/pkg/repository"
	"github.com/Cadeusept/notes-app/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	webSrv := new(webServer.Server)
	if err := webSrv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error running web server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("..\\configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
