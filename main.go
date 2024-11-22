package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

const PORT = ":4040"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "Hello there"})
	})

	fmt.Println("Starting server on", PORT)
	log.Fatal(app.Listen(PORT))
}
