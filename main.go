package main

import (
	"go-csv/routes"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SlackConfig struct {
	SlackWebhook string `env:"SLACKWEBHOOK"`
}

func main() {
	e := routes.Init()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hi, I'm Fanzru!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
