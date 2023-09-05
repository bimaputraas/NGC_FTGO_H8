package middleware

import (
	"ngc_echo/config"
	"ngc_echo/helpers"
	"ngc_echo/models"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func Authentication(next echo.HandlerFunc) echo.HandlerFunc{
	return func(c echo.Context) error {
		// cookie,err := c.Cookie("Authorize-Token")
		// if err != nil {
		// 	helpers.WriteResponseWithData(c,401,"Failed get cookie",cookie)
   	 	// 	return nil
		// }
		// tokenString := cookie.Value

		// get from request header
		tokenString:= c.Request().Header.Get("Authorization")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {	
			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			secret := []byte(os.Getenv("SECRETSIGN"))
			return secret, nil
		})

		if err != nil {
			helpers.WriteResponse(c,401,"Unauthorized")
			return nil
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// token claims
			id := claims["id"]

			// check user from db
			var user models.Users
			db := config.InitDatabase()
			result := db.Where("id = ?",id).First(&user)
			// if not exist
			if result.Error != nil {
				helpers.WriteResponse(c,401,"Data does not exist")
   	 			return nil
			}

			// if exist then set user to context
			c.Set("user",user)
			return next(c)
		}
		helpers.WriteResponse(c,401,"Unauthorized")
		return nil
	}
}