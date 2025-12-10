package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/TDiblik/project-template/api/database"
	database_gen "github.com/TDiblik/project-template/api/database/gen"
	"github.com/TDiblik/project-template/api/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

const (
	githubProviderName   = "GitHub"
	googleProviderName   = "Google"
	facebookProviderName = "Facebook"
	spotifyProviderName  = "Spotify"
)

type OauthRedirectHandlerResponse struct {
	OAuthState  string `json:"oauth_state"`
	RedirectURL string `json:"redirect_url"`
}

func GithubRedirectHandler(c fiber.Ctx) error {
	redirectParam := c.Query("redirect_back_to_after_oauth", string(utils.RedirectAfterOauthIndex))
	redirectBackTo := utils.ValidateRedirectAfterOauth(redirectParam)

	state, err := utils.GenerateOauthState(githubProviderName, redirectBackTo)
	if err != nil {
		return utils.InternalServerErrorResponse(c, fmt.Errorf("failed to generate OAuth state: %w", err))
	}
	return utils.OkResponse(c, OauthRedirectHandlerResponse{
		OAuthState:  state,
		RedirectURL: utils.EnvData.OAUTH_GITHUB_CONFIG.AuthCodeURL(state),
	})
}

func GoogleRedirectHandler(c fiber.Ctx) error {
	redirectParam := c.Query("redirect_back_to_after_oauth", string(utils.RedirectAfterOauthIndex))
	redirectBackTo := utils.ValidateRedirectAfterOauth(redirectParam)

	state, err := utils.GenerateOauthState(googleProviderName, redirectBackTo)
	if err != nil {
		return utils.InternalServerErrorResponse(c, fmt.Errorf("failed to generate OAuth state: %w", err))
	}
	return utils.OkResponse(c, OauthRedirectHandlerResponse{
		OAuthState:  state,
		RedirectURL: utils.EnvData.OAUTH_GOOGLE_CONFIG.AuthCodeURL(state),
	})
}

func FacebookRedirectHandler(c fiber.Ctx) error {
	redirectParam := c.Query("redirect_back_to_after_oauth", string(utils.RedirectAfterOauthIndex))
	redirectBackTo := utils.ValidateRedirectAfterOauth(redirectParam)

	state, err := utils.GenerateOauthState(facebookProviderName, redirectBackTo)
	if err != nil {
		return utils.InternalServerErrorResponse(c, fmt.Errorf("failed to generate OAuth state: %w", err))
	}
	return utils.OkResponse(c, OauthRedirectHandlerResponse{
		OAuthState:  state,
		RedirectURL: utils.EnvData.OAUTH_FACEBOOK_CONFIG.AuthCodeURL(state),
	})
}

func SpotifyRedirectHandler(c fiber.Ctx) error {
	redirectParam := c.Query("redirect_back_to_after_oauth", string(utils.RedirectAfterOauthIndex))
	redirectBackTo := utils.ValidateRedirectAfterOauth(redirectParam)

	state, err := utils.GenerateOauthState(spotifyProviderName, redirectBackTo)
	if err != nil {
		return utils.InternalServerErrorResponse(c, fmt.Errorf("failed to generate OAuth state: %w", err))
	}
	return utils.OkResponse(c, OauthRedirectHandlerResponse{
		OAuthState:  state,
		RedirectURL: utils.EnvData.OAUTH_SPOTIFY_CONFIG.AuthCodeURL(state),
	})
}

type OAuthPostReturnHandlerQuery struct {
	State string `query:"state" validate:"required"`
	Code  string `query:"code" validate:"required"`
}

type OAuthPostReturnHandlerResponse struct {
	AuthToken                string                   `json:"auth_token" validate:"required"`
	RedirectBackToAfterOauth utils.RedirectAfterOauth `json:"redirect_back_to_after_oauth" validate:"required"`
}

