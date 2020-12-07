package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wborbajr/osservice/kontroller"
)

// OSServiceRoute - generate route path to API
func OSServiceRoute(route fiber.Router) {

	// Route
	route.Get("/:doc/:os", kontroller.GetAllOS)

}