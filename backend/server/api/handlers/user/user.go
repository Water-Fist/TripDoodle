package user

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"server/api/presenter/request"
	"server/api/presenter/response"
	"server/pkg/entities"
	"server/pkg/user"
	"strconv"
)

// AddUser @Summary Add a new user
// @Description Add a new user to the database
// @Tags users
// @Accept json
// @Produce json
// @Param post body request.UserRequest true "Add user"
// @Success 200 {object} response.UserSuccessResponseType "Successfully added the user"
// @Failure 400 {object} response.UserErrorResponseType "Invalid request body"
// @Failure 500 {object} response.UserErrorResponseType "Internal server error"
// @Router /users [post]
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

// UpdateUser @Summary Update an existing user
// @Description Update user details in the database
// @Tags users
// @Accept json
// @Produce json
// @Param post body request.UpdateUserRequest true "Update user"
// @Success 200 {object} response.UserSuccessResponseType "Successfully updated the user"
// @Failure 400 {object} response.UserErrorResponseType "Invalid request body"
// @Failure 500 {object} response.UserErrorResponseType "Internal server error"
// @Router /users [put]
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

// RemoveUser @Summary Remove an existing user
// @Description Remove user from the database
// @Tags users
// @Accept json
// @Produce json
// @Param post body request.DeleteUserRequest true "Delete user"
// @Success 200 {object} response.UserSuccessResponseType "Successfully deleted the user"
// @Failure 400 {object} response.UserErrorResponseType "Invalid request body"
// @Failure 500 {object} response.UserErrorResponseType "Internal server error"
// @Router /users [delete]
func RemoveUser(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody request.DeleteUserRequest
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(response.UserErrorResponse(err))
		}
		userID := requestBody.ID
		err = service.RemoveUser(strconv.Itoa(int(userID)))
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(response.UserErrorResponse(err))
		}
		return c.JSON(&fiber.Map{
			"status": true,
			"data":   "deleted successfully",
			"err":    nil,
		})
	}
}

// GetUser @Summary Get a user
// @Description Get a user from the database based on its ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} response.UserSuccessResponseType
// @Failure 400 {object} response.UserErrorResponseType
// @Failure 500 {object} response.UserErrorResponseType
// @Router /users/{id} [get]
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

// GetUsers @Summary Get all users
// @Description Get all users from the database
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} response.UsersSuccessResponseType
// @Failure 500 {object} response.UserErrorResponseType
// @Router /users [get]
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

// CheckEmail @Summary Check if email exists
// @Description Check if email exists in the database
// @Tags users
// @Accept json
// @Produce json
// @Param email path string true "User email"
// @Success 200 {object} response.CheckResponseType
// @Failure 500 {object} response.CheckErrorResponseType
// @Router /users/check/email/{email} [get]
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

// CheckNickname @Summary Check if nickname exists
// @Description Check if nickname exists in the database
// @Tags users
// @Accept json
// @Produce json
// @Param nickname path string true "User nickname"
// @Success 200 {object} response.CheckResponseType
// @Failure 500 {object} response.CheckErrorResponseType
// @Router /users/check/nickname/{nickname} [get]
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

// Login @Summary Login
// @Description Login with email and password
// @Tags users
// @Accept json
// @Produce json
// @Param user body request.LoginRequest true "Login"
// @Success 200 {object} response.CheckResponseType
// @Failure 500 {object} response.CheckErrorResponseType
// @Router /users/login [post]
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
