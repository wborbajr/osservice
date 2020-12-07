package main

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/wborbajr/osservice/router"
)

func setupApp() {
	app := fiber.New(fiber.Config{
		Concurrency:  	256 * 1024,
		WriteTimeout: 	10 * time.Second,
		ReadTimeout: 	10 * time.Second,
		IdleTimeout:	10 * time.Second,
		BodyLimit:		4 * 1024 * 1024,
		CompressedFileSuffix: ".fiber.gz",
	})

	app.Use(limiter.New(limiter.Config{
		Expiration: 10 * time.Second,
		Max:      14,
	}))

	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2000",
		TimeZone:   "America/Sao_Paulo",
		Output:     os.Stdout,
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, OPTIONS, PUT, DELETE, POST",
	}))

	log.Printf("Loading routes...")
	// router.SetupRoute(app)
	// setup routes
	setupRoutes(app)

	// port := os.Getenv("APP_PORT")
	port := "3001" // os.Getenv("APP_SSL_PORT")

	log.Printf( "Server up and running: http://127.0.0.1:%s", port)
	// log.Fatal(app.Server().ListenAndServeTLS(":"+sslport, "./certs/server.crt", "./certs/server.key"))
	log.Fatal(app.Listen(":"+port))

}

func setupRoutes(app *fiber.App) {
	// give response when at /
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success":  true,
			"message": "You are at the endpoint 😉",
		})
	})

	// api group
	api := app.Group("/api")

	// give response when at /api
	api.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the api endpoint 😉",
		})
	})

	// send todos route group to TodoRoutes of routes package
	router.OSServiceRoute(api.Group("/osservice"))
}


func main() {
	log.Printf("Starting the web server...")
	setupApp()

}