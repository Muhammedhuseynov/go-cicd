package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Calculater(ctx *fiber.Ctx) error {
	num1, _ := strconv.Atoi(ctx.Query("n1"))
	num2, _ := strconv.Atoi(ctx.Query("n2"))
	if num1 == 0 || num2 == 0 {
		// ctx.Status(500)
		return ctx.SendStatus(500)
	}
	res := num1 + num2

	return ctx.JSON(fiber.Map{"res": res})
}
