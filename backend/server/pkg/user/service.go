package user

import "server/pkg/entities"

type Service interface {
	InsertUser(user *entities.User) (*entities.User, error)
	UpdateUser(user *entities.User) (*entities.User, error)
	RemoveUser(ID string) error
	CheckEmail(email string) (bool, error)
	CheckNickname(nickname string) (bool, error)
	GetUserById(ID string) (*entities.User, error)
	GetUsers() (*[]entities.User, error)
	Login(email string, password string) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) InsertUser(user *entities.User) (*entities.User, error) {
	return s.repository.CreateUser(user)
}

func (s *service) UpdateUser(user *entities.User) (*entities.User, error) {
	return s.repository.UpdateUser(user)
}

func (s *service) RemoveUser(ID string) error {
	return s.repository.DeleteUser(ID)
}

func (s *service) CheckEmail(email string) (bool, error) {
	return s.repository.CheckEmail(email)
}

func (s *service) CheckNickname(nickname string) (bool, error) {
	return s.repository.CheckNickname(nickname)
}

func (s *service) GetUserById(ID string) (*entities.User, error) {
	return s.repository.GetUserByID(ID)
}

func (s *service) GetUsers() (*[]entities.User, error) {
	return s.repository.GetUsers()
}

// TODO: JWT 추가
func (s *service) Login(email string, password string) (bool, error) {
	return s.repository.CheckUser(email, password)
}
