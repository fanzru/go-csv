package handlers

import (
	res "go-csv/common/res"

	"github.com/labstack/echo/v4"
)

func NotifySlack(c echo.Context) error {
	return res.ResponseSuccess(c, 200, true, "succes", nil)
}
