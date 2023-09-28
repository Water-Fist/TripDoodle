package post

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"server/api/presenter/request"
	"server/api/presenter/response"
	"server/pkg/entities"
	"testing"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreatePost(post *entities.Post) (*entities.Post, error) {
	args := m.Called(post)
	return args.Get(0).(*entities.Post), args.Error(1)
}

func (m *MockRepository) ReadPost() (*[]response.Post, error) {
	args := m.Called()
	return args.Get(0).(*[]response.Post), args.Error(1)
}

func (m *MockRepository) UpdatePost(post *entities.Post) (*entities.Post, error) {
	args := m.Called(post)
	return args.Get(0).(*entities.Post), args.Error(1)
}

func (m *MockRepository) DeletePost(ID string) error {
	args := m.Called(ID)
	return args.Error(0)
}

func TestServiceMethods(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	postRequest := &request.PostRequest{Title: "Test Title", Content: "Test Content"}
	postResponse := &entities.Post{ID: 1, Title: "Test Title", Content: "Test Content"}
	posts := &[]response.Post{{ID: 1, Title: "Test Title", Content: "Test Content"}}
	updatePostRequest := &request.UpdatePostRequest{ID: 1, Title: "Updated Title", Content: "Updated Content"}

	mockRepo.On("CreatePost", mock.Anything).Return(postResponse, nil)
	mockRepo.On("ReadPost").Return(posts, nil)
	mockRepo.On("UpdatePost", mock.Anything).Return(postResponse, nil)
	mockRepo.On("DeletePost", "1").Return(nil)

	result1, err1 := service.InsertPost(postRequest)
	assert.Nil(t, err1)
	assert.Equal(t, result1, postResponse)

	result2, err2 := service.FetchPosts()
	assert.Nil(t, err2)
	assert.Equal(t, result2, posts)

	result3, err3 := service.UpdatePost(updatePostRequest)
	assert.Nil(t, err3)
	assert.Equal(t, result3, postResponse)

	err4 := service.RemovePost("1")
	assert.Nil(t, err4)

	mockRepo.AssertExpectations(t)
}
