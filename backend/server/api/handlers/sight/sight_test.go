package sight

import (
	"bytes"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http/httptest"
	"server/api/presenter/response"
	"server/pkg/entities"
	"testing"
)

type MockSightService struct {
	mock.Mock
}

func (m *MockSightService) InsertSight(sight *entities.Sight) (*entities.Sight, error) {
	args := m.Called(sight)
	return args.Get(0).(*entities.Sight), args.Error(1)
}

func (m *MockSightService) FetchSights() (*[]response.Sight, error) {
	args := m.Called()
	return args.Get(0).(*[]response.Sight), args.Error(1)
}

func (m *MockSightService) UpdateSight(sight *entities.Sight) (*entities.Sight, error) {
	args := m.Called(sight)
	return args.Get(0).(*entities.Sight), args.Error(1)
}

func (m *MockSightService) RemoveSight(ID string) error {
	args := m.Called(ID)
	return args.Error(0)
}

func (m *MockSightService) LoadSight(Latitude float32, Longitude float32) (*[]response.SightLoad, error) {
	args := m.Called(Latitude, Longitude)
	return args.Get(0).(*[]response.SightLoad), args.Error(1)
}

func TestAddSight(t *testing.T) {
	app := fiber.New()

	service := new(MockSightService)
	app.Post("/sights", AddSight(service))

	t.Run("유효하지 않은 request body 인 경우, 400 에러 반환", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/sights", bytes.NewBuffer([]byte("invalid_json")))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, 400, resp.StatusCode)
	})

	t.Run("내부 서버 오류의 경우, 500 에러 반환", func(t *testing.T) {
		body := `{
			"name": "Test Sight",
			"type": "Test Type",
			"province": "Test Province",
			"cityCountyDistrict": "Test CityCountyDistrict",
			"legalDong": "Test LegalDong",
			"ri": "Test Ri",	
			"streetNumber": "Test StreetNumber",
			"buildingNumber": "Test BuildingNumber",
			"longitude": 1.0,
			"latitude": 1.0
		 }`
		req := httptest.NewRequest("POST", "/sights", bytes.NewBuffer([]byte(body)))
		req.Header.Set("Content-Type", "application/json")
		service.On("InsertSight", mock.Anything).Return(&entities.Sight{}, errors.New("500 error")).Once()
		resp, _ := app.Test(req)

		assert.Equal(t, 500, resp.StatusCode)
		service.AssertExpectations(t)
	})

	t.Run("성공 시, 200 반환", func(t *testing.T) {
		body := `{
			"name": "Test Sight",
			"type": "Test Type",
			"province": "Test Province",
			"cityCountyDistrict": "Test CityCountyDistrict",
			"legalDong": "Test LegalDong",
			"ri": "Test Ri",	
			"streetNumber": "Test StreetNumber",
			"buildingNumber": "Test BuildingNumber",
			"longitude": 1.0,
			"latitude": 1.0
		 }`
		req := httptest.NewRequest("POST", "/sights", bytes.NewBuffer([]byte(body)))
		req.Header.Set("Content-Type", "application/json")
		sight := entities.Sight{
			Name:               "Test Sight",
			Type:               "Test Type",
			Province:           "Test Province",
			CityCountyDistrict: "Test CityCountyDistrict",
			LegalDong:          "Test LegalDong",
			Ri:                 "Test Ri",
			StreetNumber:       "Test StreetNumber",
			BuildingNumber:     "Test BuildingNumber",
			Longitude:          1.0,
			Latitude:           1.0,
		}
		service.On("InsertSight", mock.Anything).Return(&sight, nil)
		resp, _ := app.Test(req)

		assert.Equal(t, 200, resp.StatusCode)
		service.AssertExpectations(t)
	})
}

func TestUpdateSight(t *testing.T) {
	app := fiber.New()

	service := new(MockSightService)
	app.Put("/sights", UpdateSight(service))

	t.Run("유효하지 않은 request body 인 경우, 400 에러 반환", func(t *testing.T) {
		req := httptest.NewRequest("PUT", "/sights", bytes.NewBuffer([]byte("invalid_json")))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, 400, resp.StatusCode)
	})

	t.Run("내부 서버 오류의 경우, 500 에러 반환", func(t *testing.T) {
		body := `{
			"id": 1,
			"name": "Test Sight",
			"type": "Test Type",
			"province": "Test Province",
			"cityCountyDistrict": "Test CityCountyDistrict",
			"legalDong": "Test LegalDong",
			"ri": "Test Ri",	
			"streetNumber": "Test StreetNumber",
			"buildingNumber": "Test BuildingNumber",
			"longitude": 1.0,
			"latitude": 1.0
		 }`
		req := httptest.NewRequest("PUT", "/sights", bytes.NewBuffer([]byte(body)))
		req.Header.Set("Content-Type", "application/json")
		service.On("UpdateSight", mock.Anything).Return(&entities.Sight{}, errors.New("500 error")).Once()
		resp, _ := app.Test(req)

		assert.Equal(t, 500, resp.StatusCode)
		service.AssertExpectations(t)
	})

	t.Run("성공 시, 200 반환", func(t *testing.T) {
		body := `{
			"id": 1,
			"name": "Test Sight",
			"type": "Test Type",
			"province": "Test Province",
			"cityCountyDistrict": "Test CityCountyDistrict",
			"legalDong": "Test LegalDong",
			"ri": "Test Ri",	
			"streetNumber": "Test StreetNumber",
			"buildingNumber": "Test BuildingNumber",
			"longitude": 1.0,
			"latitude": 1.0
		 }`
		req := httptest.NewRequest("PUT", "/sights", bytes.NewBuffer([]byte(body)))
		req.Header.Set("Content-Type", "application/json")
		sight := entities.Sight{
			Name:               "Test Sight",
			Type:               "Test Type",
			Province:           "Test Province",
			CityCountyDistrict: "Test CityCountyDistrict",
			LegalDong:          "Test LegalDong",
			Ri:                 "Test Ri",
			StreetNumber:       "Test StreetNumber",
			BuildingNumber:     "Test BuildingNumber",
			Longitude:          1.0,
			Latitude:           1.0,
		}
		service.On("UpdateSight", mock.Anything).Return(&sight, nil)
		resp, _ := app.Test(req)
		assert.Equal(t, 200, resp.StatusCode)
		service.AssertExpectations(t)
	})
}

