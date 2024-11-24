package utils

import "github.com/gofiber/fiber/v2"

// SendError sends a standardized error response
func SendError(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"status":  "ERROR",
		"message": message,
	})
}

// SendSuccess sends a standardized success response
func SendSuccess(c *fiber.Ctx, statusCode int, data interface{}) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"status": "SUCCESS",
		"result": data,
	})
}
