package jwt_utils

import (
	"os"

	"github.com/asutosh29/amx-restro/pkg/types"
	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	types.User
	jwt.StandardClaims
}

func GenerateJWT(payload types.User) (string, error) {
	claims := &Claims{
		payload,
		jwt.StandardClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateJWT(token string) (*Claims, error) {
	claims := &Claims{}
	jwt_token, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, err
		}
		return nil, err
	}

	if !jwt_token.Valid {
		return nil, err
	}

	return claims, nil
}
