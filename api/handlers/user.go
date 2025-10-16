package handlers

import (
	"github.com/TDiblik/project-template/api/database"
	"github.com/TDiblik/project-template/api/models"
	"github.com/TDiblik/project-template/api/utils"
	"github.com/gofiber/fiber/v3"
)

type UserMeHandlerResponse struct {
	UserInfo models.UserModelDB `json:"user_info"`
}

func UserMeHandler(c fiber.Ctx) error {
	userJWTInfo, err := utils.GetJWTFromLocals(c)
	if err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}

	db, err := database.CreateConnection()
	if err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}

	var userInfo models.UserModelDB
	if err := db.Get(&userInfo, "select * from users where id = $1", userJWTInfo.UserId); err != nil {
		return utils.NotFoundResponse(c, "be.error.user.not_found")
	}

	return utils.OkResponse(c, UserMeHandlerResponse{UserInfo: userInfo})
}