func OAuthPostReturnHandler(c fiber.Ctx) error {
	var query OAuthPostReturnHandlerQuery
	if err := utils.GetValidQuery(&query, c); err != nil {
		return utils.InvalidRequestResponse(c, err)
	}

	if query.State == "" || !utils.IsValidOauthState(query.State) {
		return utils.UnauthorizedResponse(c, fmt.Errorf("invalid OAuth state"))
	}

	provider, redirect, err := utils.GetOauthProviderAndRedirectFromOauthState(query.State)
	if err != nil {
		return utils.InvalidRequestResponse(c, fmt.Errorf("invalid provider or redirect info inside the state: %w", err))
	}

	var userUUID uuid.UUID
	switch provider {
	case githubProviderName:
		if userUUID, err = githubReturn(c, query.Code); err != nil {
			return utils.InternalServerErrorResponse(c, err)
		}
	case googleProviderName:
		if userUUID, err = googleReturn(c, query.Code); err != nil {
			return utils.InternalServerErrorResponse(c, err)
		}
	case facebookProviderName:
		if userUUID, err = facebookReturn(c, query.Code); err != nil {
			return utils.InternalServerErrorResponse(c, err)
		}
	case spotifyProviderName:
		if userUUID, err = spotifyReturn(c, query.Code); err != nil {
			return utils.InternalServerErrorResponse(c, err)
		}
	default:
		return utils.InvalidRequestResponse(c, fmt.Errorf("invalid provider name (not implemented) inside the state: %w", err))
	}

	newAuthToken, err := GetJwtPostLogin(userUUID)
	if err != nil {
		return utils.InternalServerErrorResponse(c, fmt.Errorf("unable to execute GetJwtPostLogin: %w", err))
	}
	return utils.OkResponse(c, OAuthPostReturnHandlerResponse{
		AuthToken:                newAuthToken,
		RedirectBackToAfterOauth: redirect,
	})
}

type githubUserApiResponse struct {
	ID        int64  `json:"id"`
	Login     string `json:"login"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url"`
	HtmlUrl   string `json:"html_url"`
}

func githubReturn(c fiber.Ctx, authCode string) (uuid.UUID, error) {
	token, err := utils.EnvData.OAUTH_GITHUB_CONFIG.Exchange(context.Background(), authCode)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to exchange token: %w", err)
	}

	client := utils.EnvData.OAUTH_GITHUB_CONFIG.Client(context.Background(), token)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to get user info: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return uuid.Nil, fmt.Errorf("GitHub API returned status %d: %s", resp.StatusCode, string(bodyBytes))
	}

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

	return CreateOrUpdateUser(c, database_gen.User{
		Email:         ghUserResponse.Email,
		EmailVerified: sql.NullBool{Valid: true, Bool: true},
		FirstName:     utils.SQLNullStringFromString(firstName).NullString,
		LastName:      utils.SQLNullStringFromString(lastName).NullString,
		Handle:        utils.SQLNullStringFromString(ghUserResponse.Login).NullString,
		GithubID:      utils.SQLNullStringFromString(strconv.FormatInt(ghUserResponse.ID, 10)),
		GithubHandle:  utils.SQLNullStringFromString(ghUserResponse.Login),
		GithubEmail:   utils.SQLNullStringFromString(ghUserResponse.Email),
		GithubUrl:     utils.SQLNullStringFromString(ghUserResponse.HtmlUrl),
		AvatarUrl:     utils.SQLNullStringFromString(ghUserResponse.AvatarURL),
	})
}

type googleUserApiResponse struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
}

func googleReturn(c fiber.Ctx, authCode string) (uuid.UUID, error) {
	token, err := utils.EnvData.OAUTH_GOOGLE_CONFIG.Exchange(context.Background(), authCode)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to exchange token: %w", err)
	}

	client := utils.EnvData.OAUTH_GOOGLE_CONFIG.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to get user info: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return uuid.Nil, fmt.Errorf("google API returned status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var gUser googleUserApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&gUser); err != nil {
		return uuid.Nil, fmt.Errorf("failed to decode user info: %w", err)
	}

	return CreateOrUpdateUser(c, database_gen.User{
		Email:         gUser.Email,
		EmailVerified: sql.NullBool{Valid: true, Bool: gUser.VerifiedEmail},
		FirstName:     utils.SQLNullStringFromString(gUser.GivenName).NullString,
		LastName:      utils.SQLNullStringFromString(gUser.FamilyName).NullString,
		GoogleID:      utils.SQLNullStringFromString(gUser.ID),
		GoogleEmail:   utils.SQLNullStringFromString(gUser.Email),
		AvatarUrl:     utils.SQLNullStringFromString(gUser.Picture),
	})
}

type facebookUserApiResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Picture   struct {
		Data struct {
			URL string `json:"url"`
		} `json:"data"`
	} `json:"picture"`
	Link string `json:"link"`
}

