package response

import (
	"github.com/gofiber/fiber/v2"
	"server/pkg/entities"
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

type PostSuccessResponseType struct {
	State bool        `json:"state"`
	Data  Post        `json:"data"`
	Error interface{} `json:"error"`
}

func PostsSuccessResponse(data *[]Post) *fiber.Map {
	return &fiber.Map{
		"state": true,
		"data":  data,
		"error": nil,
	}
}

type PostsSuccessResponseType struct {
	State bool        `json:"state"`
	Data  []Post      `json:"data"`
	Error interface{} `json:"error"`
}

func PostErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"state": true,
		"data":  "",
		"error": err.Error(),
	}
}

type PostErrorResponseType struct {
	State bool        `json:"state"`
	Data  []Post      `json:"data"`
	Error interface{} `json:"error"`
}
