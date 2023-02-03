package main

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// Load env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Define App Routes
	App := DefineRoutes()

	// Start App
	fmt.Printf("starting Planify API at port 8000 \n")
	log.Fatal(App.Listen(":8000"))
}

func GetFiberApp() *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: false,
		ServerHeader:  "Fiber",
		AppName:       "Planify API 1.0",
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
	})
	return app
}
