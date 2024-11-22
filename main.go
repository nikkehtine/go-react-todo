package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

const PORT = ":4040"

func main() {
	app := fiber.New()
	fmt.Println("Starting server on", PORT)
	app.Listen(PORT)
}
