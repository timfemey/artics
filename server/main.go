package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	app.Post("/newarticle")

	app.Post("/comment")

	app.Get("/article")

	app.Get("/viewcomms")

	log.Fatal(app.Listen(":5000"))
}
