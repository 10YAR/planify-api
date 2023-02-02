package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := DefineRoutes()

	fmt.Printf("starting Planify API at port 8000 \n")
	log.Fatal(app.Listen(":8000"))
}

func GetFiberApp() *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: false,
		ServerHeader:  "Fiber",
		AppName:       "Planify API 1.0",
	})

	return app
}
