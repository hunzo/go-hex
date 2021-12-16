package handler

import (
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
