package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-admin/src/controllers"
	"go-admin/src/middleware"
)

func Setup(app *fiber.App) {
	api := app.Group("api")
	admin := api.Group("admin")
	admin.Post("/register", controllers.Register)
	admin.Post("/login", controllers.Login)
	adminAuthenticated := admin.Use(middleware.IsAuthenticated)
	adminAuthenticated.Get("/user", controllers.User)
	adminAuthenticated.Post("/logout", controllers.Logout)
}
