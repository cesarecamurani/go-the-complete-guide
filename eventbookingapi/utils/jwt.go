package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const secretKey = "a1b1f17c-861b-4421-933d-011647bd961e"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}
