package helpers

import "github.com/gin-gonic/gin"

type Response struct{
	Message string
	Data interface{}
}

func ResponseWritter(c *gin.Context, code int, message string) {
	c.JSON(code,Response{
		Message : message,
		Data: "-",
	})
}

func ResponseWritterWithData(c *gin.Context, code int, message string,data interface{}) {
	c.JSON(code,Response{
		Message : message,
		Data: data,
	})
}