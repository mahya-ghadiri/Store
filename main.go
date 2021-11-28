package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"math/rand"
	"os"
	"os/signal"
	"store/api"
	"store/internal/config"
	"store/internal/redis"
	"syscall"
	"time"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())

	config.Init("config.yaml")
	rand.Seed(time.Now().UnixNano())

	productClient, orderClient := redis.Connect(config.Cfg.Redis)
	redis.Init(productClient, orderClient)

	api.RegisterRoutes(e)
	go func() {
		err := e.Start(fmt.Sprintf(":%s", config.Cfg.App.Port))
		if err != nil {
			logrus.Fatal(err)
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-signalChan
}
