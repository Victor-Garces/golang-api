package router

import (
	"Victor-Garces/golang-api/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	handler := logger.New()
	api := app.Group("/api", handler)
	BooksEndpoints(api, handler)
}

func BooksEndpoints(app fiber.Router, h fiber.Handler) {
	book := app.Group("/book", h)
	book.Get("/", handler.GetBooks)
}