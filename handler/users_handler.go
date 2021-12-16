package handler

import (
	"fmt"
	"go-hex/service"
	"strconv"

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
	ret := []UsersResponse{}
	for _, v := range rs {
		ret = append(ret, UsersResponse{
			Firstname: v.Firstname,
			LastName:  v.LastName,
		})
	}
	return c.JSON(ret)
}

func (h UserHandler) DeleteUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if err := h.userSrv.DeleteUserById(i); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}
	return c.JSON(fiber.Map{
		"success": true,
	})
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
