package main

import (
	"go-hex/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	payloadRepo := repository.New(db)

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": "status",
		})
	})

	app.Get("/getall", func(c *fiber.Ctx) error {
		d, err := payloadRepo.GetData()
		if err != nil {
			return c.SendStatus(fiber.StatusUnprocessableEntity)
		}
		return c.JSON(d)
	})

	app.Get("/create", func(c *fiber.Ctx) error {

		if err != nil {
			panic(err)
		}

		id, err := uuid.NewRandom()
		if err != nil {
			return c.SendStatus(fiber.StatusUnprocessableEntity)
		}

		p := repository.Payload{
			Uid:         id,
			FileName:    "fileName",
			FileType:    "application/json",
			FileContent: "<h1>TEST</h1>",
		}

		if err := payloadRepo.CreateData(p); err != nil {
			return c.JSON(err)
		}

		return c.Status(fiber.StatusCreated).JSON(p)
	})

	app.Listen(":8080")

}
