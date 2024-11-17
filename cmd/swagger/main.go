package main

import (
	"log"

	"github.com/arsmn/fiber-swagger/v2" // Fiber Swagger UI
	"github.com/gofiber/fiber/v2"
)

// @title CNS Eljur Unified API
// @version 1.0
// @description Unified Swagger documentation for all modules (auth, admin, constructor)
// @contact.name API Support
// @contact.email support@example.com
// @host localhost:8080
// @BasePath /
func main() {
	app := fiber.New()

	// Маршрут для Swagger UI
	app.Get("/swagger/*", swagger.HandlerDefault)

	log.Println("[INFO] Starting Swagger UI at http://localhost:8080/swagger")
	log.Fatal(app.Listen(":8080"))
}
