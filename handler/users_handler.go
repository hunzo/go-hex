package handler

import (
	"fmt"
	"go-hex/service"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userSrv service.UserSevice
}

func NewHandler(h service.UserSevice) UserHandler {
	return UserHandler{userSrv: h}
}

func (h UserHandler) GetUsers(c *fiber.Ctx) error {
	rs, err := h.userSrv.GetUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.JSON(rs)
}

func (h UserHandler) CreateUser(c *fiber.Ctx) error {
	req := UsersReq{}
	if err := c.BodyParser(&req); err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	uc := service.UserCreate{
		Firstname: req.Firstname,
		LastName:  req.LastName,
	}

	fmt.Printf("handler: %v\n", uc)

	if err := h.userSrv.CreateUser(uc); err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}
	return c.SendStatus(fiber.StatusCreated)
}
