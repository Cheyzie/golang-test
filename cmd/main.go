package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Cheyzie/golang-test/internal/server"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("error in config loading: %s", err)
	}

	handler := gin.Default()

	handler.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello world",
		})
	})

	srv := new(server.Server)

	go func() {
		if err := srv.Run(viper.GetString("server.port"), handler); err != nil {
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
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("dev")
	return viper.ReadInConfig()
}
