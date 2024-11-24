package handler

import (
	"ewallet/internal/httphandler/reqres"
	"ewallet/internal/httphandler/utils"
	"ewallet/internal/usecase"
	"log"

	"github.com/gofiber/fiber/v2"
)

type TopupHttpHandler struct {
	topupUseCase usecase.TopUpUseCase
}

func NewTopupHttpHandler(
	topupUseCase usecase.TopUpUseCase,
) TopupHttpHandler {
	return TopupHttpHandler{
		topupUseCase: topupUseCase,
	}
}

// TopUpHandler handles the top-up request
func (h *TopupHttpHandler) TopUpHandler(c *fiber.Ctx) error {
	// Parse the request body
	var req reqres.TopUpRequest
	if err := c.BodyParser(&req); err != nil || req.Amount <= 0 {
		log.Printf("Invalid top-up request: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid input",
		})
	}

	// Get the user ID from the context
	userID := c.Locals("user_id").(string)

	// Process the top-up
	topupUsecase, err := h.topupUseCase.TopUp(userID, req.Amount)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to process top-up")
	}

	return utils.SendSuccess(c, fiber.StatusOK, reqres.TopUpResponse{
		TopUpID:       topupUsecase.TopUpID,
		AmountTopUp:   topupUsecase.AmountTopUp,
		BalanceBefore: topupUsecase.BalanceBefore,
		BalanceAfter:  topupUsecase.BalanceAfter,
		CreatedDate:   topupUsecase.CreatedDate,
	})
}
