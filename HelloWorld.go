package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type person struct {
	name string
	age  int
}

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		newPerson := &person{
			name: "Judd Misael R. Baguio",
			age:  2,
		}

		return c.Status(200).JSON(&fiber.Map{
			"name": newPerson.name,
			"age":  newPerson.age,
		})
	})

	log.Fatal(app.Listen(":8080"))
}