func facebookReturn(c fiber.Ctx, authCode string) (uuid.UUID, error) {
	token, err := utils.EnvData.OAUTH_FACEBOOK_CONFIG.Exchange(context.Background(), authCode)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to exchange token: %w", err)
	}

	client := utils.EnvData.OAUTH_FACEBOOK_CONFIG.Client(context.Background(), token)
	resp, err := client.Get("https://graph.facebook.com/me?fields=id,first_name,last_name,email,picture,link")
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to get user info: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return uuid.Nil, fmt.Errorf("facebook API returned status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var fbUser facebookUserApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&fbUser); err != nil {
		return uuid.Nil, fmt.Errorf("failed to decode user info: %w", err)
	}

	return CreateOrUpdateUser(c, database_gen.User{
		Email:         fbUser.Email,
		EmailVerified: sql.NullBool{Valid: true, Bool: true},
		FirstName:     utils.SQLNullStringFromString(fbUser.FirstName).NullString,
		LastName:      utils.SQLNullStringFromString(fbUser.LastName).NullString,
		FacebookID:    utils.SQLNullStringFromString(fbUser.ID),
		FacebookEmail: utils.SQLNullStringFromString(fbUser.Email),
		FacebookUrl:   utils.SQLNullStringFromString(fbUser.Link),
		AvatarUrl:     utils.SQLNullStringFromString(fbUser.Picture.Data.URL),
	})
}

type spotifyUserApiResponse struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	Images      []struct {
		URL string `json:"url"`
	} `json:"images"`
	ExternalURLs struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
}

func spotifyReturn(c fiber.Ctx, authCode string) (uuid.UUID, error) {
	token, err := utils.EnvData.OAUTH_SPOTIFY_CONFIG.Exchange(context.Background(), authCode)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to exchange token: %w", err)
	}

	client := utils.EnvData.OAUTH_SPOTIFY_CONFIG.Client(context.Background(), token)
	resp, err := client.Get("https://api.spotify.com/v1/me")
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to get user info: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return uuid.Nil, fmt.Errorf("spotify API returned status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var spUser spotifyUserApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&spUser); err != nil {
		return uuid.Nil, fmt.Errorf("failed to decode user info: %w", err)
	}

	var firstName, lastName string
	if spUser.DisplayName != "" {
		parts := strings.SplitN(spUser.DisplayName, " ", 2)
		firstName = parts[0]
		if len(parts) > 1 {
			lastName = parts[1]
		}
	}

	var avatarURL string
	if len(spUser.Images) > 0 {
		avatarURL = spUser.Images[0].URL
	}

	return CreateOrUpdateUser(c, database_gen.User{
		Email:         spUser.Email,
		EmailVerified: sql.NullBool{Valid: true, Bool: true},
		FirstName:     utils.SQLNullStringFromString(firstName).NullString,
		LastName:      utils.SQLNullStringFromString(lastName).NullString,
		SpotifyID:     utils.SQLNullStringFromString(spUser.ID),
		SpotifyEmail:  utils.SQLNullStringFromString(spUser.Email),
		SpotifyUrl:    utils.SQLNullStringFromString(spUser.ExternalURLs.Spotify),
		AvatarUrl:     utils.SQLNullStringFromString(avatarURL),
	})
}

