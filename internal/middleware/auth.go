package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ssofiica/test-task-gazprom/internal/usecase/auth"
	"github.com/ssofiica/test-task-gazprom/internal/usecase/user"
)

func AuthMiddleware(ucAuth auth.UseCase, ucUser user.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sessionId := c.Cookies("session_id")
		if sessionId == "" {
			fmt.Println("no session_id")
			return c.Next()
		}

		fmt.Println(sessionId)

		email, err := ucAuth.GetSessionValue(c.Context(), sessionId)
		if err != nil {
			fmt.Println("authmiddlware, getSessionValue ", err.Error())
			return c.Next()
		}

		user, err := ucUser.GetByEmail(c.Context(), email)
		if err != nil {
			fmt.Println("authmiddlware, getByEmail ", err.Error())
			return c.Next()
		}
		if user != nil {
			c.Locals("email", user.Email)
			fmt.Println(user.Email)
		}
		return c.Next()
	}
}
