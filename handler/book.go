package handler

import (
	// "Victor-Garces/golang-api/router"
	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	return c.SendString("Hello, World👋!")
}