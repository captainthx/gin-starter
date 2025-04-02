package service

import (
	"gin-starter/core/dto"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenType string

const (
	AccessToken  TokenType = "Access"
	RefreshToken TokenType = "Refresh"
)

func GenerateToken(userId uint) *dto.TokenResponse {
	accessToekn, err := generateAccessToken(userId)
	if err != nil {
		return nil
	}
	refreshToken, err := generateRefreshToken(userId)
	if err != nil {
		return nil
	}
	return &dto.TokenResponse{
		AccessToken:  accessToekn,
		RefreshToken: refreshToken,
	}
}

func generateAccessToken(userId uint) (string, error) {
	tokenBuilder := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"issuer": os.Getenv("JWT_ISSUER"),
		"auth":   userId,
		"type":   string(AccessToken),
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	acessTokenString, err := tokenBuilder.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return acessTokenString, nil
}

func generateRefreshToken(userId uint) (string, error) {
	tokenBuilder := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"issuer": os.Getenv("JWT_ISSUER"),
		"auth":   userId,
		"type":   string(RefreshToken),
		"exp":    time.Now().Add(24 * time.Hour * 3).Unix(),
	})

	refreshTokenString, err := tokenBuilder.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}
	return refreshTokenString, nil
}
