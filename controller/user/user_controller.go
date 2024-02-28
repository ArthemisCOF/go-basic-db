package usercontroller

import (
	"go-fiber/datamodel"
	userservice "go-fiber/services/user"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type regiterUserRequest struct {
	UserID          string  `json:"username" validate:"required"`
	Password        string  `json:"password" validate:"required"`
	ComfirmPassword string  `json:"comfirm_password" validate:"required,eqfield=Password"`
	Email           string  `json:"email" validate:"required"`
	Cash            float64 `json:"cash" validate:"required"`
}

func RegiterUser(c *fiber.Ctx) error {
	req := new(regiterUserRequest)
	err := c.BodyParser(req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	v := validator.New()

	err = v.Struct(req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}
	user := &datamodel.User{
		UserID:   req.UserID,
		Password: req.Password,
		Email:    req.Email,
		Cash:     req.Cash,
	}
	err = userservice.RegiterUser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(201).JSON("success")
}
