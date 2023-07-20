package handler

import (
	"backend/api/presenter"
	"backend/pkg/entities"
	"backend/pkg/post"
	"errors"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func AddPost(service post.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Post
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.PostErrorResponse(err))
		}
		if requestBody.Content == "" || requestBody.Title == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.PostErrorResponse(errors.New(
				"Please specify title and content")))
		}
		result, err := service.InsertPost(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.PostErrorResponse(err))
		}
		return c.JSON(presenter.PostSuccessResponse(result))
	}
}

func UpdatePost(service post.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Post
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.PostErrorResponse(err))
		}
		result, err := service.UpdatePost(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.PostErrorResponse(err))
		}
		return c.JSON(presenter.PostSuccessResponse(result))
	}
}

func RemovePost(service post.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.DeleteRequest
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.PostErrorResponse(err))
		}
		postID := requestBody.ID
		err = service.RemovePost(postID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.PostErrorResponse(err))
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
			return c.JSON(presenter.PostErrorResponse(err))
		}
		return c.JSON(presenter.PostsSuccessResponse(fetched))
	}
}
