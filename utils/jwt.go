package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecretkey"

func GenerateToken(userId int64, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (map[string]any, error) {
	var user = make(map[string]any)

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token")
	}

	email := claims["email"].(string)
	userIdFloat, ok := claims["userId"].(float64)
	if !ok {
		return nil, errors.New("invalid user id in token")
	}
	userId := int64(userIdFloat)

	user["email"] = email
	user["userId"] = userId

	return user, nil
}
