package pkg

import (
	"os"
	"time"

	"github.com/ei-sugimoto/tatekae/api/model"
	"github.com/golang-jwt/jwt/v5"
)

func NewToken(user *model.User) (string, error) {
	KEY := os.Getenv("JWT_SECRET")
	if KEY == "" {
		KEY = "secret"
	}

	claims := jwt.MapClaims{
		"username": user.Username,
		"email":    user.Email,
		"ID":       user.ID,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // トークンの有効期限を24時間後に設定
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(KEY))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
