package sight

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"server/api/presenter/response"
	"server/pkg/entities"
	"testing"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreateSight(sight *entities.Sight) (*entities.Sight, error) {
	args := m.Called(sight)
	return args.Get(0).(*entities.Sight), args.Error(1)
}

func (m *MockRepository) ReadSight() (*[]response.Sight, error) {
	args := m.Called()
	return args.Get(0).(*[]response.Sight), args.Error(1)
}

func (m *MockRepository) UpdateSight(sight *entities.Sight) (*entities.Sight, error) {
	args := m.Called(sight)
	return args.Get(0).(*entities.Sight), args.Error(1)
}

func (m *MockRepository) DeleteSight(ID string) error {
	args := m.Called(ID)
	return args.Error(0)
}

func (m *MockRepository) LoadSight(Latitude float32, Longitude float32) (*[]response.SightLoad, error) {
	args := m.Called(Latitude, Longitude)
	return args.Get(0).(*[]response.SightLoad), args.Error(1)
}

func TestServiceMethods(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	sight := &entities.Sight{ID: 1, Name: "Test Sight"}
	sights := &[]response.Sight{{ID: 1, Name: "Test Sight"}}
	sightLoad := &[]response.SightLoad{{ID: 1, Name: "Test Sight"}}
	latitude := float32(37.123456)
	longitude := float32(127.123456)

	mockRepo.On("CreateSight", sight).Return(sight, nil)
	mockRepo.On("ReadSight").Return(sights, nil)
	mockRepo.On("UpdateSight", sight).Return(sight, nil)
	mockRepo.On("DeleteSight", "1").Return(nil)
	mockRepo.On("LoadSight", latitude, longitude).Return(sightLoad, nil)

	result1, err1 := service.InsertSight(sight)
	assert.Nil(t, err1)
	assert.Equal(t, result1, sight)

	result2, err2 := service.FetchSights()
	assert.Nil(t, err2)
	assert.Equal(t, result2, sights)

	result3, err3 := service.UpdateSight(sight)
	assert.Nil(t, err3)
	assert.Equal(t, result3, sight)

	err4 := service.RemoveSight("1")
	assert.Nil(t, err4)

	result5, err5 := service.LoadSight(latitude, longitude)
	assert.Nil(t, err5)
	assert.Equal(t, result5, sightLoad)

	mockRepo.AssertExpectations(t)
}
