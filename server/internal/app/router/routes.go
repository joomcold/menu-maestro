package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joomcold/go-next-docker/internal/app/controllers"
)

func SetupRoutes(app *fiber.App) {
	auth := app.Group("/")
	auth.Post("/login", controllers.Login)
	auth.Post("/logout", controllers.Logout)

	user := app.Group("/")
	profile := "/profile"

	user.Post("/register", controllers.Register)
	user.Get(profile, controllers.Profile)
	user.Post(profile, controllers.UpdateUser)
	user.Delete(profile, controllers.CancelUser)
}
