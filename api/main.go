package main

import (
	"log"
	"os"

	"github.com/TDiblik/project-template/api/database"
	"github.com/TDiblik/project-template/api/router"
	"github.com/TDiblik/project-template/api/utils"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func main() {
	if fiber.IsChild() {
		log.Printf("[%d] Child: \n", os.Getppid())
	} else {
		log.Printf("[%d] Master: \n", os.Getppid())
	}
	log.Println("Initializing the Template API server")

	utils.SetupENV()
	utils.SetupValidator()
	// utils.SetupCronJobs()

	log.Println("Checking database connectivity: start")
	db, err := database.CreateConnection()
	if err != nil {
		log.Fatalln("Unable to connect to a database", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalln("Unable to ping a database", err)
	}
	log.Println("Checking database connectivity: done")

	// 	images_path := filepath.Join(utils.EnvData.IMAGES_PATH, db_name)
	// 	if err := os.MkdirAll(images_path, utils.FoldrePerms); err != nil {
	// 		log.Fatal("Error ensuring images path (\"", images_path, "\") for db (\"", db_name, "\"): ", err)
	// 	}
	// }
	// log.Println("Checking every database connectivity: done")

	log.Println("Setting up the app: start")
	app := fiber.New(fiber.Config{
		AppName:       "Template APP",
		CaseSensitive: true,
		BodyLimit:     15 * 1024 * 1024, // 15mb
	})
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())
	log.Println("Setting up the app: done")

	log.Println("Setting up the routes: start")
	router.SetupRoutes(app)
	log.Println("Setting up the routes: done")

	log.Println("Initialization completed")
	log.Fatal(app.Listen(":"+utils.EnvData.API_PORT, fiber.ListenConfig{
		EnablePrefork:     !utils.EnvData.Debug,
		EnablePrintRoutes: false,
	}))
}
