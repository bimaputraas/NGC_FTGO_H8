package helper

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(id int) (string,error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
	})

	// Sign and get the complete encoded token as a string using the secret
	sign := []byte(os.Getenv("JWTSIGN"))
	tokenString, err := token.SignedString(sign)
	if err != nil {
		return "",err
	}
	return tokenString,nil
}