package post

import (
	"backend/api/presenter/response"
	"backend/pkg/entities"
)

type Service interface {
	InsertPost(post *entities.Post) (*entities.Post, error)
	FetchPosts() (*[]response.Post, error)
	UpdatePost(post *entities.Post) (*entities.Post, error)
	RemovePost(ID string) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) InsertPost(post *entities.Post) (*entities.Post, error) {
	return s.repository.CreatePost(post)
}

func (s *service) FetchPosts() (*[]response.Post, error) {
	return s.repository.ReadPost()
}

func (s *service) UpdatePost(post *entities.Post) (*entities.Post, error) {
	return s.repository.UpdatePost(post)
}

func (s *service) RemovePost(ID string) error {
	return s.repository.DeletePost(ID)
}
