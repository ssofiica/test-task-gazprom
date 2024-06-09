package auth

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/satori/uuid"
	"github.com/ssofiica/test-task-gazprom/internal/entity"
	"github.com/ssofiica/test-task-gazprom/internal/entity/dto"
	"github.com/ssofiica/test-task-gazprom/internal/usecase/auth"
)

type Delivery struct {
	uc auth.UseCase
}

func NewDeliveryLayer(ucProps auth.UseCase) *Delivery {
	return &Delivery{
		uc: ucProps,
	}
}

func (d *Delivery) SignUp(c *fiber.Ctx) error {
	email := ""
	emailCtx := c.Locals("email")
	if emailCtx != nil {
		email = emailCtx.(string)
	}
	fmt.Println(email)
	if email != "" {
		return c.Status(fiber.StatusBadRequest).SendString("Вы уже зарегистрированы")
	}

	//body
	signupInfo := dto.SignUp{}
	if err := c.BodyParser(&signupInfo); err != nil {
		fmt.Println("Auth delivery, SignUp, err: ", err.Error())
		return c.SendStatus(400)
	}

	fmt.Println(signupInfo)

	sessionId := uuid.NewV4().String()
	session := entity.Session{
		Id:    sessionId,
		Email: signupInfo.Email,
	}
	//registration
	err := d.uc.SignUp(c.Context(), &signupInfo, &session)
	if err != nil {
		fmt.Println(err)
		c.Status(500)
		return c.JSON(map[string]string{"error": "Ошибка сервера"})
	}

	//setting cookie
	cookie := new(fiber.Cookie)
	cookie.Name = "session_id"
	cookie.Value = sessionId
	cookie.Expires = time.Now().Add(14 * 24 * time.Hour)
	c.Cookie(cookie)

	c.Status(200)
	return c.SendString("Вы зарегистрированы")
}

func (d *Delivery) SignIn(c *fiber.Ctx) error {
	return nil
}

// cookie := new(fiber.Cookie)
//   cookie.Name = "john"
//   cookie.Value = "doe"
//   cookie.Expires = time.Now().Add(24 * time.Hour)

//   // Set cookie
//   c.Cookie(cookie)

//c.Cookies("name")         // "john"
