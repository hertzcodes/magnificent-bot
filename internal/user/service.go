package user

import (
	"errors"

	"github.com/hertzCodes/magnificent-bot/internal/user/domain"
	"github.com/hertzCodes/magnificent-bot/internal/user/port"
)

var (
	ErrUserOnCreate = errors.New("failed to create user")
	ErrUserOnUpdate = errors.New("failed to update user")
	ErrUserNotFound = errors.New("user not found")
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateUser(user domain.User) (domain.UserID, error) {
	userID, err := s.repo.Create(user)

	return userID, err
}

func (s *service) UpdateUser(user domain.User) error {
	err := s.repo.Update(user)

	return err
}

func (s *service) GetUserByDiscordID(id domain.DiscordID) (*domain.User, error) {
	user, err := s.repo.GetByDiscordID(id)

	return user, err
}
