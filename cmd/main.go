package main

import (
	"context"
	"github.com/gznrf/todo-app"
	"github.com/gznrf/todo-app/pkg/handler"
	"github.com/gznrf/todo-app/pkg/repository"
	"github.com/gznrf/todo-app/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("Ошибка в чтении конфига %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("ошибка загрузки переменных из окружения %s", err.Error())
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
		logrus.Fatalf("Ошибка подклчения к базе данных %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Ошибка запуска сервера %s", err.Error())
		}
	}()

	logrus.Print("app started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("app shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error on server shuting down: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("error on db connection closing: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
