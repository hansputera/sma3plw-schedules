package main

import (
	"smanti_schedules/teachers"

	"github.com/gofiber/fiber/v2"
	"github.com/bytedance/sonic"
)

func main() {
	app := fiber.New(fiber.Config{
		ServerHeader: "SMANTI-Schedules",
		AppName: "Smanti Schedules",
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		c.SendString("Hello World!")
		return nil
	})

	app.Get("/teachers", func(c *fiber.Ctx) error {
		data, err := teachers.GetTeachers(TEACHERS_PATH)
		if err != nil {
			c.JSON(map[string]string{
				"error": err.Error(),
			})
			return nil
		}

		c.JSON(map[string]teachers.TeacherMaps{
			"data": data,
		})
		return nil
	})

	app.Listen(":9000")
}