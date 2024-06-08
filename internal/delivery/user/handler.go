package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ssofiica/test-task-gazprom/internal/usecase/user"
)

type Delivery struct {
	uc user.UseCase
}

func NewDeliveryLayer(ucProps user.UseCase) *Delivery {
	return &Delivery{
		uc: ucProps,
	}
}

func (d *Delivery) GetAll(c *fiber.Ctx) error {
	return nil
}

func (d *Delivery) Search(c *fiber.Ctx) error {
	return nil
}
