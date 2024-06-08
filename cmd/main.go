package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/ssofiica/test-task-gazprom/config"
	authDelivery "github.com/ssofiica/test-task-gazprom/internal/delivery"
	authRepo "github.com/ssofiica/test-task-gazprom/internal/repository"
	authUseCase "github.com/ssofiica/test-task-gazprom/internal/usecase"
	"github.com/ssofiica/test-task-gazprom/pkg/connection"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	cfg := config.NewConfig()

	db := connection.InitPostgres(cfg)
	redis := connection.InitRedis(cfg, cfg.Redis.DatabaseSession)

	aRepo := authRepo.NewRepoLayer(db, redis)
	authUC := authUseCase.NewUseCaseLayer(aRepo)
	authHandler := authDelivery.NewDeliveryLayer(authUC)

	api := app.Group("/api/v1")
	api.Post("/signup", authHandler.SignUp)
	api.Post("/signin", authHandler.SignIn)
	api.Get("/user/all")          // get all users
	api.Get("/user/search")       // get one user by name and surname
	api.Post("/user/subscribe")   // subscribe user
	api.Post("/user/unsubscribe") // unsubscribe user
	api.Get("/user/notification") // get today notification about bday

	log.Fatal(app.Listen(":8000"))
}
