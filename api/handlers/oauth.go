package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/TDiblik/project-template/api/database"
	"github.com/TDiblik/project-template/api/models"
	"github.com/TDiblik/project-template/api/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

const githubProviderName = "GitHub"

type GithubRedirectResponse struct {
	OAuthState  string `json:"oauth_state"`
	RedirectURL string `json:"redirect_url"`
}

// todo: add param: isMobile, redirectToPage
func GithubRedirect(c fiber.Ctx) error {
	state, err := utils.GenerateOauthState(githubProviderName)
	if err != nil {
		return utils.InternalServerErrorResponse(c, fmt.Errorf("failed to generate OAuth state: %w", err))
	}
	utils.Log(state)
	return utils.OkResponse(c, GithubRedirectResponse{
		OAuthState:  state,
		RedirectURL: utils.EnvData.OAUTH_CONFIG_GITHUB.AuthCodeURL(state),
	})
}

type OAuthPostReturQuery struct {
	State string `query:"state" validate:"required"`
	Code  string `query:"code" validate:"required"`
}

type OAuthPostReturnResponse struct {
	AuthToken string `json:"auth_token" validate:"required"`
	// todo: Add redirectToPage, which is taken from token
}

func OAuthPostReturn(c fiber.Ctx) error {
	var query OAuthPostReturQuery
	if err := utils.GetValidQuery(&query, c); err != nil {
		return utils.InvalidRequestResponse(c, err)
	}

	if query.State == "" || !utils.IsValidOauthState(query.State) {
		return utils.UnauthorizedResponse(c, fmt.Errorf("invalid OAuth state"))
	}

	provider, err := utils.GetOauthProvider(query.State)
	if err != nil {
		return utils.InvalidRequestResponse(c, fmt.Errorf("invalid provider name inside the state: %w", err))
	}

	var userUUID uuid.UUID
	// when adding a new oauth provider and user table fields, add new "else if" here:
	if provider == githubProviderName {
		if userUUID, err = githubReturn(query.Code); err != nil {
			return utils.InternalServerErrorResponse(c, err)
		}
	} else {
		return utils.InvalidRequestResponse(c, fmt.Errorf("invalid provider name (not implemented) inside the state: %w", err))
	}

	db, err := database.CreateConnection()
	if err != nil {
		return fmt.Errorf("unable to create connection to db inside CreateOrUpdateUser: %w", err)
	}

	var userInfo models.UsersModelDB
	err = db.Get(&userInfo, `select * from users where id = $1`, userUUID)
	if err != nil {
		return utils.InternalServerErrorResponse(c, fmt.Errorf("unable to select existing user: %w", err))
	}

	newAuthToken, err := utils.GenerateJWT(userInfo)
	if err != nil {
		return utils.InternalServerErrorResponse(c, fmt.Errorf("unable to generate new oauth token: %w", err))
	}

	return utils.OkResponse(c, OAuthPostReturnResponse{
		AuthToken: newAuthToken,
	})
}

type githubUserApiResponse struct {
	ID        int64  `json:"id"`
	Login     string `json:"login"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url"`
}

func githubReturn(authCode string) (uuid.UUID, error) {
	token, err := utils.EnvData.OAUTH_CONFIG_GITHUB.Exchange(context.Background(), authCode)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to exchange token: %w", err)
	}

	client := utils.EnvData.OAUTH_CONFIG_GITHUB.Client(context.Background(), token)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to get user info: %w", err)
	}
	defer resp.Body.Close()

	var ghUserResponse githubUserApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&ghUserResponse); err != nil {
		return uuid.Nil, fmt.Errorf("failed to decode user info: %w", err)
	}
	var firstName, lastName string
	if ghUserResponse.Name != "" {
		parts := strings.SplitN(ghUserResponse.Name, " ", 2)
		firstName = parts[0]
		if len(parts) > 1 {
			lastName = parts[1]
		}
	}

	return CreateOrUpdateUser(models.UsersModelDB{
		Email:         ghUserResponse.Email,
		EmailVerified: true,
		FirstName:     utils.SQLNullStringFromString(firstName),
		LastName:      utils.SQLNullStringFromString(lastName),
		Handle:        utils.SQLNullStringFromString(ghUserResponse.Login),
		GithubId:      utils.SQLNullStringFromString(strconv.FormatInt(ghUserResponse.ID, 10)),
		GithubHandle:  utils.SQLNullStringFromString(ghUserResponse.Login),
		AvatarUrl:     utils.SQLNullStringFromString(ghUserResponse.AvatarURL),
	})
}

