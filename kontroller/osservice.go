package kontroller

import "github.com/gofiber/fiber/v2"

// GetAllOS - retrieve customer order service
func GetAllOS(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"success": false,
		"message": "OS not found",
	})

}