package controllers

import (
	"api/database"
	"api/types"
	"api/utils"
	"github.com/gofiber/fiber/v2"
)

func GetAppointments(c *fiber.Ctx) error {
	res, err := database.DoQuery("SELECT * FROM appointment")
	if err != nil {
		return c.JSON(utils.E503("Error while getting appointments", err))
	}

	var appointments []types.Appointment
	for res.Next() {
		var appointment types.Appointment
		err := res.Scan(&appointment.ID, &appointment.CustomerName, &appointment.AppointmentDate, &appointment.StartTime, &appointment.EndTime, &appointment.Status, &appointment.ShopId)
		if err != nil {
			return c.JSON(utils.E503("Error while getting appointments", err))
		}

		appointments = append(appointments, appointment)
	}

	return c.JSON(appointments)
}

func GetAppointment(c *fiber.Ctx) error {
	id := c.Params("id")
	res, err := database.DoQuery("SELECT * FROM appointment WHERE id = ?", id)
	if err != nil {
		return c.JSON(utils.E503("Error while getting appointment", err))
	}

	var appointment types.Appointment
	for res.Next() {
		err := res.Scan(&appointment.ID, &appointment.CustomerName, &appointment.AppointmentDate, &appointment.StartTime, &appointment.EndTime, &appointment.Status, &appointment.ShopId)
		if err != nil {
			return c.JSON(utils.E503("Error while getting appointment", err))
		}
	}

	return c.JSON(appointment)
}

func CreateAppointment(c *fiber.Ctx) error {
	appointment := new(types.Appointment)

	if err := c.BodyParser(appointment); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := utils.ValidateStruct(*appointment)
	if errors != "" {
		return c.JSON(utils.E400("Bad request :\n"+errors, nil))
	}

	_, err := database.DoQuery("INSERT INTO appointment (customer_name, appointment_date, start_time, end_time, status, shop_id) VALUES (?, ?, ?, ?, ?, ?)", appointment.CustomerName, appointment.AppointmentDate, appointment.StartTime, appointment.EndTime, appointment.Status, appointment.ShopId)
	if err != nil {
		return c.JSON(utils.E503("Error while creating appointment", err))
	}

	return c.JSON(types.HttpResponse{Status: 1, Message: "Appointment created successfully", HttpCode: 200})
}

func UpdateAppointment(c *fiber.Ctx) error {
	id := c.Params("id")

	appointment := new(types.Appointment)

	if err := c.BodyParser(appointment); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := utils.ValidateStruct(*appointment)
	if errors != "" {
		return c.JSON(utils.E400("Bad request :\n"+errors, nil))
	}

	_, err := database.DoQuery("UPDATE appointment SET customer_name = ?, appointment_date = ?, start_time = ?, end_time = ?, status = ?, shop_id = ? WHERE id = ?", appointment.CustomerName, appointment.AppointmentDate, appointment.StartTime, appointment.EndTime, appointment.Status, appointment.ShopId, id)
	if err != nil {
		return c.JSON(utils.E503("Error while updating appointment", err))
	}

	return c.JSON(types.HttpResponse{Status: 1, Message: "Appointment updated successfully", HttpCode: 200})
}

func DeleteAppointment(c *fiber.Ctx) error {
	id := c.Params("id")

	_, err := database.DoQuery("DELETE FROM appointment WHERE id = ?", id)
	if err != nil {
		return c.JSON(utils.E503("Error while deleting appointment", err))
	}

	return c.JSON(types.HttpResponse{Status: 1, Message: "Appointment deleted successfully", HttpCode: 200})
}
