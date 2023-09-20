package dto

import (
	"github.com/labstack/echo/v4"
)

type Response struct {
	Message string
	StatusCode	int
	Data    interface{}
}

func SuccessResponse(c echo.Context,code int,message string) {
	c.JSON(code,Response{
		Message: message,
		StatusCode: code,
		Data: "-",
	})
}

func SuccessResponseResponseWithData(c echo.Context,code int,message string,data interface{}) {
	c.JSON(code,Response{
		Message: message,
		StatusCode: code,
		Data: data,
	})
}

func ErrorResponse(code int,message string) *echo.HTTPError{
	err := echo.NewHTTPError(code,Response{
		Message: message,
		StatusCode: code,
		Data: "-",
	})
	return err
}