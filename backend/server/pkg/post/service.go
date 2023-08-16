package post

import (
	"server/api/presenter/request"
	"server/api/presenter/response"
	"server/pkg/entities"
)

type Service interface {
	InsertPost(post *request.PostRequest) (*entities.Post, error)
	FetchPosts() (*[]response.Post, error)
	UpdatePost(post *request.UpdatePostRequest) (*entities.Post, error)
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

func (s *service) InsertPost(post *request.PostRequest) (*entities.Post, error) {
	entityPost := &entities.Post{
		Title:    post.Title,
		Content:  post.Content,
		ImageUrl: post.ImageUrl,
		State:    post.State,
		SightId:  post.SightID,
	}
	return s.repository.CreatePost(entityPost)
}

func (s *service) FetchPosts() (*[]response.Post, error) {
	return s.repository.ReadPost()
}

func (s *service) UpdatePost(post *request.UpdatePostRequest) (*entities.Post, error) {
	entityPost := &entities.Post{
		ID:       post.ID,
		Title:    post.Title,
		Content:  post.Content,
		ImageUrl: post.ImageUrl,
		State:    post.State,
	}
	return s.repository.UpdatePost(entityPost)
}

func (s *service) RemovePost(ID string) error {
	return s.repository.DeletePost(ID)
}
