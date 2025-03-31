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
	acessToken := jwt.New(jwt.SigningMethodHS256)
	claims := acessToken.Claims.(jwt.MapClaims)
	claims["issuer"] = os.Getenv("JWT_ISSUER")
	claims["auth"] = userId
	claims["type"] = string(AccessToken)
	claims["exp"] = time.Now().Add(24 * time.Hour * 3).Unix()

	acessTokenString, err := acessToken.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}
	return acessTokenString, nil
}

func generateRefreshToken(userId uint) (string, error) {
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	claims := refreshToken.Claims.(jwt.MapClaims)
	claims["issuer"] = os.Getenv("JWT_ISSUER")
	claims["auth"] = userId
	claims["type"] = string(RefreshToken)
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	refreshTokenString, err := refreshToken.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}
	return refreshTokenString, nil
}
