package router

import (
	"app/controller"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	app.Get("/", controller.Calculater)
}
