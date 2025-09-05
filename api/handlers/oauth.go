package handlers

import (
	"context"
	"fmt"

	"github.com/TDiblik/project-template/api/utils"
	"github.com/gofiber/fiber/v3"
)

var (
	oauthStateStringGithub = "randomstatestring"
)

func GithubRedirect(c fiber.Ctx) error {
	return c.Redirect().To(utils.EnvData.OAUTH_CONFIG_GITHUB.AuthCodeURL(oauthStateStringGithub))
}

func GithubReturn(c fiber.Ctx) error {
	state := c.Query("state")
	if state != oauthStateStringGithub {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid OAuth state")
	}

	code := c.Query("code")
	token, err := utils.EnvData.OAUTH_CONFIG_GITHUB.Exchange(context.Background(), code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to exchange token")
	}

	// Optional: fetch user info from GitHub
	client := utils.EnvData.OAUTH_CONFIG_GITHUB.Client(context.Background(), token)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get user info")
	}
	defer resp.Body.Close()

	return c.SendString(fmt.Sprintf("GitHub user info response status: %s", resp.Status))
}
