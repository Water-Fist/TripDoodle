package sight

import (
	"server/api/presenter/response"
	"server/pkg/entities"
)

type Service interface {
	InsertSight(sight *entities.Sight) (*entities.Sight, error)
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

func (s *service) InsertSight(sight *entities.Sight) (*entities.Sight, error) {
	return s.repository.CreateSight(sight)
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
