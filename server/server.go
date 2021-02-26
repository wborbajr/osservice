package server

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/pkg/errors"
	"github.com/wborbajr/osservice/router"
)

var port string

func init() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error reading .env file: ", err)
	// }

	port = os.Getenv("APP_PORT")
}

// SetupApp - Create GoFiber app
func SetupApp() {

	l := log.New(os.Stdout, "OSService", log.LstdFlags)

	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if _, ok := err.(*fiber.Error); ok {
				return errors.New("Fiber Error")
			}
			return errors.New("Managed Error")
		},
		Concurrency:          256 * 1024,
		WriteTimeout:         10 * time.Second,
		ReadTimeout:          10 * time.Second,
		IdleTimeout:          10 * time.Second,
		ErrorLog:
		BodyLimit:            4 * 1024 * 1024,
		CompressedFileSuffix: ".fiber.gz",
	})

	app.Use(limiter.New(limiter.Config{
		Next:       nil,
		Max:        1000,
		Expiration: 1 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			c.Set("Content-Type", "application/json")
			return c.Status(429).SendString(`{"message":"Too much request #blocked"}`)
		},
	}))

	app.Use(recover.New())

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
	// setup routes
	setupRoutes(app)

	// fmt.Println(getOutboundIP())

	log.Printf("Server up and running: http://127.0.0.1:%s", port)
	// log.Fatal(app.Server().ListenAndServeTLS(":"+sslport, "./certs/server.crt", "./certs/server.key"))
	log.Fatal(app.Listen(":" + port))
}

// func getOutboundIP() net.IP {
// 	conn, err := net.Dial("udp", "8.8.8.8:80")
// 	if err != nil {
// 			log.Fatal(err)
// 	}
// 	defer conn.Close()

// 	localAddr := conn.LocalAddr().(*net.UDPAddr)

// 	return localAddr.IP
// }

func setupRoutes(app *fiber.App) {
	// give response when at /
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
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
