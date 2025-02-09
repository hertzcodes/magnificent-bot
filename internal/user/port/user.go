package port

import "github.com/hertzCodes/magnificent-bot/internal/user/domain"

type Repo interface {
	Create(user domain.User) (domain.UserID, error)
	Update(user domain.User) error
	GetByDiscordID(id domain.DiscordID) (*domain.User, error)
}
