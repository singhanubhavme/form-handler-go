package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/singhanubhavme/form-handler-go/controllers"
	"github.com/singhanubhavme/form-handler-go/middlewares"
)

func FormRoute(app fiber.Router) {
	app.Post("/create", middlewares.Protected(), controllers.CreateForm)
	app.Post("/:user/:formid", controllers.SubmitForm)
}
