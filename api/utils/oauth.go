package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type RedirectAfterOauth string

// when adding redirect url after oauth, add it here
const (
	RedirectAfterOauthIndex    RedirectAfterOauth = "index"
	RedirectAfterOauthSettings RedirectAfterOauth = "settings"
)

// when adding redirect url after oauth, add it here
var validRedirectAfterOauth = map[RedirectAfterOauth]struct{}{
	RedirectAfterOauthIndex:    {},
	RedirectAfterOauthSettings: {},
}

func ValidateRedirectAfterOauth(value string) RedirectAfterOauth {
	r := RedirectAfterOauth(value)
	if _, ok := validRedirectAfterOauth[r]; ok {
		return r
	}
	return RedirectAfterOauthIndex
}

func (e RedirectAfterOauth) EnumValues() []any {
	return []any{RedirectAfterOauthIndex, RedirectAfterOauthSettings}
}

func GenerateOauthState(oauthProviderName string, redirectAfterOauth RedirectAfterOauth) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "project-template.inc"
	claims["jti"] = uuid.New().String()
	claims["exp"] = time.Now().Add(10 * time.Minute).Unix()
	claims["oauth_provider_name"] = oauthProviderName
	claims["redirect_back_to_after_oauth"] = string(redirectAfterOauth)
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

func GetOauthProviderAndRedirectFromOauthState(state string) (string, RedirectAfterOauth, error) {
	token, err := jwt.Parse(state, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrInvalidType
		}
		return EnvData.OAUTH_SECRET_BYTES, nil
	})
	if err != nil {
		return "", RedirectAfterOauthIndex, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		provider, ok := claims["oauth_provider_name"].(string)
		if !ok {
			return "", RedirectAfterOauthIndex, jwt.ErrTokenInvalidClaims
		}

		redirectStr, _ := claims["redirect_back_to_after_oauth"].(string)
		redirect := ValidateRedirectAfterOauth(redirectStr)

		return provider, redirect, nil
	}

	return "", RedirectAfterOauthIndex, jwt.ErrTokenInvalidClaims
}
