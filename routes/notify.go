package routes

import (
	"go-csv/infrastructure/service"

	"github.com/labstack/echo/v4"
)

func RouteNotify(e *echo.Echo) *echo.Echo {
	e.POST("/api/notify", service.NotifySlack)
	return e
}
