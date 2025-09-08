package utils

import (
	"io/fs"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type IEnvData struct {
	Debug bool

	API_PORT                  string
	API_PROD_URL              string
	FE_PROD_URL               string
	DB_CONNECTION_STRING      string
	DB_MIGRATIONS_PATH        string
	DB_DEV_FORCE_MIGRATE_DOWN bool

	AUTH_SECRET       string
	AUTH_SECRET_BYTES []byte

	OAUTH_GITHUB_CLIENT_ID     string
	OAUTH_GITHUB_CLIENT_SECRET string
	OAUTH_CONFIG_GITHUB        *oauth2.Config

	// MAIL_CLIENT_AUTOMATION_HOST     string
	// MAIL_CLIENT_AUTOMATION_PORT     int
	// MAIL_CLIENT_AUTOMATION_USERNAME string
	// MAIL_CLIENT_AUTOMATION_PASSWORD string

	// MAIL_CLIENT_SALES_HOST     string
	// MAIL_CLIENT_SALES_PORT     int
	// MAIL_CLIENT_SALES_USERNAME string
	// MAIL_CLIENT_SALES_PASSWORD string
}

var EnvData IEnvData
var FoldrePerms fs.FileMode = 0o777

func SetupENV(env_files ...string) {
	log.Println("Setting up env variables: start")

	err := godotenv.Load(env_files...)
	if err != nil {
		log.Println("Unable to load .env file: ", err)
		log.Println("This is normal in production environments, since all environment variables are set in the cloud.")
	}

	switch os.Getenv("GO_ENV") {
	case "development":
		EnvData.Debug = true
		log.SetPrefix("[DEBUG] " + log.Prefix())
	case "production":
		EnvData.Debug = false
		log.SetPrefix("[PROD] " + log.Prefix())
	default:
		log.Fatalln("Error determening GO_ENV (", os.Getenv("GO_ENV"), ")")
	}

	EnvData.API_PORT = getEnvKeyOrPanic("API_PORT")
	EnvData.DB_CONNECTION_STRING = getEnvKeyOrPanic("DB_CONNECTION_STRING")
	EnvData.DB_MIGRATIONS_PATH = getEnvKeyOrPanic("DB_MIGRATIONS_PATH")
	if strings.ToLower(os.Getenv("DB_DEV_FORCE_MIGRATE_DOWN")) == "true" {
		if !EnvData.Debug {
			log.Fatalln("Cannot use DB_DEV_FORCE_MIGRATE_DOWN while in production mode!")
		}
		EnvData.DB_DEV_FORCE_MIGRATE_DOWN = true
	} else {
		EnvData.DB_DEV_FORCE_MIGRATE_DOWN = false
	}
	EnvData.API_PROD_URL = getEnvKeyOrPanic("API_PROD_URL")
	if !strings.HasSuffix(EnvData.API_PROD_URL, "/") {
		EnvData.API_PROD_URL += "/"
	}
	EnvData.FE_PROD_URL = getEnvKeyOrPanic("FE_PROD_URL")
	if !strings.HasSuffix(EnvData.FE_PROD_URL, "/") {
		EnvData.FE_PROD_URL += "/"
	}

	EnvData.AUTH_SECRET = getEnvKeyOrPanic("AUTH_SECRET")
	EnvData.AUTH_SECRET_BYTES = []byte(EnvData.AUTH_SECRET)

	// when adding a new oauth provider and user table fields, add the checks here:
	EnvData.OAUTH_GITHUB_CLIENT_ID = getEnvKeyOrPanic("OAUTH_GITHUB_CLIENT_ID")
	EnvData.OAUTH_GITHUB_CLIENT_SECRET = getEnvKeyOrPanic("OAUTH_GITHUB_CLIENT_SECRET")
	EnvData.OAUTH_CONFIG_GITHUB = &oauth2.Config{
		ClientID:     EnvData.OAUTH_GITHUB_CLIENT_ID,
		ClientSecret: EnvData.OAUTH_GITHUB_CLIENT_SECRET,
		Scopes:       []string{"read:user", "user:email"},
		Endpoint:     github.Endpoint,
		RedirectURL:  JoinUrlOrPanic(EnvData.API_PROD_URL, "/api/v1/auth/oauth/github/return"),
	}

	// EnvData.MAIL_CLIENT_AUTOMATION_HOST = getEnvKeyOrPanic("MAIL_CLIENT_AUTOMATION_HOST")
	// if port, err := strconv.ParseInt(getEnvKeyOrPanic("MAIL_CLIENT_AUTOMATION_PORT"), 10, 64); err != nil {
	// 	log.Fatalf("Invalid MAIL_CLIENT_AUTOMATION_PORT: %v", err)
	// } else {
	// 	EnvData.MAIL_CLIENT_AUTOMATION_PORT = int(port)
	// }
	// EnvData.MAIL_CLIENT_AUTOMATION_USERNAME = getEnvKeyOrPanic("MAIL_CLIENT_AUTOMATION_USERNAME")
	// EnvData.MAIL_CLIENT_AUTOMATION_PASSWORD = getEnvKeyOrPanic("MAIL_CLIENT_AUTOMATION_PASSWORD")

	// EnvData.MAIL_CLIENT_SALES_HOST = getEnvKeyOrPanic("MAIL_CLIENT_SALES_HOST")
	// if port, err := strconv.ParseInt(getEnvKeyOrPanic("MAIL_CLIENT_SALES_PORT"), 10, 64); err != nil {
	// 	log.Fatalf("Invalid MAIL_CLIENT_SALES_PORT: %v", err)
	// } else {
	// 	EnvData.MAIL_CLIENT_SALES_PORT = int(port)
	// }
	// EnvData.MAIL_CLIENT_SALES_USERNAME = getEnvKeyOrPanic("MAIL_CLIENT_SALES_USERNAME")
	// EnvData.MAIL_CLIENT_SALES_PASSWORD = getEnvKeyOrPanic("MAIL_CLIENT_SALES_PASSWORD")

	log.Println("Setting up env variables: done")
}

func getEnvKeyOrPanic(key string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		log.Fatal("Error loading ", key)
	}
	return val
}

func Log(v ...any) {
	log.Println(v...)
	// if EnvData.Debug {
	// 	log.Println("DEBUG: ", v)
	// }
}

func LogErr(e error) {
	log.Println("ERROR: ", e)
	// if EnvData.Debug {
	// 	log.Println("DEBUG ERROR: ", e)
	// }
}
