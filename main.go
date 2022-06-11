package main

import (
	"os"

	"github.com/EdricT07/workhours/database"
	"github.com/EdricT07/workhours/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	database.Connect()
	app := fiber.New()

	//default logger config
	app.Use(logger.New())
	// set cors
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	// set the routes of the webapp
	routes.Setup(app)

	app.Listen(":" + os.Getenv("APPPORT"))

}
