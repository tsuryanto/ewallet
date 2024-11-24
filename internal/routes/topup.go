package routes

import (
	handler "ewallet/internal/httphandler"
	"ewallet/internal/middleware"
	"ewallet/internal/repository"
	"ewallet/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

func TopupRoutes(app *fiber.App) {
	userBalanceRepo := repository.NewUserBalanceRepository()
	topupRepo := repository.NewTopupRepository()
	topUpUseCase := usecase.NewTopUpUseCase(userBalanceRepo, topupRepo)
	topUpHandler := handler.NewTopupHttpHandler(topUpUseCase)

	// Your code here
	app.Post("/topup", middleware.AuthMiddleware, topUpHandler.TopUpHandler)
}