func CreateOrUpdateUser(c fiber.Ctx, possiblyNewUser database_gen.User) (uuid.UUID, error) {
	queries, ctx, err := database.CreateConnection()
	if err != nil {
		return uuid.Nil, fmt.Errorf("unable to create connection to db: %w", err)
	}

	possiblyUserEmail := possiblyNewUser.Email
	loggedInUserInfo, err := utils.GetUserInfoFromJWT(c)
	processingLoggedInUser := loggedInUserInfo != nil
	if err == nil && processingLoggedInUser {
		possiblyUserEmail = loggedInUserInfo.UserEmail
	}

	emailExists, err := queries.CheckEmailExists(ctx, possiblyUserEmail)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to check email exists: %w", err)
	}

	if !emailExists {
		finalHandle := possiblyNewUser.Handle
		if !finalHandle.Valid || finalHandle.String == "" || finalHandle.String == "user" {
			uniqueHandle, err := getUniqueHandle(ctx, queries, possiblyNewUser.FirstName, possiblyNewUser.LastName)
			if err != nil {
				return uuid.Nil, fmt.Errorf("failed to generate unique handle: %w", err)
			}
			finalHandle = utils.SQLNullStringFromString(uniqueHandle).NullString
		} else {
			exists, err := queries.CheckHandleExists(ctx, finalHandle)
			if err != nil {
				return uuid.Nil, err
			}
			if exists {
				uniqueHandle, err := getUniqueHandle(ctx, queries, possiblyNewUser.FirstName, possiblyNewUser.LastName)
				if err != nil {
					return uuid.Nil, err
				}
				finalHandle = utils.SQLNullStringFromString(uniqueHandle).NullString
			}
		}

		createdUser, err := queries.CreateUser(ctx, database_gen.CreateUserParams{
			Email:         possiblyNewUser.Email,
			EmailVerified: possiblyNewUser.EmailVerified,
			PasswordHash:  possiblyNewUser.PasswordHash,
			Handle:        finalHandle,
			FirstName:     possiblyNewUser.FirstName,
			LastName:      possiblyNewUser.LastName,
			AvatarUrl:     possiblyNewUser.AvatarUrl,
			GithubID:      possiblyNewUser.GithubID,
			GithubEmail:   possiblyNewUser.GithubEmail,
			GithubHandle:  possiblyNewUser.GithubHandle,
			GithubUrl:     possiblyNewUser.GithubUrl,
			GoogleID:      possiblyNewUser.GoogleID,
			GoogleEmail:   possiblyNewUser.GoogleEmail,
			FacebookID:    possiblyNewUser.FacebookID,
			FacebookEmail: possiblyNewUser.FacebookEmail,
			FacebookUrl:   possiblyNewUser.FacebookUrl,
			SpotifyID:     possiblyNewUser.SpotifyID,
			SpotifyEmail:  possiblyNewUser.SpotifyEmail,
			SpotifyUrl:    possiblyNewUser.SpotifyUrl,
		})
		if err != nil {
			return uuid.Nil, fmt.Errorf("unable to insert new user: %w", err)
		}
		return createdUser.ID, nil
	}

	existingUser, err := queries.GetUserByEmailOrOauth(ctx, possiblyUserEmail)
	if err != nil {
		return uuid.Nil, fmt.Errorf("unable to select existing user: %w", err)
	}

	if processingLoggedInUser && loggedInUserInfo.UserId != existingUser.ID {
		return uuid.Nil, fmt.Errorf("email is already linked to a different account")
	}

	updatedUser, err := queries.UpdateUserFull(ctx, database_gen.UpdateUserFullParams{
		ID:            existingUser.ID,
		Column2:       possiblyNewUser.PasswordHash.String,
		Handle:        possiblyNewUser.Handle,
		FirstName:     possiblyNewUser.FirstName,
		LastName:      possiblyNewUser.LastName,
		AvatarUrl:     possiblyNewUser.AvatarUrl,
		EmailVerified: possiblyNewUser.EmailVerified,
		GithubID:      possiblyNewUser.GithubID,
		GithubEmail:   possiblyNewUser.GithubEmail,
		GithubHandle:  possiblyNewUser.GithubHandle,
		GithubUrl:     possiblyNewUser.GithubUrl,
		GoogleID:      possiblyNewUser.GoogleID,
		GoogleEmail:   possiblyNewUser.GoogleEmail,
		FacebookID:    possiblyNewUser.FacebookID,
		FacebookEmail: possiblyNewUser.FacebookEmail,
		FacebookUrl:   possiblyNewUser.FacebookUrl,
		SpotifyID:     possiblyNewUser.SpotifyID,
		SpotifyEmail:  possiblyNewUser.SpotifyEmail,
		SpotifyUrl:    possiblyNewUser.SpotifyUrl,
	})

	if err != nil {
		return uuid.Nil, fmt.Errorf("unable to update existing user: %w", err)
	}

	return updatedUser.ID, nil
}

func GetJwtPostLogin(userUUID uuid.UUID) (string, error) {
	queries, ctx, err := database.CreateConnection()
	if err != nil {
		return "", fmt.Errorf("unable to create connection to db: %w", err)
	}

	updatedUser, err := queries.SetLastLoginNow(ctx, userUUID)
	if err != nil {
		return "", fmt.Errorf("unable to set last login: %w", err)
	}
	newAuthToken, err := utils.GenerateJWT(updatedUser)
	if err != nil {
		return "", fmt.Errorf("unable to generate new oauth token: %w", err)
	}

	return newAuthToken, nil
}

func getUniqueHandle(ctx context.Context, q *database_gen.Queries, firstName, lastName sql.NullString) (string, error) {
	base := "user"
	if firstName.Valid && firstName.String != "" {
		base = utils.NormalizeHandle(firstName.String)
		if lastName.Valid && lastName.String != "" {
			base = utils.NormalizeHandle(string(firstName.String[0]) + lastName.String)
		}
	}

	handle := base
	randomSuffixLen := 4
	if base == "user" {
		randomSuffixLen = 6
	}

	for {
		exists, err := q.CheckHandleExists(ctx, utils.SQLNullStringFromString(handle).NullString)
		if err != nil {
			return "", err
		}
		if !exists {
			return handle, nil
		}
		handle = fmt.Sprintf("%s-%s", base, utils.RandomString(randomSuffixLen))
	}
}
