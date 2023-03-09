package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/Cheyzie/golang-test/internal/handler"
	"github.com/Cheyzie/golang-test/internal/repository"
	"github.com/Cheyzie/golang-test/internal/server"
	"github.com/Cheyzie/golang-test/internal/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("error in config loading: %s", err)
	}

	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatalf("error in .env file loading: %s", err)
	}

	db, err := repository.NewPostgresDB(repository.SqlConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("error occured on db connection: %s", err)
	}

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	srv := new(server.Server)
	logrus.Info("server is starting...")
	go func() {
		if err := srv.Run(viper.GetString("server.port"), handler.InitRoutes()); err != nil {
			logrus.Fatalf("error caused on server startup: %s", err)
		}
	}()

	logrus.Info("server is started...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Info("server shutting down...")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Fatalf("error caused on server shutting down: %s", err)
	}

	if err := db.Close(); err != nil {
		logrus.Fatalf("error caused on db disconecting: %s", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("dev")
	return viper.ReadInConfig()
}
