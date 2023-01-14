package main

import (
	// "log"
	"os"

	"github.com/aleks-tim/todo-app"
	handler "github.com/aleks-tim/todo-app/pkg/handlers"
	"github.com/aleks-tim/todo-app/pkg/repository"
	"github.com/aleks-tim/todo-app/pkg/service"
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
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	cfg := repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	db, err := repository.NewPostgresDB(cfg)
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)

	srv := new(todo.Server)
	port := viper.GetString("tcp.port")
	if err := srv.Run(port, handlers.InitRouters()); err != nil {
		logrus.Fatalf("Error occured while running http server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
