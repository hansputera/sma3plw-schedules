package main

import (
	"smanti_schedules/schedules"
	"smanti_schedules/services"
	"smanti_schedules/teachers"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New(fiber.Config{
		ServerHeader: "SMANTI-Schedules",
		AppName:      "Smanti Schedules",
		JSONEncoder:  sonic.Marshal,
		JSONDecoder:  sonic.Unmarshal,
	})

	app.Use(cors.New())
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

	app.Get("/times/:day", func(c *fiber.Ctx) error {
		times := services.GetScheduleTimes(c.Params("day"))

		if times == nil {
			c.JSON(map[string]string{
				"error": "fail to generate the times",
			})
			return nil
		}

		c.JSON(map[string][]string{
			"data": *times,
		})
		return nil
	})

	app.Get("/schedules/:class", func(c *fiber.Ctx) error {
		data := &teachers.TeacherMaps{}
		if len(c.Query("teachers")) > 0 {
			results, err := teachers.GetTeachers(TEACHERS_PATH)
			if err != nil {
				c.JSON(map[string]string{
					"error": err.Error(),
				})
				return nil
			}
			
			data = &results
		}
		scheds, err := schedules.GetSchedules(SCHEDULES_PATH, c.Params("class"), data)
		if err != nil {
			return err
		}

		c.JSON(map[string]schedules.ScheduleMaps{
			"data": *scheds,
		})
		return nil
	})

	app.Listen(":9000")
}
