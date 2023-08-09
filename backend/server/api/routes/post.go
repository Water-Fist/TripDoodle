package routes

import (
	handler "backend/api/handlers"
	"backend/pkg/post"
	"github.com/gofiber/fiber/v2"
)

func PostRouter(app fiber.Router, service post.Service) {
	app.Get("/posts", handler.GetPosts(service))
	app.Post("/posts", handler.AddPost(service))
	app.Put("/posts", handler.UpdatePost(service))
	app.Delete("/posts", handler.RemovePost(service))
}
