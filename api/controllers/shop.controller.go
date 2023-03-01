package controllers

import (
	"api/database"
	"api/repositories"
	"api/types"
	"api/utils"
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func GetShops(c *fiber.Ctx) error {
	shops := repositories.GetShops()
	return c.JSON(shops)
}

func GetShop(c *fiber.Ctx) error {
	id := c.Params("id")

	db := utils.GetLocal[*sql.DB](c, "db")
	shop := repositories.GetShop(db, id)

	availabilities, _ := GetAvailabilitiesOfAShop(id)
	appointments, _ := GetAppointmentsOfAShop(id)
	availabilitiesWithTimeSlots := utils.GenerateTimeSlotsOfAShop(availabilities)

	ShopInfoWithAvailabilityAppointmentsTimeSlots := types.ShopInfosAvailabilitiesAndAppointments{
		ShopInfos:      shop.ShopInfos,
		Availabilities: availabilitiesWithTimeSlots,
		Appointments:   appointments,
	}

	return c.JSON(ShopInfoWithAvailabilityAppointmentsTimeSlots)
}

func GetAvailabilitiesOfAShop(id string) ([]types.ShopAvailability, types.HttpResponse) {
	resShopAvailability, errShopAvailability := database.DoQuery("SELECT shop_availability.day_of_week, shop_availability.duration, shop_availability.start_time, shop_availability.end_time FROM shops INNER JOIN shop_availability ON shops.id = shop_availability.shop_id WHERE shops.id = ? ", id)
	var errorMessage types.HttpResponse
	if errShopAvailability != nil {
		errorMessage = utils.E503("Error while getting ShopAvailability from database", errShopAvailability)
	}

	var availabilities []types.ShopAvailability
	for resShopAvailability.Next() {
		var shopAvailability types.ShopAvailability
		err := resShopAvailability.Scan(&shopAvailability.DayOfWeek, &shopAvailability.Duration, &shopAvailability.StartTime, &shopAvailability.EndTime)
		if err != nil {
			errorMessage = utils.E503("Error while getting availabilities attributes", err)
		}
		availabilities = append(availabilities, shopAvailability)
	}
	return availabilities, errorMessage
}

func GetAppointmentsOfAShop(id string) ([]types.AppointmentDateTimeInfos, types.HttpResponse) {
	resShopAppointments, errShopAppointments := database.DoQuery("SELECT appointment_date, appointment_time, appointment_date_time FROM appointments WHERE shop_id = ? ", id)
	var errorMessage types.HttpResponse
	if errShopAppointments != nil {
		errorMessage = utils.E503("Error while getting appointments of current shop from database", errShopAppointments)
	}

	var appointments []types.AppointmentDateTimeInfos
	for resShopAppointments.Next() {
		var appointment types.AppointmentDateTimeInfos
		err := resShopAppointments.Scan(&appointment.AppointmentDate, &appointment.AppointmentTime, &appointment.AppointmentDateTime)
		if err != nil {
			errorMessage = utils.E503("Error while getting appointments attributes", err)
		}
		appointments = append(appointments, appointment)
	}
	return appointments, errorMessage
}

func CreateShop(c *fiber.Ctx) error {
	db := utils.GetLocal[*sql.DB](c, "db")
	shop := new(types.Shop)

	if err := c.BodyParser(shop); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := utils.ValidateStruct(*shop)
	if errors != "" {
		return c.JSON(utils.E400("Bad request :\n"+errors, nil))
	}

	shopId := repositories.CreateShop(db, shop)
	successMessage := fmt.Sprintf("Shop %s created successfully", shopId)
	return c.JSON(types.HttpResponse{Status: 1, Message: successMessage, HttpCode: 200})
}

func UpdateShop(c *fiber.Ctx) error {
	id := c.Params("id")
	db := utils.GetLocal[*sql.DB](c, "db")
	shop := new(types.Shop)

	if err := c.BodyParser(shop); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := utils.ValidateStruct(*shop)
	if errors != "" {
		return c.JSON(utils.E400("Bad request :\n"+errors, nil))
	}

	numberOfAffectedRows := repositories.UpdateShop(db, shop, id)
	successMessage := fmt.Sprintf("%d rows updated", numberOfAffectedRows)
	return c.JSON(types.HttpResponse{Status: 1, Message: successMessage, HttpCode: 200})
}

func DeleteShop(c *fiber.Ctx) error {
	id := c.Params("id")
	db := utils.GetLocal[*sql.DB](c, "db")

	numberOfAffectedRows, err := repositories.DeleteShop(db, id)
	if err != nil {
		return c.JSON(utils.E400("Bad request :\n"+err.Error(), err))
	}
	successMessage := fmt.Sprintf("%d rows deleted", numberOfAffectedRows)
	return c.JSON(types.HttpResponse{Status: 1, Message: successMessage, HttpCode: 200})
}
