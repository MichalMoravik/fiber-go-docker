package main

import (
    "os"
    "log"

    "github.com/gofiber/fiber/v2"
)

func hiHandler(c *fiber.Ctx) error {
    return c.SendString("Yo, World!")
}

func main() {
    app := fiber.New()

    app.Post("/yo", hiHandler)

    log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}

