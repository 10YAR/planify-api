package controllers_test

import (
	"api/controllers"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestGetShops(t *testing.T) {
	// Initialize Fiber app
	app := fiber.New()

	app.Get("/shops", controllers.GetShops)

	// Create a test request End the GetShops endpoint
	req := httptest.NewRequest("GET", "/shops", nil)
	fmt.Println("req", req)
	res, _ := app.Test(req, -1)
	assert.Equal(t, 200, res.StatusCode)
}

func TestCreateShop(t *testing.T) {
	app := fiber.New()

	app.Post("/shops", controllers.CreateShop)

	req := httptest.NewRequest("POST", "/shops", nil)
	res, _ := app.Test(req, -1)
	assert.Equal(t, 200, res.StatusCode)
}
