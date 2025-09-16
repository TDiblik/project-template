package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenerateOauthState(oauthProviderName string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "project-template.inc"
	claims["jti"] = uuid.New().String()
	claims["exp"] = time.Now().Add(10 * time.Minute).Unix() // Unix timestamp
	claims["oauth_provider_name"] = oauthProviderName
	return token.SignedString(EnvData.OAUTH_SECRET_BYTES)
}

func IsValidOauthState(state string) bool {
	token, err := jwt.Parse(state, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrInvalidType
		}
		return EnvData.OAUTH_SECRET_BYTES, nil
	})
	if err != nil {
		return false
	}
	return token.Valid
}

func GetOauthProvider(state string) (string, error) {
	token, err := jwt.Parse(state, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrInvalidType
		}
		return EnvData.OAUTH_SECRET_BYTES, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if provider, ok := claims["oauth_provider_name"].(string); ok {
			return provider, nil
		}
		return "", jwt.ErrTokenInvalidClaims
	}
	return "", jwt.ErrTokenInvalidClaims
}
