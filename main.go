package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/template/django/v3"
)

func main() {
	engine := django.New("./templates", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Get("/", func(c fiber.Ctx) error {
		return c.Render("layouts/index", fiber.Map{
			"PageTitle": "Hello, World!",
		})
	})

	app.Listen(":3000")
}
