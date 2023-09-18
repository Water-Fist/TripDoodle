package response

import (
	"github.com/gofiber/fiber/v2"
	"server/pkg/entities"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

func UserSuccessResponse(data *entities.User) *fiber.Map {
	user := User{
		ID:       data.ID,
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Nickname: data.Nickname,
	}
	return &fiber.Map{
		"state": true,
		"data":  user,
		"error": nil,
	}
}

type UserSuccessResponseType struct {
	State bool        `json:"state"`
	Data  User        `json:"data"`
	Error interface{} `json:"error"`
}

func UsersSuccessResponse(data *[]entities.User) *fiber.Map {
	return &fiber.Map{
		"state": true,
		"data":  data,
		"error": nil,
	}
}

type UsersErrorResponseType struct {
	State bool        `json:"state"`
	Data  []User      `json:"data"`
	Error interface{} `json:"error"`
}

func UserErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"state": false,
		"data":  nil,
		"error": err.Error(),
	}
}

type UserErrorResponseType struct {
	State bool        `json:"state"`
	Data  User        `json:"data"`
	Error interface{} `json:"error"`
}

func CheckResponse(data bool) *fiber.Map {
	return &fiber.Map{
		"state": true,
		"data":  data,
		"error": nil,
	}
}

func CheckErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"state": true,
		"data":  nil,
		"error": err.Error(),
	}
}
