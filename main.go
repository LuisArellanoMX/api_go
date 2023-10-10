package main

import (
	"os"

	"github.com/LuisArellanoMX/api_go/common"
	"github.com/LuisArellanoMX/api_go/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	err := run()

	if err != nil {
		panic(err)
	}
}

func run() error {
	// Init env
	err := godotenv.Load()
	if err != nil {
		return err
	}

	// Init db
	err = common.InitDB()
	if err != nil {
		return err
	}

	// Defer closing db
	defer common.CloseDB()

	// Create app
	app := fiber.New()

	// Add basic configurations
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	// Add routes
	router.AddIcecreamGroup(app)

	// Start server
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}
	app.Listen(":" + port)

	return nil
}
