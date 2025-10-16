package main

import (
	"log"
	"os"

	"github.com/chethanm99/go-url-shortner/routes"

	"github.com/chethanm99/go-url-shortner/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recovermw "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

// setupRoutes registers app routes
func setupRoutes(app *fiber.App) {
	app.Get("/healthz", routes.LivenessHandler)
	app.Get("/readyz", routes.ReadinessHandler)

	app.Post("/api/v1", routes.ShortenURL)
	app.Get("/urls", routes.GetUrls)
	app.Get("/:short", routes.ResolveURL)
	app.Delete("/api/v1/:short", routes.DeleteURL)
}

func main() {
	// Load env
	_ = godotenv.Load()

	if err := database.Connect(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = ":8080"
	}

	app := fiber.New()
	app.Use(logger.New())
	app.Use(recovermw.New())

	setupRoutes(app)

	log.Fatal(app.Listen(port))
}
