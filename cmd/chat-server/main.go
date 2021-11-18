package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"github.com/vitalik-ez/Chat-Golang/pkg/handler"
	"github.com/vitalik-ez/Chat-Golang/pkg/repository"
	"github.com/vitalik-ez/Chat-Golang/pkg/service"
	"github.com/vitalik-ez/Chat-Golang/server"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		log.Fatalf("fatal to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)

	hub := handler.NewHub()

	handlers := handler.NewHandler(services)

	srv := new(server.Server)

	go hub.Run()

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes(hub)); err != nil {
		log.Fatalf("Error occured while running http server", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()

}
