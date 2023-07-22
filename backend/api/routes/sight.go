package routes

import (
	handler "backend/api/handlers"
	"backend/pkg/sight"
	"github.com/gofiber/fiber/v2"
)

func SightRouter(app fiber.Router, service sight.Service) {
	app.Get("/sights", handler.GetSights(service))
	app.Get("/sights/location", handler.LoadSight(service))
	app.Post("/sights", handler.AddSight(service))
	app.Put("/sights", handler.UpdateSight(service))
	app.Delete("/sights", handler.RemoveSight(service))
}
