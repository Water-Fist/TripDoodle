package sight

import (
	"backend/api/presenter"
	"backend/pkg/entities"
)

type Service interface {
	InsertSight(sight *entities.Sight) (*entities.Sight, error)
	FetchSights() (*[]presenter.Sight, error)
	UpdateSight(sight *entities.Sight) (*entities.Sight, error)
	RemoveSight(ID string) error
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

func (s *service) FetchSights() (*[]presenter.Sight, error) {
	return s.repository.ReadSight()
}

func (s *service) UpdateSight(sight *entities.Sight) (*entities.Sight, error) {
	return s.repository.UpdateSight(sight)
}

func (s *service) RemoveSight(ID string) error {
	return s.repository.DeleteSight(ID)
}
