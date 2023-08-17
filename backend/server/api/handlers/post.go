package handler

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"server/api/presenter/request"
	"server/api/presenter/response"
	"server/pkg/post"
)

// @Summary Add a new post
// @Description Add a new post to the database
// @Tags posts
// @Accept json
// @Produce json
// @Param post body request.PostRequest true "Add post"
// @Success 200 {object} response.PostSuccessResponseType "Successfully added the post"
// @Failure 400 {object} response.PostErrorResponseType "Invalid request body"
// @Failure 500 {object} response.PostErrorResponseType "Internal server error"
// @Router /posts [post]
func AddPost(service post.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody request.PostRequest
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

// @Summary Update an existing post
// @Description Update post details in the database
// @Tags posts
// @Accept json
// @Produce json
// @Param post body request.UpdatePostRequest true "Update post"
// @Success 200 {object} response.PostSuccessResponseType "Successfully updated the post"
// @Failure 400 {object} response.PostErrorResponseType "Invalid request body"
// @Failure 500 {object} response.PostErrorResponseType "Internal server error"
// @Router /posts [put]
func UpdatePost(service post.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody request.UpdatePostRequest
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

// @Summary Remove a post
// @Description Remove a post from the database based on its ID
// @Tags posts
// @Accept json
// @Produce json
// @Param post body request.DeletePostRequest true "Update post"
// @Success 200 {object} response.PostSuccessResponseType "Successfully Removed the post"
// @Failure 400 {object} response.PostErrorResponseType "Invalid request body"
// @Failure 500 {object} response.PostErrorResponseType "Internal server error"
// @Router /posts [delete]
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

// @Summary Get all posts
// @Description Get all posts from the database
// @Tags posts
// @Produce json
// @Success 200 {object} response.PostsSuccessResponseType
// @Failure 500 {object} response.PostErrorResponseType
// @Router /posts [get]
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
