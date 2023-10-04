package sight

import (
	"server/api/presenter/request"
	"server/api/presenter/response"
	"server/pkg/entities"
)

type Service interface {
	InsertSight(sight *request.SightRequest) (*entities.Sight, error)
	FetchSights() (*[]response.Sight, error)
	UpdateSight(sight *entities.Sight) (*entities.Sight, error)
	RemoveSight(ID string) error
	LoadSight(Latitude float32, Longitude float32) (*[]response.SightLoad, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) InsertSight(sight *request.SightRequest) (*entities.Sight, error) {
	body := &entities.Sight{
		Name:               sight.Name,
		Type:               sight.Type,
		Province:           sight.Province,
		CityCountyDistrict: sight.CityCountyDistrict,
		LegalDong:          sight.LegalDong,
		Ri:                 sight.Ri,
		StreetNumber:       sight.StreetNumber,
		BuildingNumber:     sight.BuildingNumber,
		Latitude:           sight.Latitude,
		Longitude:          sight.Longitude,
	}
	return s.repository.CreateSight(body)
}

func (s *service) FetchSights() (*[]response.Sight, error) {
	return s.repository.ReadSight()
}

func (s *service) UpdateSight(sight *entities.Sight) (*entities.Sight, error) {
	return s.repository.UpdateSight(sight)
}

func (s *service) RemoveSight(ID string) error {
	return s.repository.DeleteSight(ID)
}

func (s *service) LoadSight(Latitude float32, Longitude float32) (*[]response.SightLoad, error) {
	return s.repository.LoadSight(Latitude, Longitude)
}
