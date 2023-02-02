package main

import (
	"api/controllers"
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

	// Auth
	auth := app.Group("/auth")
	auth.Post("/login", controllers.Login)
	auth.Post("/register", controllers.Register)

	// Appointments
	appointments := app.Group("/appointments")
	appointments.Get("/", controllers.GetAppointments)
	appointments.Get("/:id", controllers.GetAppointment)
	appointments.Post("/", controllers.CreateAppointment)
	appointments.Patch("/:id", controllers.UpdateAppointment)
	appointments.Delete("/:id", controllers.DeleteAppointment)

	// Users
	users := app.Group("/users")
	users.Get("/", controllers.GetUsers)
	users.Get("/:id", controllers.GetUser)
	users.Post("/", controllers.CreateUser)
	users.Patch("/:id", controllers.UpdateUser)
	users.Delete("/:id", controllers.DeleteUser)

	// Shops
	shops := app.Group("/shops")
	shops.Get("/", controllers.GetShops)
	shops.Get("/:id", controllers.GetShop)
	shops.Post("/", controllers.CreateShop)
	shops.Patch("/:id", controllers.UpdateShop)
	shops.Delete("/:id", controllers.DeleteShop)

	return app
}
