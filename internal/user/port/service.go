package port

import "github.com/hertzCodes/magnificent-bot/internal/user/domain"

type Service interface {
	CreateUser(user domain.User) (domain.UserID, error)
	UpdateUser(user domain.User) error
	GetUserByDiscordID(id domain.DiscordID) (*domain.User, error)
}
