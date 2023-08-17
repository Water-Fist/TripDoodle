package handler

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"server/api/presenter/request"
	"server/api/presenter/response"
	"server/pkg/entities"
	"server/pkg/sight"
	"strconv"
)

// @Summary Add a new sight
// @Description Add a new sight to the database
// @Tags sights
// @Accept json
// @Produce json
// @Param sight body request.SightRequest true "Add Sight"
// @Success 200 {object} response.SightSuccessResponseType "Successfully added the sight"
// @Failure 400 {object} response.SightErrorResponseType "Invalid request body"
// @Failure 500 {object} response.SightErrorResponseType "Internal server error"
// @Router /sights [post]
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

// @Summary Update an existing sight
// @Description Update sight details in the database
// @Tags sights
// @Accept json
// @Produce json
// @Param sight body request.UpdateSightRequest true "Update sight"
// @Success 200 {object} response.SightSuccessResponseType "Successfully updated the sight"
// @Failure 400 {object} response.SightErrorResponseType "Invalid request body"
// @Failure 500 {object} response.SightErrorResponseType "Internal server error"
// @Router /sights [put]
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

// @Summary Remove a sight
// @Description Remove a sight from the database based on its ID
// @Tags sights
// @Accept json
// @Produce json
// @Param sight body request.DeleteSightRequest true "Update sight"
// @Success 200 {object} response.SightSuccessResponseType "Successfully Removed the sight"
// @Failure 400 {object} response.SightErrorResponseType "Invalid request body"
// @Failure 500 {object} response.SightErrorResponseType "Internal server error"
// @Router /sights [delete]
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

// @Summary Get all sights
// @Description Get all sights from the database
// @Tags sights
// @Produce json
// @Success 200 {object} response.SightsSuccessResponseType "Successfully Got the sight"
// @Failure 500 {object} response.SightErrorResponseType "Internal server error"
// @Router /sights [get]
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

// @Summary Get Current sights
// @Description Gets information about the tourist attractions currently located in the database Fetch all sights from the database
// @Tags sights
// @Produce json
// @Param Latitude query float32 true "Latitude for the sight"
// @Param Longitude query float32 true "Longitude for the sight"
// @Success 200 {object} response.SightsSuccessResponseType "Successfully Got the sight"
// @Failure 400 {object} response.SightErrorResponseType "Invalid request body"
// @Failure 500 {object} response.SightErrorResponseType "Internal server error"
// @Router /sights/location [get]
func LoadSight(service sight.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		latitude := c.Query("Latitude")
		longitude := c.Query("Longitude")

		if latitude == "" || longitude == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(response.SightErrorResponse(errors.New("Latitude and Longitude parameters are required")))
		}

		lat, err := strconv.ParseFloat(latitude, 32)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(response.SightErrorResponse(errors.New("Invalid Latitude value")))
		}

		lon, err := strconv.ParseFloat(longitude, 32)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(response.SightErrorResponse(errors.New("Invalid Longitude value")))
		}

		result, err := service.LoadSight(float32(lat), float32(lon))
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(response.SightErrorResponse(err))
		}
		return c.JSON(response.SightsLoadSuccessResponse(result))
	}
}
