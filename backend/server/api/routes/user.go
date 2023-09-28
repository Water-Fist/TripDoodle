package routes

import (
	"github.com/gofiber/fiber/v2"
	handler "server/api/handlers/user"
	"server/pkg/user"
)

func UserRouter(app fiber.Router, service user.Service) {
	app.Post("/users", handler.AddUser(service))
	app.Put("/users", handler.UpdateUser(service))
	app.Get("/users", handler.GetUsers(service))
	app.Get("/users/:id", handler.GetUser(service))
	app.Delete("/users", handler.RemoveUser(service))
	app.Post("/users/login", handler.Login(service))
	app.Get("/users/check/email/:email", handler.CheckEmail(service))
	app.Get("/users/check/nickname/:nickname", handler.CheckNickname(service))
}
