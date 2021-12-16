package main

import (
	"go-hex/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	routes.SetupRouters(app)

	app.Listen(":8080")

}
