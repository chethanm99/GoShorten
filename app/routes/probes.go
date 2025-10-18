package routes

import (
	"github.com/gofiber/fiber/v2"
)

func LivenessHandler(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func ReadinessHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "ok",
	})
}
