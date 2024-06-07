package auth

import (
	"github.com/gofiber/fiber/v2"
	auth "github.com/ssofiica/test-task-gazprom/internal/usecase"
)

type Delivery struct {
	uc auth.UseCase
}

func NewDeliveryLayer(ucProps auth.UseCase) *Delivery {
	return &Delivery{
		uc: ucProps,
	}
}

func (repo *Delivery) SignUp(c *fiber.Ctx) error {
	return nil
}

func (repo *Delivery) SignIn(c *fiber.Ctx) error {
	return nil
}
