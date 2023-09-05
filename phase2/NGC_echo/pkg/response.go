package pkg

import "github.com/labstack/echo/v4"

type ResponseJSON struct {
	Message string
	Data    interface{}
}

func WriteResponse(c echo.Context, code int, message string) {
	c.JSON(code, ResponseJSON{
		Message: message,
		Data:    "-",
	})
}

func WriteResponseWithData(c echo.Context, code int, message string, data interface{}) {
	c.JSON(code, ResponseJSON{
		Message: message,
		Data:    data,
	})
}
