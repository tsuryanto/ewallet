package handler

import (
	"ewallet/internal/httphandler/reqres"
	"ewallet/internal/httphandler/utils"
	"ewallet/internal/repository"
	"ewallet/internal/usecase"
	"log"

	"github.com/gofiber/fiber/v2"
)

type UserHttpHandler struct {
	userUseCase usecase.UserUsecase
}

func NewUserHttpHandler(
	userUseCase usecase.UserUsecase,
) UserHttpHandler {
	return UserHttpHandler{
		userUseCase: userUseCase,
	}
}

// RegisterHandler handles the registration request
func (h *UserHttpHandler) RegisterHandler(c *fiber.Ctx) error {
	userRepository := repository.NewUserRepository()
	userUseCase := usecase.NewUserUsecase(userRepository)

	var req reqres.RegisterRequest
	// Parse the request body
	if err := c.BodyParser(&req); err != nil {
		log.Printf("Error parsing request: %v", err)
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid input")
	}

	// Call use case to register the user
	user, err := userUseCase.RegisterUser(req.FirstName, req.LastName, req.PhoneNumber, req.Address, req.Pin)
	if err != nil {
		// Handle specific error cases
		if err.Error() == "phone number already registered" {
			return utils.SendError(c, fiber.StatusConflict, "Phone Number already registered")
		}
		// Default internal error
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to register user")
	}

	// Send successful response
	return utils.SendSuccess(c, fiber.StatusOK, user)
}
func (h *UserHttpHandler) LoginHandler(c *fiber.Ctx) error {
	// Parse the request body
	var loginReq reqres.LoginRequest
	if err := c.BodyParser(&loginReq); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid input")
	}

	// Call use case to login the user
	token, err := h.userUseCase.LoginUser(loginReq.Username, loginReq.Password)
	if err != nil {
		return utils.SendError(c, fiber.StatusOK, "Failed to register user")
	}

	return utils.SendSuccess(c, fiber.StatusOK, fiber.Map{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
	})
}
