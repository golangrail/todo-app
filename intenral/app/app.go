package app

import (
	handler "github.com/lov3allmy/todo-app/intenral/delivery/http"
	"github.com/lov3allmy/todo-app/intenral/server"
	"github.com/sirupsen/logrus"
)

func Run() {
	handlers := new(handler.Handler)
	srv := new(server.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}
