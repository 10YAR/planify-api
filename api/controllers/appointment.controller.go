package controllers

import (
	"api/database"
	"api/types"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func GetAppointments(c *fiber.Ctx) error {
	// TODO
	db := database.Mysql()

	res, err := db.Query("SELECT * FROM appointments")
	if err != nil {
		return err
	}
	database.DeferClose(db)

	var appointments []types.Appointment
	for res.Next() {
		var appointment types.Appointment
		err := res.Scan(&appointment.ID, &appointment.CustomerName, &appointment.AppointmentDate, &appointment.StartTime, &appointment.EndTime, &appointment.Status, &appointment.ShopId)
		if err != nil {
			return err
		}
		appointments = append(appointments, appointment)
	}
	return c.JSON(appointments)
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
