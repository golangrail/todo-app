package app

import (
	handler "github.com/lov3allmy/todo-app/intenral/delivery/http"
	"github.com/lov3allmy/todo-app/intenral/repository"
	"github.com/lov3allmy/todo-app/intenral/server"
	"github.com/lov3allmy/todo-app/intenral/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Run() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")
	return viper.ReadInConfig()
}