func CreateOrUpdateUser(possiblyNewUser models.UsersModelDB) (uuid.UUID, error) {
	db, err := database.CreateConnection()
	if err != nil {
		return uuid.Nil, fmt.Errorf("unable to create connection to db inside CreateOrUpdateUser: %w", err)
	}

	var emailExists, handleExists bool
	err = db.QueryRow(`select 
		exists(select 1 from users where email = $1) as email_exists,
		exists(select 1 from users where handle = $2) as handle_exists`,
		possiblyNewUser.Email, possiblyNewUser.Handle).Scan(&emailExists, &handleExists)
	if err != nil {
		return uuid.Nil, fmt.Errorf("unable to query exists staments inside CreateOrUpdateUser: %w", err)
	}

	var existingUser models.UsersModelDB
	if !emailExists {
		if handleExists {
			possiblyNewUser.Handle = models.SQLNullString{}
		}

		rows, err := db.NamedQuery(`
			insert into users (
				email, email_verified, handle, first_name, last_name, avatar_url, github_id, github_handle
			) 
			values (
				:email, :email_verified, :handle, :first_name, :last_name, :avatar_url, :github_id, :github_handle
			)
			returning *
		`, possiblyNewUser)
		if err != nil {
			return uuid.Nil, fmt.Errorf("unable to insert new user: %w", err)
		}
		defer rows.Close()
		if !rows.Next() {
			return uuid.Nil, fmt.Errorf("no user inserted, but no error provided ? this should never occured, user object: %+v", possiblyNewUser)
		}
		if err := rows.StructScan(&existingUser); err != nil {
			return uuid.Nil, fmt.Errorf("unable to scan inserted user: %w", err)
		}
	}
	if emailExists {
		err := db.Get(&existingUser, `select * from users where email = $1`, possiblyNewUser.Email)
		if err != nil {
			return uuid.Nil, fmt.Errorf("unable to select existing user: %w", err)
		}

		if !existingUser.Handle.Valid && !handleExists {
			existingUser.Handle = possiblyNewUser.Handle
		}
		if !existingUser.FirstName.Valid {
			existingUser.FirstName = possiblyNewUser.FirstName
		}
		if !existingUser.LastName.Valid {
			existingUser.LastName = possiblyNewUser.LastName
		}
		if !existingUser.AvatarUrl.Valid {
			existingUser.AvatarUrl = possiblyNewUser.AvatarUrl
		}
		if !existingUser.EmailVerified {
			existingUser.EmailVerified = possiblyNewUser.EmailVerified
		}

		// when adding a new oauth provider and user table fields, add the checks here:
		if !existingUser.GithubId.Valid {
			existingUser.GithubId = possiblyNewUser.GithubId
		}
		if !existingUser.GithubHandle.Valid {
			existingUser.GithubHandle = possiblyNewUser.GithubHandle
		}

		// when adding a new oauth provider and user table fields, change the query here:
		if _, err := db.NamedExec(`
			update users set
				handle = :handle,
				first_name = :first_name,
				last_name = :last_name,
				avatar_url = :avatar_url,
				email_verified = :email_verified,
				github_id = :github_id,
				github_handle = :github_handle
			where id = :id
		`, existingUser); err != nil {
			return uuid.Nil, fmt.Errorf("unable to update existing user: %w", err)
		}
	}

	return existingUser.Id, nil
}
