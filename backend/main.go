package main

import (
	"backend/database"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	// Connect with database
	database.ConnectDb()
	app := fiber.New()

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pingpong by fiber\n")
	})

	log.Fatal(app.Listen(":3000"))
}
