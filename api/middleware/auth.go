package middleware

import (
	"fmt"

	"github.com/TDiblik/project-template/api/constants"
	"github.com/TDiblik/project-template/api/utils"
	"github.com/gofiber/fiber/v3"
)

func AuthedMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		token_raw := c.Get(constants.TOKEN_HEADER_NAME)
		if len(token_raw) == 0 {
			return utils.UnauthentizatedResponse(c, fmt.Errorf("token_raw was empty"))
		}

		token_claims, err := utils.ValidateJWT(token_raw)
		if err != nil {
			return utils.UnauthentizatedResponse(c, err)
		}

		tokenInfo, err := utils.TokenClaimsToJwtInfo(token_claims)
		if err != nil {
			return utils.InternalServerErrorResponse(c, err)
		}
		utils.SetJWTToLocals(c, tokenInfo)

		return c.Next()
	}
}
