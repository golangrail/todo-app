package app

import (
	"github.com/lov3allmy/todo-app/intenral/server"
	"github.com/sirupsen/logrus"
)

func Run() {
	srv := new(server.Server)
	if err := srv.Run("8000"); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}
