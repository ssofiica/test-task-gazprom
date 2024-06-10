package user

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ssofiica/test-task-gazprom/internal/usecase/user"
	"github.com/ssofiica/test-task-gazprom/pkg/myerrors"
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
		return c.Status(500).JSON(fiber.Map{"error": myerrors.InternalServer.Error()})
	}
	err = c.JSON(users)
	if err != nil {
		fmt.Println("User delivery, GetAll, error in marshaling: ", err.Error())
		return c.Status(500).JSON(fiber.Map{"error": myerrors.InternalServer.Error()})
	}
	return c.SendStatus(200)
}

func (d *Delivery) Search(c *fiber.Ctx) error {
	search := NameAndSurname{}
	if err := c.BodyParser(&search); err != nil {
		fmt.Println("User delivery, Search, err: ", err.Error())
		return c.SendStatus(400)
	}

	users, err := d.uc.Search(c.Context(), search.Name, search.Surname)
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(500).JSON(fiber.Map{"error": myerrors.InternalServer.Error()})
	}

	err = c.JSON(users)
	if err != nil {
		fmt.Println("User delivery, Search, error in marshaling: ", err.Error())
		return c.Status(500).JSON(fiber.Map{"error": myerrors.InternalServer.Error()})
	}
	return c.SendStatus(200)
}

func (d *Delivery) Subscribe(c *fiber.Ctx) error {
	email := ""
	emailCtx := c.Locals("email")
	if emailCtx != nil {
		email = emailCtx.(string)
	}
	if email == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": myerrors.Unauthorized.Error()})
	}

	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Param have to be number")
	}

	user, err := d.uc.GetByEmail(c.Context(), email)
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": myerrors.InternalServer.Error()})
	}

	err = d.uc.Subscribe(c.Context(), uint64(id), user.Id)
	if err != nil {
		fmt.Println(err.Error())
		if errors.Is(err, myerrors.NoSubcribeBdayUser) {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		if errors.Is(err, myerrors.NoUser) {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(500).JSON(fiber.Map{"error": myerrors.InternalServer.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"detail": "Подписка успешно оформлена"})
}

func (d *Delivery) UnSubscribe(c *fiber.Ctx) error {
	email := ""
	emailCtx := c.Locals("email")
	if emailCtx != nil {
		email = emailCtx.(string)
	}
	if email == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": myerrors.Unauthorized.Error()})
	}

	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Param have to be number")
	}

	user, err := d.uc.GetByEmail(c.Context(), email)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": myerrors.InternalServer.Error()})
	}

	err = d.uc.UnSubscribe(c.Context(), uint64(id), user.Id)
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, myerrors.NoUnsubcribeBdayUser) {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		if errors.Is(err, myerrors.NoUser) {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(500).JSON(fiber.Map{"error": myerrors.InternalServer.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"detail": "Подписка успешно отменена"})
}

func (d *Delivery) GetTodayBirthdayUsers(c *fiber.Ctx) error {
	email := ""
	emailCtx := c.Locals("email")
	if emailCtx != nil {
		email = emailCtx.(string)
	}
	if email == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": myerrors.Unauthorized.Error()})
	}

	user, err := d.uc.GetByEmail(c.Context(), email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": myerrors.InternalServer.Error()})
	}

	users, err := d.uc.GetTodayBirthdayUsers(c.Context(), uint64(user.Id))
	if err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(fiber.Map{"error": myerrors.InternalServer.Error()})
	}

	return c.Status(200).JSON(users)
}

type NameAndSurname struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}
