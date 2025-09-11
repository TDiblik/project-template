package router

import (
	"time"

	"github.com/TDiblik/gofiber-swagger/gofiberswagger"
	"github.com/TDiblik/project-template/api/constants"
	"github.com/TDiblik/project-template/api/handlers"
	"github.com/TDiblik/project-template/api/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func SetupRoutes(app *fiber.App) {
	base := gofiberswagger.NewRouter(app)
	api := base.Group("/api")
	api.Get("/health", nil, func(c fiber.Ctx) error {
		return utils.OkResponse(c, fiber.Map{})
	})
	api_v1 := api.Group("/v1")

	// Auth
	api_auth := api_v1.Group("/auth")

	api_oauth := api_auth.Group("/oauth")
	api_oauth.Get("/return", &gofiberswagger.RouteInfo{
		Parameters: gofiberswagger.NewParameters(
			gofiberswagger.NewQueryParameter("state"),
			gofiberswagger.NewQueryParameter("code"),
		),
	}, handlers.OAuthPostReturn)
	api_oauth_redirect := api_oauth.Group("/redirect")
	api_oauth_redirect.Get("/github", &gofiberswagger.RouteInfo{
		Responses: gofiberswagger.NewResponses(
			gofiberswagger.NewResponseInfo[handlers.GithubRedirectResponse]("200", "ok"),
		),
	}, handlers.GithubRedirect)

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
						"x-user-token": {
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
				"x-user-token": {},
			}},
			SwaggerUI:          gofiberswagger.DefaultUIConfig,
			CreateSwaggerFiles: true,
			SwaggerFilesPath:   "./generated/swagger",
			AppendMethodToTags: false,
			FilterOutAppUse:    true,
		})
	}

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
}
