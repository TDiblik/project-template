package handlers

import (
	"github.com/TDiblik/project-template/api/database"
	"github.com/TDiblik/project-template/api/models"
	"github.com/TDiblik/project-template/api/utils"
	"github.com/gofiber/fiber/v3"
)

type GetUserMeHandlerResponse struct {
	UserInfo models.UserModelDB `json:"user_info"`
}

func GetUserMeHandler(c fiber.Ctx) error {
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

	return utils.OkResponse(c, GetUserMeHandlerResponse{UserInfo: userInfo})
}

type PatchUserMeHandlerRequest struct {
	FirstName        string                          `json:"first_name,omitempty" validate:"omitempty,min=1,max=50"`
	LastName         string                          `json:"last_name,omitempty" validate:"omitempty,min=1,max=50"`
	PreferedTheme    utils.ThemePosibilities         `json:"prefered_theme,omitempty" validate:"omitempty,oneof=light dark"`
	PreferedLanguage utils.TranslationsPossibilities `json:"prefered_language,omitempty" validate:"omitempty,oneof=cs en"`
}
type PatchUserMeHandlerResponse struct{}

func PatchUserMeHandler(c fiber.Ctx) error {
	userJWTInfo, err := utils.GetJWTFromLocals(c)
	if err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}

	var req PatchUserMeHandlerRequest
	if err := utils.GetValidRequestBody(&req, c); err != nil {
		return utils.InvalidRequestResponse(c, err)
	}

	db, err := database.CreateConnection()
	if err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}

	var user models.UserModelDB
	if err := db.Get(&user, "select * from users where id = $1", userJWTInfo.UserId); err != nil {
		return utils.NotFoundResponse(c, "be.error.user.not_found")
	}

	if req.FirstName != "" {
		user.FirstName = utils.SQLNullStringFromString(req.FirstName)
	}
	if req.LastName != "" {
		user.LastName = utils.SQLNullStringFromString(req.LastName)
	}
	if req.PreferedTheme != "" {
		user.PreferedTheme = utils.SQLNullStringFromString(string(req.PreferedTheme))
	}
	if req.PreferedLanguage != "" {
		user.PreferedLanguage = utils.SQLNullStringFromString(string(req.PreferedLanguage))
	}

	_, err = db.NamedExec(`
		update users set
			first_name = :first_name,
			last_name = :last_name,
			prefered_theme = :prefered_theme,
			prefered_language = :prefered_language
		where id = :id
	`, user)
	if err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}

	return utils.OkResponse(c, PatchUserMeHandlerResponse{})
}
