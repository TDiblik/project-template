package middleware

import (
	"errors"

	"github.com/TDiblik/project-template/api/utils"
	"github.com/gofiber/fiber/v3"
)

func AuthedMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		tokenInfo, err := utils.GetUserInfoFromJWT(c)
		if err != nil {
			if errors.Is(err, utils.JWTNoTokenErr) || errors.Is(err, utils.JWTInvalidTokenErr) {
				return utils.UnauthentizatedResponse(c, err)
			} else {
				return utils.InternalServerErrorResponse(c, err)
			}
		}
		utils.SetJWTToLocals(c, tokenInfo)
		return c.Next()
	}
}
