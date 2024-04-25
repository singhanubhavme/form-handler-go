package routes

import (
	"github.com/singhanubhavme/form-handler-go/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app fiber.Router) {
	app.Post("/register", controllers.RegisterUser)
	app.Post("/login", controllers.LoginUser)
}
