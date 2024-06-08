package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/ssofiica/test-task-gazprom/config"
	authDelivery "github.com/ssofiica/test-task-gazprom/internal/delivery/auth"
	userDelivery "github.com/ssofiica/test-task-gazprom/internal/delivery/user"
	"github.com/ssofiica/test-task-gazprom/internal/middleware"
	authRepo "github.com/ssofiica/test-task-gazprom/internal/repository/auth"
	userRepo "github.com/ssofiica/test-task-gazprom/internal/repository/user"
	authUseCase "github.com/ssofiica/test-task-gazprom/internal/usecase/auth"
	userUseCase "github.com/ssofiica/test-task-gazprom/internal/usecase/user"
	"github.com/ssofiica/test-task-gazprom/pkg/connection"
)

func main() {
	cfg := config.NewConfig()

	db := connection.InitPostgres(cfg)
	redis := connection.InitRedis(cfg, cfg.Redis.DatabaseSession)

	aRepo := authRepo.NewRepoLayer(db, redis)
	authUC := authUseCase.NewUseCaseLayer(aRepo)
	authHandler := authDelivery.NewDeliveryLayer(authUC)

	uRepo := userRepo.NewRepoLayer(db)
	userUC := userUseCase.NewUseCaseLayer(uRepo)
	userHandler := userDelivery.NewDeliveryLayer(userUC)

	app := fiber.New()
	app.Use(cors.New())
	app.Use(middleware.AuthMiddleware(authUC, userUC))

	api := app.Group("/api/v1")
	api.Post("/signup", authHandler.SignUp)
	api.Post("/signin", authHandler.SignIn)
	api.Get("/user/all", userHandler.GetAll) // get all users
	// api.Get("/user/search", userHandler.Search) // get one user by name and surname
	// api.Post("/user/subscribe")                 // subscribe user
	// api.Post("/user/unsubscribe")               // unsubscribe user
	// api.Get("/user/notification")               // get today notification about bday

	log.Fatal(app.Listen(":8000"))
}
