package presenter

import (
	"backend/pkg/entities"
	"github.com/gofiber/fiber/v2"
)

type Sight struct {
	ID        int    `json:"id"`
	Name      string `json:"title"`
	Latitude  string `json:"content"`
	Longitude string `json:"imageUrl"`
	Area      bool   `json:"state"`
}

func SightSuccessResponse(data *entities.Sight) *fiber.Map {
	sight := Sight{
		ID:        data.ID,
		Name:      data.Name,
		Latitude:  data.Latitude,
		Longitude: data.Longitude,
		Area:      data.Area,
	}
	return &fiber.Map{
		"state": true,
		"data":  sight,
		"error": nil,
	}
}

func SightsSuccessResponse(data *[]Sight) *fiber.Map {
	return &fiber.Map{
		"state": true,
		"data":  data,
		"error": nil,
	}
}

func SightErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"state": true,
		"data":  "",
		"error": err.Error(),
	}
}