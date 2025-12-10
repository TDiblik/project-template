package handlers

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/TDiblik/project-template/api/database"
	database_gen "github.com/TDiblik/project-template/api/database/gen"
	"github.com/TDiblik/project-template/api/utils"
	"github.com/gofiber/fiber/v3"
)

type LoginHandlerRequestBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type AuthHandlerResponse struct {
	AuthToken string `json:"auth_token" validate:"required"`
}

func LoginHandler(c fiber.Ctx) error {
	var req LoginHandlerRequestBody
	if err := utils.GetValidRequestBody(&req, c); err != nil {
		return utils.InvalidRequestResponse(c, err)
	}

	db, db_ctx, err := database.CreateConnection()
	if err != nil {
		return utils.InternalServerErrorResponse(c, fmt.Errorf("unable to create db connection: %w", err))
	}

	user, err := db.GetUserAuthByEmail(db_ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return utils.ConflictResponse(c, "be.error.login.username_or_password_incorrect")
		}
		return utils.InternalServerErrorResponse(c, err)
	}

	var userUUID = user.ID
	var userPasswordHash = user.PasswordHash

	// Handle accounts created by oauth without a password
	if !userPasswordHash.Valid || userPasswordHash.String == "" {
		return utils.ConflictResponse(c, "be.error.login.username_or_password_incorrect")
	}

	if !utils.CompareHashAndPassword(userPasswordHash.String, req.Password) {
		return utils.ConflictResponse(c, "be.error.login.username_or_password_incorrect")
	}

	newAuthToken, err := GetJwtPostLogin(userUUID)
	if err != nil {
		return utils.InternalServerErrorResponse(c, fmt.Errorf("unable to execute GetJwtPostLogin: %w", err))
	}

	return utils.OkResponse(c, AuthHandlerResponse{
		AuthToken: newAuthToken,
	})
}

type SignUpHandlerRequestBody struct {
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=6"`
	UseUsername bool   `json:"useUsername"`
	FirstName   string `json:"firstName,omitempty"`
	LastName    string `json:"lastName,omitempty"`
	Username    string `json:"username,omitempty"`
}

func (req *SignUpHandlerRequestBody) validateSignUpRequestBody() error {
	if len([]byte(req.Password)) > 72 {
		return fmt.Errorf("password cannot exceed 72 bytes")
	}
	if req.UseUsername {
		if strings.TrimSpace(req.Username) == "" {
			return fmt.Errorf("username is required")
		}
		if len(req.Username) < 3 {
			return fmt.Errorf("username must have at least 3 characters")
		}
	} else {
		if strings.TrimSpace(req.FirstName) == "" {
			return fmt.Errorf("first name is required")
		}
		if strings.TrimSpace(req.LastName) == "" {
			return fmt.Errorf("last name is required")
		}
	}
	return nil
}

func SignUpHandler(c fiber.Ctx) error {
	var req SignUpHandlerRequestBody
	if err := utils.GetValidRequestBody(&req, c); err != nil {
		return utils.InvalidRequestResponse(c, err)
	}
	if err := req.validateSignUpRequestBody(); err != nil {
		return utils.InvalidRequestResponse(c, err)
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return utils.InternalServerErrorResponse(c, fmt.Errorf("failed to hash password: %w", err))
	}

	db, db_ctx, err := database.CreateConnection()
	if err != nil {
		return utils.InternalServerErrorResponse(c, fmt.Errorf("unable to create db connection: %w", err))
	}

	if emailExists, err := db.CheckEmailExists(db_ctx, req.Email); err != nil {
		return utils.InternalServerErrorResponse(c, fmt.Errorf("failed to check existing user: %w", err))
	} else if emailExists {
		return utils.ConflictResponse(c, "be.error.login.email_already_in_use")
	}

	var newUser = database_gen.User{
		Email:         req.Email,
		EmailVerified: sql.NullBool{Valid: true, Bool: false},
		FirstName:     utils.SQLNullStringFromString(req.FirstName).NullString,
		LastName:      utils.SQLNullStringFromString(req.LastName).NullString,
		PasswordHash:  utils.SQLNullStringFromString(hashedPassword).NullString,
	}
	if req.UseUsername {
		if handleExists, err := db.CheckHandleExists(db_ctx, utils.SQLNullStringFromString(req.Username).NullString); err != nil {
			return utils.InternalServerErrorResponse(c, fmt.Errorf("failed to check existing user: %w", err))
		} else if handleExists {
			return utils.ConflictResponse(c, "be.error.login.handle_already_in_use")
		}
		newUser.Handle = utils.SQLNullStringFromString(req.Username).NullString
	}

	userUUID, err := CreateOrUpdateUser(c, newUser)
	if err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}
	newAuthToken, err := GetJwtPostLogin(userUUID)
	if err != nil {
		return utils.InternalServerErrorResponse(c, fmt.Errorf("unable to execute GetJwtPostLogin: %w", err))
	}

	return utils.OkResponse(c, AuthHandlerResponse{
		AuthToken: newAuthToken,
	})
}
