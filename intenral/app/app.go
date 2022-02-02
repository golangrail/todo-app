package app

import (
	"github.com/joho/godotenv"
	// Postgres driver
	_ "github.com/lib/pq"
	handler "github.com/lov3allmy/todo-app/intenral/delivery/http"
	"github.com/lov3allmy/todo-app/intenral/repository"
	"github.com/lov3allmy/todo-app/intenral/server"
	"github.com/lov3allmy/todo-app/intenral/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func Run() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
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
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
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
