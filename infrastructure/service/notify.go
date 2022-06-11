package service

import (
	res "go-csv/lib/common/res"

	"github.com/labstack/echo/v4"
)

func NotifySlack(c echo.Context) error {
	return res.ResponseSuccess(c, "succes", nil)
}
