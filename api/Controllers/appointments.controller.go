package Controllers

import "github.com/gofiber/fiber/v2"

func GetAppointments(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}
