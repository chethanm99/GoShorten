package routes

import (
	"context"
	"os"
	"time"

	"github.com/asaskevich/govalidator"
	helper "github.com/chethanm99/go-url-shortner/api/helpers"
	"github.com/chethanm99/go-url-shortner/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func ShortenURL(c *fiber.Ctx) error {
	type request struct {
		URL         string        `json:"url"`
		CustomShort string        `json:"short"`
		Expiry      time.Duration `json:"expiry"`
	}

	body := new(request)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	// Validate input
	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid URL",
		})
	}

	if !helper.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"error": "cannot use this domain",
		})
	}

	// Enforce HTTP scheme
	body.URL = helper.EnforceHTTP(body.URL)

	// Generate ID
	var id string
	if body.CustomShort == "" {
		id = uuid.New().String()[:6]
	} else {
		id = body.CustomShort
		val, _ := database.RDB.Get(context.Background(), id).Result()
		if val != "" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Custom short URL already in use",
			})
		}
	}

	// Expiry
	if body.Expiry == 0 {
		body.Expiry = 24
	}

	err := database.RDB.Set(context.Background(), id, body.URL, body.Expiry*time.Hour).Err()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to save to database",
		})
	}

	resp := fiber.Map{
		"url":     body.URL,
		"short":   os.Getenv("DOMAIN") + "/" + id, // Correctly set the short URL
		"expiry":  body.Expiry,
		"created": time.Now().UTC(),
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}
