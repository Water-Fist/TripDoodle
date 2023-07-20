package presenter

import (
	"backend/pkg/entities"
	"github.com/gofiber/fiber/v2"
)

type Post struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	ImageUrl string `json:"imageUrl"`
	State    bool   `json:"state"`
	SightID  int    `json:"sightId"`
}

func PostSuccessResponse(data *entities.Post) *fiber.Map {
	post := Post{
		ID:       data.ID,
		Title:    data.Title,
		Content:  data.Content,
		ImageUrl: data.ImageUrl,
		State:    data.State,
		SightID:  data.SightId,
	}
	return &fiber.Map{
		"state": true,
		"data":  post,
		"error": nil,
	}
}

func PostsSuccessResponse(data *[]Post) *fiber.Map {
	return &fiber.Map{
		"state": true,
		"data":  data,
		"error": nil,
	}
}

func PostErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"state": true,
		"data":  "",
		"error": err.Error(),
	}
}
