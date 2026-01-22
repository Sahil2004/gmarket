package utils

import (
	"time"

	"github.com/Sahil2004/gmarket/server/config"
	"github.com/Sahil2004/gmarket/server/dtos"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(user dtos.UserDTO) (string, error) {
	claims := jwt.MapClaims{
		"user": user,
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(config.AppConfig().AccessSecret))

	if err != nil {
		return "", err
	}
	
	return t, nil
}

func GenerateRefreshToken(user dtos.UserDTO) (string, error) {
	claims := jwt.MapClaims{
		"user": user,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(config.AppConfig().RefreshSecret))

	if err != nil {
		return "", err
	}

	return t, nil
}