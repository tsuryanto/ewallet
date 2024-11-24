package routes

import (
	handler "ewallet/internal/httphandler"
	"ewallet/internal/repository"
	"ewallet/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

// RegisterRoutes will register all user-related routes
func RegisterRoutes(app *fiber.App) {
	userRepository := repository.NewUserRepository()
	userUserCase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHttpHandler(userUserCase)

	app.Post("/register", userHandler.RegisterHandler)
	app.Post("/login", userHandler.LoginHandler)
}
