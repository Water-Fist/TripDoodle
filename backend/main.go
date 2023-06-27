package main

import (
	"backend/database"
	"backend/router"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"log"
)

func setUpPosts(app *fiber.App) {
	var path = "/post"

	//app.Get(path, router.AllPosts)
	//app.Post(path, router.AddPost)
	app.Get(path+"/:id", router.GetPost)
	//app.Put("/update", router.PostUpdate)
	//app.Delete("/delete", router.PostDelete)
}

func main() {
	// Connect with database
	database.ConnectDb()
	app := fiber.New()

	setUpPosts(app)

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pingpong by fiber\n")
	})

	log.Fatal(app.Listen(":3000"))
}
