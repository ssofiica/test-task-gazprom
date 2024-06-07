package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	auth2 "github.com/ssofiica/test-task-gazprom/internal/delivery"
	auth "github.com/ssofiica/test-task-gazprom/internal/repository"
	auth1 "github.com/ssofiica/test-task-gazprom/internal/usecase"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	authRepo := auth.NewRepoLayer(nil, nil)
	authUC := auth1.NewUseCaseLayer(authRepo)
	authHandler := auth2.NewDeliveryLayer(authUC)

	app.Post("/api/v1/signup", authHandler.SignUp)
	app.Post("/api/v1/signin", authHandler.SignIn)
	app.Get("/api/v1/user/all")          // get all users
	app.Get("/api/v1/user/search")       // get one user by name and surname
	app.Get("/api/v1/user/search")       // get one user by name and surname
	app.Post("/api/v1/user/subscribe")   // subscribe user
	app.Post("/api/v1/user/unsubscribe") // unsubscribe user
	app.Get("/api/v1/user/notification") // get today notification about bday
}
