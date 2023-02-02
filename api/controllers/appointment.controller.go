package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func GetAppointments(c *fiber.Ctx) error {
	// TODO
	return c.SendString("Gets All Appointments")
}

func GetAppointment(c *fiber.Ctx) error {
	// TODO
	id := c.Params("id")
	return c.SendString(fmt.Sprintf("Get Appointment ID %s", id))
}

func CreateAppointment(c *fiber.Ctx) error {
	// TODO
	return c.SendString("Create an Appointment")
}

func UpdateAppointment(c *fiber.Ctx) error {
	// TODO
	id := c.Params("id")
	return c.SendString(fmt.Sprintf("Update Appointment ID %s", id))
}

func DeleteAppointment(c *fiber.Ctx) error {
	// TODO
	id := c.Params("id")
	return c.SendString(fmt.Sprintf("Delete Appointment ID %s", id))
}
