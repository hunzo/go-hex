package routes

import (
	"go-hex/handler"
	"go-hex/repository"
	"go-hex/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupRouters(r *fiber.App) {
	db, err := gorm.Open(sqlite.Open("./db/test.sqlite"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUsers(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewHandler(userService)

	r.Get("/", userHandler.GetUsers)
	r.Post("/create", userHandler.CreateUser)
	r.Get("/delete/:id", userHandler.DeleteUserById)

}
