package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	webServer "github.com/Cadeusept/notes-app"
	"github.com/Cadeusept/notes-app/pkg/handler"
	"github.com/Cadeusept/notes-app/pkg/repository"
	"github.com/Cadeusept/notes-app/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load("../.env"); err != nil {
		logrus.Fatalf("error loading .env file: %s", err.Error())
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
		logrus.Fatalf("error connecting to database: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	webSrv := new(webServer.Server)
	go func() {
		if err := webSrv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("error running web server: %s", err.Error())
		}
	}()

	logrus.Print("Notes app started")

	c_quit := make(chan os.Signal, 1)
	signal.Notify(c_quit, syscall.SIGTERM, syscall.SIGINT)
	sig := <-c_quit

	logrus.Printf("catched signal: %s. Notes app shutting down...", sig.String())

	if err := webSrv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured during shutdown: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("error occured during database connection closure: %s", err.Error())
	}

	logrus.Print("Notes app shut down successfully")
}

func initConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
