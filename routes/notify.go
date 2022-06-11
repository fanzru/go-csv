package routes

import (
	"go-csv/handlers"

	"github.com/labstack/echo/v4"
)

func RouteNotify(e *echo.Echo) *echo.Echo {
	e.POST("/api/notify", handlers.NotifySlack)
	return e
}
