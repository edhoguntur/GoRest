package main

import (
	//"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	//Initial Commit
	//fmt.Println("Hello World!")

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":3030")

}
