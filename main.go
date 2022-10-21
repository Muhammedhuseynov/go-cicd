package main

import (
	"app/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	router.Route(app)
	// http://localhost:5050/?n1=1&n2=3
	app.Listen(":5050")
}
