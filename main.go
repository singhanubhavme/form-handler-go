package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/singhanubhavme/form-handler-go/configs"
	"github.com/singhanubhavme/form-handler-go/routes"
)

const port string = ":3001"

type response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    interface{}
}

func main() {
	app := fiber.New()
	configs.InitDb()
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(response{
			Status:  "success",
			Message: "success",
		})
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "*",
	}))

	apiUser := app.Group("/api/user")
	routes.UserRoute(apiUser)

	apiForm := app.Group("/api/form")
	routes.FormRoute(apiForm)

	log.Fatal(app.Listen(port))
}
