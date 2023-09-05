package pkg

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
	})

	// run load env
	secret := []byte(os.Getenv("SECRETSIGN"))
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
