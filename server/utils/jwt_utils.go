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

func GenerateRefreshToken(userId string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(config.AppConfig().RefreshSecret))

	if err != nil {
		return "", err
	}

	return t, nil
}

func ValidateTokens(accessToken string, refreshToken string) (*jwt.Token, *jwt.Token, error) {
	refreshTokenParsed, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig().RefreshSecret), nil
	})
	
	if err != nil {
		return nil, nil, err
	}

	accessTokenParsed, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig().AccessSecret), nil
	})

	if err != nil {
		return nil, refreshTokenParsed, err
	}

	return accessTokenParsed, refreshTokenParsed, nil
}