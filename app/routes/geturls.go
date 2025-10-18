package routes

import (
	"context"

	"github.com/chethanm99/go-url-shortner/database"
	"github.com/gofiber/fiber/v2"
)

func GetUrls(c *fiber.Ctx) error {
	ctx := context.Background()
	result := make(map[string]fiber.Map)

	iter := database.RDB.Scan(ctx, 0, "*", 0).Iterator()
	for iter.Next(ctx) {
		key := iter.Val()

		// Get URL
		longURL, err := database.RDB.Get(ctx, key).Result()
		if err != nil {
			continue
		}

		// Get expiry (TTL)
		ttl, _ := database.RDB.TTL(ctx, key).Result()
		result[key] = fiber.Map{
			"url":    longURL,
			"expiry": ttl.String(),
		}
	}

	if err := iter.Err(); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch URLs",
		})
	}

	return c.JSON(result)
}
