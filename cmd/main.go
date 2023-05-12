package main

import (
	practice "github.com/Asstrahanec/PracticeManagementAPI"
	"github.com/Asstrahanec/PracticeManagementAPI/pkg/handler"
	"github.com/Asstrahanec/PracticeManagementAPI/pkg/repository"
	"github.com/Asstrahanec/PracticeManagementAPI/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initialising configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(practice.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running server %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
