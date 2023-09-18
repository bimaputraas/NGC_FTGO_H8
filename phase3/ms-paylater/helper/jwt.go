package helper

import (
	"errors"
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

func ParseJWT(tokenString string) (float64,error){

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		sign := []byte(os.Getenv("JWTSIGN"))
		return sign, nil
	})
	if err != nil {
		return 0,errors.New("invalid token")
	}

	// success
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := claims["id"].(float64)
		return id,nil
	}
	
	// fail
	return 0,errors.New("invalid token")	
}