func TestRemoveSight(t *testing.T) {
	app := fiber.New()

	service := new(MockSightService)
	app.Delete("/sights", RemoveSight(service))

	t.Run("유효하지 않은 request body 인 경우, 400 에러 반환", func(t *testing.T) {
		req := httptest.NewRequest("DELETE", "/sights", bytes.NewBuffer([]byte("invalid_json")))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, 400, resp.StatusCode)
	})

	t.Run("내부 서버 오류의 경우, 500 에러 반환", func(t *testing.T) {
		body := `{"ID": "12345"}`
		req := httptest.NewRequest("DELETE", "/sights", bytes.NewBuffer([]byte(body)))
		req.Header.Set("Content-Type", "application/json")
		service.On("RemoveSight", mock.Anything).Return(errors.New("500 error")).Once()
		resp, _ := app.Test(req)

		assert.Equal(t, 500, resp.StatusCode)
		service.AssertExpectations(t)
	})

	t.Run("성공 시, 200 반환", func(t *testing.T) {
		body := `{"ID": "12345"}`
		req := httptest.NewRequest("DELETE", "/sights", bytes.NewBuffer([]byte(body)))
		req.Header.Set("Content-Type", "application/json")
		service.On("RemoveSight", mock.Anything).Return(nil)
		resp, _ := app.Test(req)

		assert.Equal(t, 200, resp.StatusCode)
		service.AssertExpectations(t)
	})
}

func TestGetSights(t *testing.T) {
	app := fiber.New()

	service := new(MockSightService)
	app.Get("/sights", GetSights(service))

	t.Run("내부 서버 오류의 경우, 500 에러 반환", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/sights", nil)
		service.On("FetchSights").Return(&[]response.Sight{}, errors.New("500 error")).Once()
		resp, _ := app.Test(req)

		assert.Equal(t, 500, resp.StatusCode)
		service.AssertExpectations(t)
	})

	t.Run("성공 시, 200 반환", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/sights", nil)
		sights := []response.Sight{
			{
				ID:                 1,
				Name:               "Test Sight",
				Type:               "Test Type",
				Province:           "Test Province",
				CityCountyDistrict: "Test CityCountyDistrict",
				LegalDong:          "Test LegalDong",
				Ri:                 "Test Ri",
				StreetNumber:       "Test StreetNumber",
				BuildingNumber:     "Test BuildingNumber",
				Longitude:          1.0,
				Latitude:           1.0,
			},
		}
		service.On("FetchSights").Return(&sights, nil).Once()
		resp, _ := app.Test(req)

		assert.Equal(t, 200, resp.StatusCode)
		service.AssertExpectations(t)
	})
}

func TestLoadSight(t *testing.T) {
	app := fiber.New()

	service := new(MockSightService)
	app.Get("/sights/location", LoadSight(service))

	t.Run("유효하지 않은 request body 인 경우, 400 에러 반환", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/sights/location", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, 400, resp.StatusCode)
	})

	t.Run("내부 서버 오류의 경우, 500 에러 반환", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/sights/location?Latitude=1.0&Longitude=1.0", nil)
		service.On("LoadSight", mock.Anything, mock.Anything).Return(&[]response.SightLoad{}, errors.New("500 error")).Once()
		resp, _ := app.Test(req)

		assert.Equal(t, 500, resp.StatusCode)
		service.AssertExpectations(t)
	})

	t.Run("성공 시, 200 반환", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/sights/location?Latitude=1.0&Longitude=1.0", nil)
		sights := []response.SightLoad{
			{
				ID:   1,
				Name: "Test Sight",
			},
		}
		service.On("LoadSight", mock.Anything, mock.Anything).Return(&sights, nil).Once()
		resp, _ := app.Test(req)

		assert.Equal(t, 200, resp.StatusCode)
		service.AssertExpectations(t)
	})
}
