package res

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type SuccessResponse struct {
	Code    int         `json:"code"`
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func ResponseErr(c echo.Context, code int, message string) error {
	res := ErrorResponse{
		Code:    code,
		Status:  false,
		Message: message,
	}
	return c.JSON(http.StatusBadRequest, res)
}
func ResponseSuccess(c echo.Context, message string, data interface{}) error {
	res := SuccessResponse{
		Code:    200,
		Status:  true,
		Message: message,
		Data:    data,
	}
	return c.JSON(http.StatusOK, res)
}
