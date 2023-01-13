package main

import (
	"log"

	"github.com/aleks-tim/todo-app"
	handler "github.com/aleks-tim/todo-app/pkg/handlers"
	"github.com/aleks-tim/todo-app/pkg/repository"
	"github.com/aleks-tim/todo-app/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	servises := service.NewService(repos)
	handlers := handler.NewHandler(servises)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRouters()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
