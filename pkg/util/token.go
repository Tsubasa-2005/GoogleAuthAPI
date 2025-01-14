package util

import (
	"fmt"

	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/taxio/errors"
)

type UserToken struct {
	ID       int64
	Name     string
	ImageURL string
}

func GenerateToken(sub string, user UserToken) (string, error) {
	key := os.Getenv("SECRET_KEY")

	claims := jwt.MapClaims{
		"sub":     sub,
		"user_id": user.ID,
		"name":    user.Name,
		"image":   user.ImageURL,
		"exp":     time.Now().Add(time.Hour).Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(key))
	if err != nil {
		return "", errors.Wrap(err)
	}

	return token, nil
}

func ParseToken(token string) (*jwt.Token, error) {
	key := os.Getenv("SECRET_KEY")

	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			err := fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			return nil, errors.Wrap(err)
		}
		return []byte(key), nil
	})
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return t, nil
}
