package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/medivh13/mnc-test/internal/transport/http/middleware"

	handlers "github.com/medivh13/mnc-test/internal/transport/http"

	"github.com/medivh13/mnc-test/internal/services"

	"github.com/apex/log"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

func main() {

	errChan := make(chan error)

	e := echo.New()
	m := middleware.NewMiddleware()

	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.SetConfigName("config-dev")

	err := viper.ReadInConfig()
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.Use(m.CORS)
	srv := services.NewService()
	handlers.NewHttpHandler(e, srv)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		errChan <- e.Start(":" + viper.GetString("server.port"))
	}()

	e.Logger.Print("Starting ", viper.GetString("appName"))
	err = <-errChan
	log.Error(err.Error())

}
