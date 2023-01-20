package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New(fiber.Config{
		ServerHeader: "SMANTI-Schedules",
		AppName: "Smanti Schedules",
	})

	app.Get("/", func(c *fiber.Ctx) error {
		c.SendString("Hello World!")
		return nil
	})

	app.Listen(":9000")
}