package routes

import (
	"context"

	"github.com/chethanm99/go-url-shortner/database"
	"github.com/gofiber/fiber/v2"
)

func DeleteURL(c *fiber.Ctx) error {
	id := c.Params("short")
	url, err := database.RDB.Del(context.Background(), id).Result()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to delete url",
		})
	}

	if url == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "short url not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "URL deleted successfully",
	})
}
