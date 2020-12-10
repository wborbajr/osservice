package kontroller

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// GetAllOS - retrieve customer order service
func GetAllOS(c *fiber.Ctx) error {

	// Parsing parameters
	paramDoc := c.Params("doc")
	paramOs := c.Params("os")

	log.Println(paramDoc)
	log.Println(paramOs)

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"success": false,
		"message": "OS not found",
	})

}