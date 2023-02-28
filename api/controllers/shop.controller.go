package controllers

import (
	"api/database"
	"api/repositories"
	"api/types"
	"api/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func GetShops(c *fiber.Ctx) error {
	shops, err := repositories.GetShops()
	if (err != types.HttpResponse{}) {
		return c.JSON(err)
	}
	return c.JSON(shops)
}

func GetShop(c *fiber.Ctx) error {
	id := c.Params("id")
	shop, err := repositories.GetShop(id)
	if (err != types.HttpResponse{}) {
		return c.JSON(err)
	}

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
	// TODO
	return c.SendString("Create an Shop")
	//return c.SendStatus(200)
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
