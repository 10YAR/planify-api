package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	// TODO
	return c.SendString("Login")
}

func Register(c *fiber.Ctx) error {
	// TODO
	return c.SendString("Register")
}
