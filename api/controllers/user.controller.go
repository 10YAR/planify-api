package controllers

import (
	"api/database"
	"api/types"
	"api/utils"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	res, err := database.DoQuery("SELECT * FROM users")
	if err != nil {
		return c.JSON(utils.E503("Error while getting users"))
	}

	var users []types.User
	for res.Next() {
		var user types.User
		err := res.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Role)
		if err != nil {
			return c.JSON(utils.E503("Error while getting users"))
		}

		users = append(users, user)
	}

	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	res, err := database.DoQuery("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return c.JSON(utils.E503("Error while getting user"))
	}

	var user types.User
	for res.Next() {
		err := res.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Role)
		if err != nil {
			return c.JSON(utils.E503("Error while getting user"))
		}
	}

	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	user := new(types.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := utils.ValidateStruct(*user)
	if errors != "" {
		return c.JSON(utils.E400("Bad request :\n" + errors))
	}

	_, err := database.DoQuery("INSERT INTO users (firstName, lastName, email, password, role) VALUES (?, ?, ?, ?, ?)", user.FirstName, user.LastName, user.Email, user.Password, user.Role)
	if err != nil {
		return c.JSON(utils.E503("Error while creating user"))
	}

	return c.JSON(types.HttpResponse{Status: 1, Message: "User created successfully", HttpCode: 200})
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
		return c.JSON(utils.E400("Bad request :\n" + errors))
	}

	_, err := database.DoQuery("UPDATE users SET firstName = ?, lastName = ?, email = ?, password = ?, role = ? WHERE id = ?", user.FirstName, user.LastName, user.Email, user.Password, user.Role, id)
	if err != nil {
		return c.JSON(utils.E503("Error while updating user"))
	}

	return c.JSON(types.HttpResponse{Status: 1, Message: "User updated successfully", HttpCode: 200})
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	_, err := database.DoQuery("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return c.JSON(utils.E503("Error while deleting user"))
	}

	return c.JSON(types.HttpResponse{Status: 1, Message: "User deleted successfully", HttpCode: 200})
}
