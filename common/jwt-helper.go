package common

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("secret") // Replace with your secret key

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(email string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}

func VerifyToken(tknStr string) (*Claims, error) {
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		fmt.Printf("Error parsing token: %v\n", err)
		if err == jwt.ErrSignatureInvalid {
			return nil, err
		}
		return nil, err
	}

	if !tkn.Valid {
		fmt.Println("Token is invalid")
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
