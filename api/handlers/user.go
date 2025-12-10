package handlers

import (
	"github.com/TDiblik/project-template/api/database"
	database_gen "github.com/TDiblik/project-template/api/database/gen"
	"github.com/TDiblik/project-template/api/utils"
	"github.com/gofiber/fiber/v3"
)

type GetUserMeHandlerResponse struct {
	UserInfo database_gen.User `json:"user_info"`
}

func GetUserMeHandler(c fiber.Ctx) error {
	userJWTInfo, err := utils.GetJWTFromLocals(c)
	if err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}

	db, db_ctx, err := database.CreateConnection()
	if err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}

	user, err := db.GetUserById(db_ctx, userJWTInfo.UserId)
	if err != nil {
		return utils.NotFoundResponse(c, "be.error.user.not_found")
	}

	return utils.OkResponse(c, GetUserMeHandlerResponse{UserInfo: user})
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

	db, db_ctx, err := database.CreateConnection()
	if err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}

	user, err := db.GetUserById(db_ctx, userJWTInfo.UserId)
	if err != nil {
		return utils.NotFoundResponse(c, "be.error.user.not_found")
	}

	if req.FirstName != "" {
		user.FirstName = utils.SQLNullStringFromString(req.FirstName).NullString
	}
	if req.LastName != "" {
		user.LastName = utils.SQLNullStringFromString(req.LastName).NullString
	}
	if req.PreferedTheme != "" {
		user.PreferedTheme = utils.SQLNullStringFromString(string(req.PreferedTheme))
	}
	if req.PreferedLanguage != "" {
		user.PreferedLanguage = utils.SQLNullStringFromString(string(req.PreferedLanguage))
	}

	_, err = db.UpdateUserPreferences(db_ctx, database_gen.UpdateUserPreferencesParams{
		FirstName:        user.FirstName.String,
		LastName:         user.LastName.String,
		PreferedTheme:    user.PreferedTheme.(string),
		PreferedLanguage: user.PreferedLanguage.(string),
	})
	if err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}

	return utils.OkResponse(c, PatchUserMeHandlerResponse{})
}

type PostUserMeAvatarHandlerResponse struct{}

func PostUserMeAvatarHandler(c fiber.Ctx) error {
	file, err := c.FormFile("avatar")
	if err != nil {
		return utils.InvalidRequestResponse(c, err)
	}

	userJWTInfo, err := utils.GetJWTFromLocals(c)
	if err != nil {
		return utils.InvalidRequestResponse(c, err)
	}

	db, db_ctx, err := database.CreateConnection()
	if err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}

	user, err := db.GetUserById(db_ctx, userJWTInfo.UserId)
	if err != nil {
		return utils.NotFoundResponse(c, "be.error.user.not_found")
	}

	newAvatarImageUUID, err := utils.SaveImage(c, file, utils.GetAvatarImageFolder(), 450, 450)
	if err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}
	newAvatarImageUrlPath, err := utils.GetAvatarImageUrl(newAvatarImageUUID)
	if err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}
	user.AvatarUrl = utils.SQLNullStringFromString(newAvatarImageUrlPath)

	_, err = db.UpdateUserAvatar(db_ctx, database_gen.UpdateUserAvatarParams{
		ID:        user.ID,
		AvatarUrl: user.AvatarUrl,
	})
	if err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}

	return utils.OkResponse(c, PostUserMeAvatarHandlerResponse{})
}
