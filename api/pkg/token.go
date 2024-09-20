package pkg

import (
	"fmt"
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

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	KEY := os.Getenv("JWT_SECRET")
	if KEY == "" {
		KEY = "secret"
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("ACCESS_SECRET_KEY"), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, err

}

type UserInfo struct {
	ID       int
	Username string
	Email    string
}

func GetUserInfo(claims jwt.MapClaims) (*UserInfo, error) {
	id := (claims)["ID"].(int)
	username := (claims)["username"].(string)
	email := (claims)["email"].(string)

	if id == 0 || username == "" || email == "" {
		err := fmt.Errorf("invalid claims")
		return nil, err
	}

	return &UserInfo{
		ID:       id,
		Username: username,
		Email:    email,
	}, nil

}
