package user

import (
	"fmt"

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
	users, err := d.uc.GetAll(c.Context())
	if err != nil {
		fmt.Println("User delivery, GetAll, err: ", err.Error())
		return c.SendStatus(500)
	}
	err = c.JSON(users)
	if err != nil {
		fmt.Println("User delivery, GetAll, error in marshaling: ", err.Error())
		return c.SendStatus(500)
	}
	return c.SendStatus(200)
}

func (d *Delivery) Search(c *fiber.Ctx) error {
	return nil
}
