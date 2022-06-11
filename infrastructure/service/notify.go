package service

import (
	"go-csv/domain/models"
	res "go-csv/lib/common/res"
	"go-csv/usecase"
	"go-csv/usecase/definitions"

	"github.com/labstack/echo/v4"
)

var Interactor definitions.Notify

func NotifySlack(c echo.Context) error {
	data := &[]models.Mahasiswa{}
	if err := c.Bind(&data); err != nil {
		return res.ResponseErr(c, 400, err.Error())
	}
	err := usecase.DoNotifySlackSend(models.NotifSlackParam{
		Text: "Hello",
	})
	if err != nil {
		return res.ResponseErr(c, 400, err.Error())
	}

	return res.ResponseSuccess(c, "success", data)
}
