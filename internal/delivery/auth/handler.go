package auth

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/satori/uuid"
	"github.com/ssofiica/test-task-gazprom/internal/entity"
	"github.com/ssofiica/test-task-gazprom/internal/entity/dto"
	"github.com/ssofiica/test-task-gazprom/internal/usecase/auth"
	"github.com/ssofiica/test-task-gazprom/internal/usecase/user"
)

type Delivery struct {
	ucAuth auth.UseCase
	ucUser user.UseCase
}

func NewDeliveryLayer(ucaProps auth.UseCase, ucuProps user.UseCase) *Delivery {
	return &Delivery{
		ucAuth: ucaProps,
		ucUser: ucuProps,
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

	user, err := d.ucUser.GetByEmail(c.Context(), signupInfo.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Ошибка сервера")
	}
	if user != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Вы уже зарегистрированы")
	}

	sessionId := uuid.NewV4().String()
	session := entity.Session{
		Id:    sessionId,
		Email: signupInfo.Email,
	}
	//registration
	err = d.ucAuth.SignUp(c.Context(), &signupInfo, &session)
	if err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(map[string]string{"error": "Ошибка сервера"})
	}

	//setting cookie
	cookie := new(fiber.Cookie)
	cookie.Name = "session_id"
	cookie.Value = sessionId
	cookie.Expires = time.Now().Add(14 * 24 * time.Hour)
	c.Cookie(cookie)

	return c.Status(200).SendString("Вы зарегистрированы")
}

func (d *Delivery) SignIn(c *fiber.Ctx) error {
	email := ""
	emailCtx := c.Locals("email")
	if emailCtx != nil {
		email = emailCtx.(string)
	}
	fmt.Println(email)
	if email != "" {
		return c.Status(fiber.StatusBadRequest).SendString("Вы уже авторизированы")
	}

	//body
	signinInfo := dto.SignIn{}
	if err := c.BodyParser(&signinInfo); err != nil {
		fmt.Println("Auth delivery, SignIn, err: ", err.Error())
		return c.SendStatus(400)
	}

	fmt.Println(signinInfo)

	user, err := d.ucUser.GetByEmail(c.Context(), signinInfo.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Ошибка сервера")
	}
	if user == nil {
		return c.Status(fiber.StatusBadRequest).SendString("Неверный адрес почты")
	}

	sessionId := uuid.NewV4().String()
	session := entity.Session{
		Id:    sessionId,
		Email: signinInfo.Email,
	}

	err = d.ucAuth.SignIn(c.Context(), user, &signinInfo, &session)
	if err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(map[string]string{"error": "Ошибка сервера"})
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "session_id"
	cookie.Value = sessionId
	cookie.Expires = time.Now().Add(14 * 24 * time.Hour)
	c.Cookie(cookie)

	return c.Status(200).SendString("Вы успешно авторизированы")
}
