package middlewares

import (
	"api/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func ProcessAuth(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["id"] == nil {
		return c.JSON(utils.E401("Unauthorized", nil))
	}

	// Poursuivre la requête
	return c.Next()
}
