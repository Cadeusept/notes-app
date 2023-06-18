package main

import (
	"log"

	webServer "github.com/Cadeusept/notes-app"
	"github.com/Cadeusept/notes-app/pkg/handler"
	"github.com/Cadeusept/notes-app/pkg/repository"
	"github.com/Cadeusept/notes-app/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	webSrv := new(webServer.Server)
	if err := webSrv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error running web server: %s", err.Error())
	}
}
