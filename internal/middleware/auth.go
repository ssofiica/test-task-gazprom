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
			return c.Next()
		}
		email, err := ucAuth.GetSessionValue(c.Context(), sessionId)
		if err != nil {
			fmt.Println(err)
			return c.Next()
		}

		user, err := ucUser.GetByEmail(c.Context(), email)
		if err != nil {
			fmt.Println(err)
			return c.Next()
		}
		if user != nil {
			c.Locals("email", user.Email)
		}
		return c.Next()
	}
}
