package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	// TODO
	return c.SendString("Gets All Users")
}

func GetUser(c *fiber.Ctx) error {
	// TODO
	id := c.Params("id")
	return c.SendString(fmt.Sprintf("Get User ID %s", id))
}

func CreateUser(c *fiber.Ctx) error {
	// TODO
	return c.SendString("Create an User")
}

func UpdateUser(c *fiber.Ctx) error {
	// TODO
	id := c.Params("id")
	return c.SendString(fmt.Sprintf("Update User ID %s", id))
}

func DeleteUser(c *fiber.Ctx) error {
	// TODO
	id := c.Params("id")
	return c.SendString(fmt.Sprintf("Delete User ID %s", id))
}
