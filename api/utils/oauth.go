package utils

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/gofiber/fiber/v3"
)

func GenerateAndSetOauthStateCookie(c fiber.Ctx, cookie_name string, cookie_allowed_path string) (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	state := hex.EncodeToString(b)
	c.ClearCookie(cookie_name)
	c.Cookie(&fiber.Cookie{
		Name:     cookie_name,
		Value:    state,
		HTTPOnly: true,                // Prevent JS access
		Secure:   !EnvData.Debug,      // Only over HTTPS
		SameSite: "Lax",               // Prevent CSRF
		Path:     cookie_allowed_path, // Accessible across site
		MaxAge:   600,                 // 10 minutes expiry
	})
	return state, nil
}

func GetOauthStateCookie(c fiber.Ctx, cookie_name string) string {
	return c.Cookies(cookie_name)
}
