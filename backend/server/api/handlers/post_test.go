package handler

import (
	"bytes"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http/httptest"
	"server/api/presenter/request"
	"server/api/presenter/response"
	"server/pkg/entities"
	"testing"
)

type MockPostService struct {
	mock.Mock
}

func (m *MockPostService) InsertPost(post *request.PostRequest) (*entities.Post, error) {
	args := m.Called(post)
	return args.Get(0).(*entities.Post), args.Error(1)
}

func (m *MockPostService) FetchPosts() (*[]response.Post, error) {
	args := m.Called()
	return args.Get(0).(*[]response.Post), args.Error(1)
}

func (m *MockPostService) UpdatePost(post *request.UpdatePostRequest) (*entities.Post, error) {
	args := m.Called(post)
	return args.Get(0).(*entities.Post), args.Error(1)
}

func (m *MockPostService) RemovePost(ID string) error {
	args := m.Called(ID)
	return args.Error(0)
}

func TestAddPost(t *testing.T) {
	app := fiber.New()

	mockService := new(MockPostService)
	app.Post("/posts", AddPost(mockService))

	t.Run("it should return 400 for invalid request body", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/posts", bytes.NewBuffer([]byte("invalid_json")))
		resp, _ := app.Test(req)

		assert.Equal(t, 400, resp.StatusCode)
	})

	t.Run("it_should_return_400_if_title_or_content_is_missing", func(t *testing.T) {
		body := `{"title": "", "content": "some content"}`
		req := httptest.NewRequest("POST", "/posts", bytes.NewBuffer([]byte(body)))
		resp, _ := app.Test(req)

		assert.Equal(t, 400, resp.StatusCode)
	})

	t.Run("it should return 500 for internal server error", func(t *testing.T) {
		body := `{
        "title": "ooo",
        "content": "ㄴㄴ",
        "imageUrl": "ㄴㄴㄴ",
        "state": true,
        "sightID": 1
    }`
		req := httptest.NewRequest("POST", "/posts", bytes.NewBuffer([]byte(body)))
		mockService.On("InsertPost", mock.Anything).Return(nil, errors.New("Some error")).Once()
		resp, _ := app.Test(req)

		assert.Equal(t, 500, resp.StatusCode)
	})

	t.Run("it should return 200 for successful post insertion", func(t *testing.T) {
		body := `{
        "title": "ooo",
        "content": "ㄴㄴ",
        "imageUrl": "ㄴㄴㄴ",
        "state": true,
        "sightID": 1
    }`
		req := httptest.NewRequest("POST", "/posts", bytes.NewBuffer([]byte(body)))
		post := entities.Post{
			ID:       1,
			Title:    "test title",
			Content:  "test content",
			ImageUrl: "test url",
			State:    true,
		}
		mockService.On("InsertPost", mock.Anything).Return(&post, nil).Once()
		resp, _ := app.Test(req)

		assert.Equal(t, 200, resp.StatusCode)
	})
}
