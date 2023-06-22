package main

import (
	"log"
	"os"

	_ "github.com/lib/pq"

	webServer "github.com/Cadeusept/notes-app"
	"github.com/Cadeusept/notes-app/pkg/handler"
	"github.com/Cadeusept/notes-app/pkg/repository"
	"github.com/Cadeusept/notes-app/pkg/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("error loading .env file: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("error connecting to database: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	webSrv := new(webServer.Server)
	if err := webSrv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error running web server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
