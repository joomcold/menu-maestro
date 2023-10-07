package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/joomcold/go-next-docker/internal/app/router"
	"github.com/joomcold/go-next-docker/internal/database"
	"github.com/joomcold/go-next-docker/internal/initializers"
)

func init() {
	initializers.Environment()
	initializers.Postgresql()
	database.Migrations()
}

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	app.Use(encryptcookie.New(encryptcookie.Config{
		Key:    os.Getenv("COOKIE_KEY"),
		Except: []string{"csrf_token"},
	}))

	// app.Use(csrf.New(csrf.Config{
	// 	CookieName:     "csrf_token",
	// 	CookieSecure:   true,
	// 	CookieHTTPOnly: true,
	// 	ErrorHandler: func(c *fiber.Ctx, err error) error {
	// 		return c.JSON(fiber.Map{"csrf_token:": fiber.StatusForbidden})
	// 	},
	// }))

	app.Use(requestid.New())

	app.Use(logger.New(logger.Config{
		TimeZone: "Asia/Bangkok",
	}))

	// Assign routes
	router.SetupRoutes(app)

	err := app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}
