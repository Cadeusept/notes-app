package main

import (
	"log"

	webServer "github.com/Cadeusept/notes-app"
	"github.com/Cadeusept/notes-app/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	webSrv := new(webServer.Server)
	if err := webSrv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error running web server: %s", err.Error())
	}
}
