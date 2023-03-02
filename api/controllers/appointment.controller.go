package controllers

import (
	"api/repositories"
	"api/types"
	"api/utils"
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func GetAppointments(c *fiber.Ctx) error {
	db := utils.GetLocal[*sql.DB](c, "db")
	appointments, err := repositories.GetAppointments(db)

	if err != nil {
		return c.JSON(utils.E404("Appointments not found", err))
	}
	return c.JSON(appointments)
}

func GetAppointment(c *fiber.Ctx) error {
	id := c.Params("id")

	db := utils.GetLocal[*sql.DB](c, "db")
	appointment, err := repositories.GetAppointment(db, id)

	if err != nil {
		return c.JSON(utils.E404("Appointment not found", err))
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

	db := utils.GetLocal[*sql.DB](c, "db")
	appointmentId, err := repositories.CreateAppointment(db, appointment)

	if err != nil {
		return c.JSON(utils.E400("Bad request :\n"+err.Error(), err))
	}

	successMessage := fmt.Sprintf("Appointment %d created successfully", appointmentId)
	return c.JSON(types.HttpResponse{Status: 1, Message: successMessage, HttpCode: 200})
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

	db := utils.GetLocal[*sql.DB](c, "db")
	numberOfAffectedRows, err := repositories.UpdateAppointment(db, appointment, id)

	if err != nil {
		return c.JSON(utils.E503("Error while updating appointment", err))
	}

	if numberOfAffectedRows == 0 {
		return c.JSON(utils.E404("Appointment not found", nil))
	}

	successMessage := fmt.Sprintf("Appointment %d updated successfully", numberOfAffectedRows)
	return c.JSON(types.HttpResponse{Status: 1, Message: successMessage, HttpCode: 200})
}

func DeleteAppointment(c *fiber.Ctx) error {
	id := c.Params("id")

	db := utils.GetLocal[*sql.DB](c, "db")
	rowsAffected, err := repositories.DeleteAppointment(db, id)

	if err != nil {
		return c.JSON(utils.E503("Error while deleting appointment", err))
	}

	if rowsAffected == 0 {
		return c.JSON(utils.E404("Appointment not found", nil))
	}

	successMessage := fmt.Sprintf("Appointment %d deleted successfully", rowsAffected)
	return c.JSON(types.HttpResponse{Status: 1, Message: successMessage, HttpCode: 200})
}
