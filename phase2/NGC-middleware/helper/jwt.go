package helper

import (
	"errors"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(claims jwt.MapClaims) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(os.Getenv("JWTSIGNATURE")))

	return token,err
}

func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWTSIGNATURE")), nil
	})

	if err != nil {
		fmt.Println(err.Error())
		return jwt.MapClaims{}, errors.New("failed to decode token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return jwt.MapClaims{}, errors.New("failed to decode token")
	}
}