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

func (m *MockPostService) InsertPost(req *request.PostRequest) (*entities.Post, error) {
	args := m.Called(req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Post), args.Error(1)
}

func (m *MockPostService) FetchPosts() (*[]response.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockPostService) UpdatePost(post *request.UpdatePostRequest) (*entities.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockPostService) RemovePost(ID string) error {
	//TODO implement me
	panic("implement me")
}

func TestAddPost(t *testing.T) {
	app := fiber.New()

	mockService := new(MockPostService)
	app.Post("/posts", AddPost(mockService))

	t.Run("유효하지 않은 request body 인 경우, 400 에러 반환", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/posts", bytes.NewBuffer([]byte("invalid_json")))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, 400, resp.StatusCode)
	})

	t.Run("제목, 내용, 관광지 PK이 누락된 경우, 400 에러 반환", func(t *testing.T) {
		body := `{
			  "title": "",
			  "content": "test",
			  "imageUrl": "test",
			  "sightId": 1
		  }`
		req := httptest.NewRequest("POST", "/posts", bytes.NewBuffer([]byte(body)))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, 400, resp.StatusCode)
	})

	t.Run("내부 서버 오류의 경우, 500 에러 반환", func(t *testing.T) {
		body := `{
			  "title": "test",
			  "content": "test",
			  "imageUrl": "test",
			  "sightId": 1
		  }`
		req := httptest.NewRequest("POST", "/posts", bytes.NewBuffer([]byte(body)))
		req.Header.Set("Content-Type", "application/json")
		mockService.On("InsertPost", mock.Anything).Return(nil, errors.New("500 error")).Once()
		resp, _ := app.Test(req)

		assert.Equal(t, 500, resp.StatusCode)
	})

	t.Run("성공 시, 200 반환", func(t *testing.T) {
		body := `{
			  "title": "test",
			  "content": "test",
			  "imageUrl": "test",
			  "sightId": 1
		  }`
		req := httptest.NewRequest("POST", "/posts", bytes.NewBuffer([]byte(body)))
		req.Header.Set("Content-Type", "application/json")
		post := entities.Post{
			ID:       1,
			Title:    "test",
			Content:  "test",
			ImageUrl: "test",
			State:    false,
			SightId:  1,
		}
		mockService.On("InsertPost", mock.Anything).Return(&post, nil).Once()
		resp, _ := app.Test(req)

		assert.Equal(t, 200, resp.StatusCode)
	})

	mockService.AssertExpectations(t)
}
