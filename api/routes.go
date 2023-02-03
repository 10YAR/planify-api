package main

import (
	"api/controllers"
	"api/middlewares"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"os"
)

func DefineRoutes() *fiber.App {
	app := GetFiberApp()

	// Welcome (Public)
	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("Welcome to the Planify API - 1.0")
		if err != nil {
			return err
		}
		return nil
	})

	// Auth (Public)
	auth := app.Group("/auth")
	auth.Post("/", controllers.Login)
	auth.Post("/register", controllers.Register)

	// Appointments (Public)
	appointments := app.Group("/appointments")
	appointments.Get("/", controllers.GetAppointments)
	appointments.Get("/:id", controllers.GetAppointment)

	// Users (Public)
	users := app.Group("/users")

	// Shops (Public)
	shops := app.Group("/shops")
	shops.Get("/", controllers.GetShops)
	shops.Get("/:id", controllers.GetShop)

	// Authenticated Routes
	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("SECRET_KEY")),
	}))

	// Appointments (Authenticated)
	appointments.Post("/", middlewares.ProcessAuth, controllers.CreateAppointment)
	appointments.Patch("/:id", middlewares.ProcessAuth, controllers.UpdateAppointment)
	appointments.Delete("/:id", middlewares.ProcessAuth, controllers.DeleteAppointment)

	// Users (Authenticated)
	users.Get("/", middlewares.ProcessAuth, controllers.GetUsers)
	users.Get("/:id", middlewares.ProcessAuth, controllers.GetUser)
	users.Post("/", middlewares.ProcessAuth, controllers.CreateUser)
	users.Patch("/:id", middlewares.ProcessAuth, controllers.UpdateUser)
	users.Delete("/:id", middlewares.ProcessAuth, controllers.DeleteUser)

	// Shops (Authenticated)
	shops.Post("/", middlewares.ProcessAuth, controllers.CreateShop)
	shops.Patch("/:id", middlewares.ProcessAuth, controllers.UpdateShop)
	shops.Delete("/:id", middlewares.ProcessAuth, controllers.DeleteShop)

	return app
}
