package main

import (
	"api/Controllers"
	"github.com/gofiber/fiber/v2"
)

func DefineRoutes() *fiber.App {
	app := GetFiberApp()

	// Welcome
	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("Welcome to the Planify API - 1.0")
		if err != nil {
			return err
		}
		return nil
	})

	// Appointments
	appointments := app.Group("/appointments")
	appointments.Get("/", Controllers.GetAppointments)

	return app
}
