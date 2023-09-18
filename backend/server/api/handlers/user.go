package handler

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"server/api/presenter/request"
	"server/api/presenter/response"
	"server/pkg/entities"
	"server/pkg/user"
)

func AddUser(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.User
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(response.UserErrorResponse(err))
		}
		if requestBody.Name == "" || requestBody.Email == "" || requestBody.Password == "" || requestBody.Nickname == "" {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(response.UserErrorResponse(errors.New(
				"Please specify name, email, password and nickname")))
		}
		result, err := service.InsertUser(&requestBody)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(response.UserErrorResponse(err))
		}
		return c.JSON(response.UserSuccessResponse(result))
	}
}

func UpdateUser(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.User
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(response.UserErrorResponse(err))
		}
		result, err := service.UpdateUser(&requestBody)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(response.UserErrorResponse(err))
		}
		return c.JSON(response.UserSuccessResponse(result))
	}
}

func RemoveUser(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody request.DeleteUserRequest
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(response.UserErrorResponse(err))
		}
		userID := requestBody.ID
		err = service.RemoveUser(userID)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(response.UserErrorResponse(err))
		}
		return c.JSON(response.UserSuccessResponse(nil))
	}
}

func GetUser(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		result, err := service.GetUserById(id)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(response.UserErrorResponse(err))
		}
		return c.JSON(response.UserSuccessResponse(result))
	}
}

func GetUsers(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var result, err = service.GetUsers()
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(response.UserErrorResponse(err))
		}
		return c.JSON(response.UsersSuccessResponse(result))
	}
}

func CheckEmail(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		email := c.Params("email")
		result, err := service.CheckEmail(email)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(response.CheckErrorResponse(err))
		}
		return c.JSON(response.CheckResponse(result))
	}
}

func CheckNickname(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		nickname := c.Params("nickname")
		result, err := service.CheckNickname(nickname)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(response.CheckErrorResponse(err))
		}
		return c.JSON(response.CheckResponse(result))
	}
}

func Login(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody request.LoginRequest
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(response.UserErrorResponse(err))
		}
		result, err := service.Login(requestBody.Email, requestBody.Password)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(response.CheckErrorResponse(err))
		}
		return c.JSON(response.CheckResponse(result))
	}
}
