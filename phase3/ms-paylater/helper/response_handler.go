package helper

import (
	"github.com/labstack/echo/v4"
)

type Response struct {
	Message string
	Data    interface{}
}

type ErrResponse struct {
	Message    string
	StatusCode int `json:"status_code"`
}

func WriteResponse(c echo.Context,code int, message string) error{
	err := c.JSON(code,Response{
		Message: message,
		Data: "-",
	})
	if err != nil {
		return nil
	}
	return err
}

func WriteResponseWithData(c echo.Context,code int, message string,data interface{}) error {
	err := c.JSON(code,Response{
		Message: message,
		Data: data,
	})
	if err != nil {
		return nil
	}
	return err
}

func WriteErrorResponse(code int, message string) error{
	err := echo.NewHTTPError(code,ErrResponse{
		Message: message,
		StatusCode: code,
	})

	return err
}