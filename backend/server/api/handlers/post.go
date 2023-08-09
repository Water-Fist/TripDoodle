package handler

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"server/api/presenter/request"
	"server/api/presenter/response"
	"server/pkg/entities"
	"server/pkg/post"
)

func AddPost(service post.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Post
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(response.PostErrorResponse(err))
		}
		if requestBody.Content == "" || requestBody.Title == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(response.PostErrorResponse(errors.New(
				"Please specify title and content")))
		}
		result, err := service.InsertPost(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(response.PostErrorResponse(err))
		}
		return c.JSON(response.PostSuccessResponse(result))
	}
}

func UpdatePost(service post.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Post
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(response.PostErrorResponse(err))
		}
		result, err := service.UpdatePost(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(response.PostErrorResponse(err))
		}
		return c.JSON(response.PostSuccessResponse(result))
	}
}

func RemovePost(service post.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody request.DeletePostRequest
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(response.PostErrorResponse(err))
		}
		postID := requestBody.ID
		err = service.RemovePost(postID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(response.PostErrorResponse(err))
		}
		return c.JSON(&fiber.Map{
			"status": true,
			"data":   "updated successfully",
			"err":    nil,
		})
	}
}

func GetPosts(service post.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.FetchPosts()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(response.PostErrorResponse(err))
		}
		return c.JSON(response.PostsSuccessResponse(fetched))
	}
}
