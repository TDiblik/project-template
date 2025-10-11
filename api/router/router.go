package router

import (
	"time"

	"github.com/TDiblik/gofiber-swagger/gofiberswagger"
	"github.com/TDiblik/project-template/api/constants"
	"github.com/TDiblik/project-template/api/handlers"
	"github.com/TDiblik/project-template/api/middleware"
	"github.com/TDiblik/project-template/api/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func SetupRoutes(app *fiber.App) {
	base := gofiberswagger.NewRouter(app)
	api := base.Group("/api")
	api.Get("/health", &gofiberswagger.RouteInfo{Responses: gofiberswagger.NewResponses(gofiberswagger.NewResponseInfo[struct{}]("200", "ok"))}, func(c fiber.Ctx) error {
		return utils.OkResponse(c, fiber.Map{})
	})
	api_v1 := api.Group("/v1")
	api_v1_public := api_v1.Group("/public")

	// Auth
	api_auth := api_v1_public.Group("/auth")
	api_auth.Post("/login", &gofiberswagger.RouteInfo{
		RequestBody: gofiberswagger.NewRequestBody[handlers.LoginHandlerRequestBody](),
		Responses: utils.NewSwaggerResponsesWithErrors(
			gofiberswagger.NewResponseInfo[handlers.AuthHandlerResponse]("200", "ok"),
		),
	}, handlers.LoginHandler)
	api_auth.Post("/signup", &gofiberswagger.RouteInfo{
		RequestBody: gofiberswagger.NewRequestBody[handlers.SignUpHandlerRequestBody](),
		Responses: utils.NewSwaggerResponsesWithErrors(
			gofiberswagger.NewResponseInfo[handlers.AuthHandlerResponse]("200", "ok"),
		),
	}, handlers.SignUpHandler)

	api_oauth := api_auth.Group("/oauth")
	api_oauth.Get("/return", &gofiberswagger.RouteInfo{
		Parameters: gofiberswagger.NewParameters(
			gofiberswagger.NewQueryParameterRequired("state"),
			gofiberswagger.NewQueryParameterRequired("code"),
		),
		Responses: utils.NewSwaggerResponsesWithErrors(
			gofiberswagger.NewResponseInfo[handlers.OAuthPostReturnHandlerResponse]("200", "ok"),
		),
	}, handlers.OAuthPostReturnHandler)

	api_oauth_redirect := api_oauth.Group("/redirect")
	api_oauth_redirect.Get("/github", &gofiberswagger.RouteInfo{
		Parameters: gofiberswagger.NewParameters(
			gofiberswagger.INewQueryParameter[utils.RedirectAfterOauth]("redirect_back_to_after_oauth"),
		),
		Responses: utils.NewSwaggerResponsesWithErrors(
			gofiberswagger.NewResponseInfo[handlers.OauthRedirectHandlerResponse]("200", "ok"),
		),
	}, handlers.GithubRedirectHandler)
	api_oauth_redirect.Get("/google", &gofiberswagger.RouteInfo{
		Parameters: gofiberswagger.NewParameters(
			gofiberswagger.INewQueryParameter[utils.RedirectAfterOauth]("redirect_back_to_after_oauth"),
		),
		Responses: utils.NewSwaggerResponsesWithErrors(
			gofiberswagger.NewResponseInfo[handlers.OauthRedirectHandlerResponse]("200", "ok"),
		),
	}, handlers.GoogleRedirectHandler)
	api_oauth_redirect.Get("/facebook", &gofiberswagger.RouteInfo{
		Parameters: gofiberswagger.NewParameters(
			gofiberswagger.INewQueryParameter[utils.RedirectAfterOauth]("redirect_back_to_after_oauth"),
		),
		Responses: utils.NewSwaggerResponsesWithErrors(
			gofiberswagger.NewResponseInfo[handlers.OauthRedirectHandlerResponse]("200", "ok"),
		),
	}, handlers.FacebookRedirectHandler)
	api_oauth_redirect.Get("/spotify", &gofiberswagger.RouteInfo{
		Parameters: gofiberswagger.NewParameters(
			gofiberswagger.INewQueryParameter[utils.RedirectAfterOauth]("redirect_back_to_after_oauth"),
		),
		Responses: utils.NewSwaggerResponsesWithErrors(
			gofiberswagger.NewResponseInfo[handlers.OauthRedirectHandlerResponse]("200", "ok"),
		),
	}, handlers.SpotifyRedirectHandler)

	api_v1_private := api_v1.Group("/private")
	api_user := api_v1_private.Group("/user", middleware.AuthedMiddleware())
	api_user.Get("/me", &gofiberswagger.RouteInfo{
		Responses: utils.NewSwaggerResponsesWithErrors(
			gofiberswagger.NewResponseInfo[handlers.UserMeHandlerResponse]("200", "ok"),
		),
	}, handlers.UserMeHandler)

	if utils.EnvData.Debug {
		gofiberswagger.Register(app, gofiberswagger.Config{
			Swagger: gofiberswagger.SwaggerConfig{
				OpenAPI: "3.1.1",
				Info: &gofiberswagger.Info{
					Title:   "project-template",
					Version: "0.0.1",
				},
				Components: &gofiberswagger.Components{
					SecuritySchemes: map[string]*gofiberswagger.SecuritySchemeRef{
						constants.TOKEN_HEADER_NAME: {
							Value: &gofiberswagger.SecurityScheme{
								Type: "apiKey",
								Name: constants.TOKEN_HEADER_NAME,
								In:   "header",
							},
						},
					},
				},
			},
			AutomaticallyRequireAuth: true,
			RequiredAuth: &gofiberswagger.SecurityRequirements{{
				constants.TOKEN_HEADER_NAME: {},
			}},
			SwaggerUI:          gofiberswagger.DefaultUIConfig,
			CreateSwaggerFiles: true,
			SwaggerFilesPath:   "./generated/swagger",
			AppendMethodToTags: false,
			FilterOutAppUse:    true,
		})
	}

	if !utils.EnvData.Debug {
		cache_duration := time.Second * 60 * 60 * 60
		app.Use("/", static.New("./public", static.Config{
			IndexNames:    []string{"index.html"},
			Compress:      true,
			CacheDuration: cache_duration,
		}))
		app.Get("/*", func(c fiber.Ctx) error {
			return c.SendFile("./public/index.html", fiber.SendFile{
				Compress:      true,
				CacheDuration: cache_duration,
			})
		})
	} else {
		api.Get("/", nil, func(c fiber.Ctx) error {
			return utils.OkResponse(c, fiber.Map{
				"message": "this is a default index dev page",
			})
		})
	}
}
