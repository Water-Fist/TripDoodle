package handler

import (
	"backend/api/presenter/request"
	"backend/api/presenter/response"
	"backend/pkg/entities"
	"backend/pkg/sight"
	"errors"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func AddSight(service sight.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Sight
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(response.SightErrorResponse(err))
		}
		if requestBody.Name == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(response.SightErrorResponse(errors.New(
				"Please specify title and content")))
		}
		result, err := service.InsertSight(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(response.SightErrorResponse(err))
		}
		return c.JSON(response.SightSuccessResponse(result))
	}
}

func UpdateSight(service sight.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Sight
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(response.SightErrorResponse(err))
		}
		result, err := service.UpdateSight(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(response.SightErrorResponse(err))
		}
		return c.JSON(response.SightSuccessResponse(result))
	}
}

func RemoveSight(service sight.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody request.DeleteSightRequest
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(response.SightErrorResponse(err))
		}
		sightID := requestBody.ID
		err = service.RemoveSight(sightID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(response.SightErrorResponse(err))
		}
		return c.JSON(&fiber.Map{
			"status": true,
			"data":   "updated successfully",
			"err":    nil,
		})
	}
}

func GetSights(service sight.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.FetchSights()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(response.SightErrorResponse(err))
		}
		return c.JSON(response.SightsSuccessResponse(fetched))
	}
}

func LoadSight(service sight.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody request.LoadSightRequest
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(response.SightErrorResponse(err))
		}
		result, err := service.LoadSight(requestBody.Latitude, requestBody.Longitude)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(response.SightErrorResponse(err))
		}
		return c.JSON(response.SightsSuccessResponse(result))
	}
}
