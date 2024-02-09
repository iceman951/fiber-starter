package main

import (
	"fiber-starter/api"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	app := fiber.New()

	app.Use(requestid.New())

	apiGroup := app.Group("/api")

	apiGroup.Get("/greet", api.Greet)
	apiGroup.Post("/add", api.AddHandler)
	apiGroup.Post("/subtract", api.SubtractHandler)
	apiGroup.Post("/multiply", api.MultiplyHandler)
	apiGroup.Post("/divide", api.DivideHandler)
	apiGroup.Post("/is-odd", api.IsOddHandler)

	log.Fatal(app.Listen(":3000"))
}
