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

func GetUsers(c *fiber.Ctx) error {
	db := utils.GetLocal[*sql.DB](c, "db")
	users, err := repositories.GetUsers(db)

	if err != nil {
		return c.JSON(utils.E404("Users not found", err))
	}

	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := utils.GetLocal[*sql.DB](c, "db")
	res, err := database.DoQuery(db, "SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return c.JSON(utils.E503("Error while getting user", err))
	}

	var user types.User
	for res.Next() {
		err := res.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Role)
		if err != nil {
			return c.JSON(utils.E503("Error while getting user", err))
		}
	}

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")

	user := new(types.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := utils.ValidateStruct(*user)
	if errors != "" {
		return c.JSON(utils.E400("Bad request :\n"+errors, nil))
	}

	db := utils.GetLocal[*sql.DB](c, "db")
	res, err := database.DoQuery(db, "SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return c.JSON(utils.E503("Error while getting user", err))
	}

	var oldUser types.User
	for res.Next() {
		err := res.Scan(&oldUser.ID, &oldUser.FirstName, &oldUser.LastName, &oldUser.Email, &oldUser.Password, &oldUser.Role)
		if err != nil {
			return c.JSON(utils.E503("Error while getting user", err))
		}
	}

	if !utils.CheckPasswordHash(user.Password, oldUser.Password) {
		user.Password, _ = utils.HashPassword(user.Password)
	}

	_, errRepo := repositories.UpdateUser(db, user, id)

	if err != nil {
		return c.JSON(utils.E400("Bad request :\n"+err.Error(), errRepo))
	}

	successMessage := fmt.Sprintf("User %s updated successfully", id)
	return c.JSON(types.HttpResponse{Status: 1, Message: successMessage, HttpCode: 200})
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	db := utils.GetLocal[*sql.DB](c, "db")
	_, err := repositories.DeleteUser(db, id)
	if err != nil {
		return c.JSON(utils.E400("Bad request :\n"+err.Error(), err))
	}

	successMessage := fmt.Sprintf("User %s deleted successfully", id)
	return c.JSON(types.HttpResponse{Status: 1, Message: successMessage, HttpCode: 200})
}
