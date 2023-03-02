package controllers

import (
	"api/repositories"
	"api/types"
	"api/utils"
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func GetShops(c *fiber.Ctx) error {
	db := utils.GetLocal[*sql.DB](c, "db")
	shops, err := repositories.GetShops(db)

	if err != nil {
		return c.JSON(utils.E404("Shops not found", err))
	}

	return c.JSON(shops)
}

func GetShop(c *fiber.Ctx) error {
	id := c.Params("id")

	db := utils.GetLocal[*sql.DB](c, "db")
	shop, err := repositories.GetShop(db, id)

	if err != nil {
		return c.JSON(utils.E404("Shop not found", err))
	}

	availabilities, _ := repositories.GetShopAvailabilities(db, id)
	appointments, _ := repositories.GetShopAppointments(db, id)
	availabilitiesWithTimeSlots := utils.GenerateTimeSlotsOfAShop(availabilities)

	ShopInfoWithAvailabilityAppointmentsTimeSlots := types.ShopInfosAvailabilitiesAndAppointments{
		ShopInfos:      shop.ShopInfos,
		Availabilities: availabilitiesWithTimeSlots,
		Appointments:   appointments,
	}

	return c.JSON(ShopInfoWithAvailabilityAppointmentsTimeSlots)
}

//func GetShopAvailabilities(c *fiber.Ctx) error {
//	id := c.Params("id")
//	db := utils.GetLocal[*sql.DB](c, "db")
//	shopAvailabilities, errShopAvailabilities := repositories.GetShopAvailabilities(db, id)
//
//	if errShopAvailabilities != nil {
//		return c.JSON(utils.E404("ShopAvailabilities not found", errShopAvailabilities))
//	}
//
//	return c.JSON(shopAvailabilities)
//}

func GetShopAppointments(c *fiber.Ctx) error {
	id := c.Params("shopId")
	db := utils.GetLocal[*sql.DB](c, "db")
	shopAppointments, errShopAppointments := repositories.GetShopAppointments(db, id)

	if errShopAppointments != nil {
		return c.JSON(utils.E404("ShopAppointments not found", errShopAppointments))
	}

	return c.JSON(shopAppointments)
}

func GetShopsByUserId(c *fiber.Ctx) error {
	userId := c.Params("userId")

	db := utils.GetLocal[*sql.DB](c, "db")
	shops, err := repositories.GetShopsByUserId(db, userId)

	if err != nil {
		return c.JSON(utils.E404("Shops not found", err))
	}

	return c.JSON(shops)
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

	shopId, err := repositories.CreateShop(db, shop)
	if err != nil {
		return c.JSON(utils.E400("Bad request :\n"+err.Error(), err))
	}

	successMessage := fmt.Sprintf("Shop %d created successfully", shopId)
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

	_, err := repositories.UpdateShop(db, shop, id)
	if err != nil {
		return c.JSON(utils.E400("Bad request :\n"+err.Error(), err))
	}

	successMessage := fmt.Sprintf("Shop %s updated successfully", id)
	return c.JSON(types.HttpResponse{Status: 1, Message: successMessage, HttpCode: 200})
}

func DeleteShop(c *fiber.Ctx) error {
	id := c.Params("id")
	db := utils.GetLocal[*sql.DB](c, "db")

	_, err := repositories.DeleteShop(db, id)
	if err != nil {
		return c.JSON(utils.E400("Bad request :\n"+err.Error(), err))
	}

	successMessage := fmt.Sprintf("Shop %s deleted successfully", id)
	return c.JSON(types.HttpResponse{Status: 1, Message: successMessage, HttpCode: 200})
}
