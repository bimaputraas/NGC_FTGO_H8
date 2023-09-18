package middleware

import (
	"ms-paylater/helper"

	"github.com/labstack/echo/v4"
)

func (a Auth) Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error{
		// parse and check
		tokenString := c.Request().Header.Get("Authorization")
		id,err := helper.ParseJWT(tokenString)
		if err != nil {
			return helper.WriteErrorResponse(401,err.Error())
		}
		if tokenString == "" {
			return helper.WriteErrorResponse(401,"token does not exist")
		}

		// get user
		user,err := a.Repository.GetUserById(int(id))
		if err != nil {
			return helper.WriteErrorResponse(401,"undefined user")
		}
		c.Set("user",user)
		return next(c)
	}
}