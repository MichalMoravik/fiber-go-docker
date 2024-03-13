package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func hiHandler(c *fiber.Ctx) error {
	return c.SendString("Yo, World!")
}

func main() {
	app := fiber.New()

	app.Get("/yo", hiHandler)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
