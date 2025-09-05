package utils

import (
	"io/fs"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type IEnvData struct {
	Debug bool

	API_PORT             string
	DB_CONNECTION_STRING string

	AUTH_SECRET       string
	AUTH_SECRET_BYTES []byte

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
		log.Fatal("Error determening GO_ENV (", os.Getenv("GO_ENV"), ")")
	}

	EnvData.API_PORT = getEnvKeyOrPanic("API_PORT")
	EnvData.DB_CONNECTION_STRING = getEnvKeyOrPanic("DB_CONNECTION_STRING")

	EnvData.AUTH_SECRET = getEnvKeyOrPanic("AUTH_SECRET")
	EnvData.AUTH_SECRET_BYTES = []byte(EnvData.AUTH_SECRET)

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
