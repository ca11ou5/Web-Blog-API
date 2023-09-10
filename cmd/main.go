package main

import (
	"blog/pkg/handler"
	"blog/pkg/repository"
	"blog/pkg/service"
	"blog/server"
	"github.com/spf13/viper"
	"log"
)

func main() {
	err := initConfig()
	if err != nil {
		log.Fatal(err)
	}

	db := repository.ConnectToDB(repository.Config{
		Host:     viper.GetString("db.host"),
		User:     viper.GetString("db.user"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		Port:     viper.GetString("db.port"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	err = srv.Run(viper.GetString("server.port"), handlers.InitRoutes())
	if err != nil {
		log.Fatal(err)
	}
}

func initConfig() error {
	viper.SetConfigFile("configs/main.yaml")
	return viper.ReadInConfig()
}
