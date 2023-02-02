package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func GetShops(c *fiber.Ctx) error {
	// TODO
	return c.SendString("Gets All Shops")
}

func GetShop(c *fiber.Ctx) error {
	// TODO
	id := c.Params("id")
	return c.SendString(fmt.Sprintf("Get Shop ID %s", id))
}

func CreateShop(c *fiber.Ctx) error {
	// TODO
	return c.SendString("Create an Shop")
}

func UpdateShop(c *fiber.Ctx) error {
	// TODO
	id := c.Params("id")
	return c.SendString(fmt.Sprintf("Update Shop ID %s", id))
}

func DeleteShop(c *fiber.Ctx) error {
	// TODO
	id := c.Params("id")
	return c.SendString(fmt.Sprintf("Delete Shop ID %s", id))
}
