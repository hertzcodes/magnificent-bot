package mapper

import (
	"github.com/hertzCodes/magnificent-bot/internal/user/domain"
	"github.com/hertzCodes/magnificent-bot/pkg/adapters/storage/entities"
	"gorm.io/gorm"
)

func UserDomainToStorage(userDomain domain.User) *entities.User {
	return &entities.User{
		Model: gorm.Model{
			ID:        uint(userDomain.ID),
			CreatedAt: userDomain.CreatedAt,
			DeletedAt: gorm.DeletedAt(ToNullTime(userDomain.DeletedAt)),
		},
		DiscordID:       uint64(userDomain.DiscordID),
		DiscordUsername: userDomain.DiscordUsername,
		RobloxID:        uint64(userDomain.RobloxID),
		RobloxUsername:  userDomain.RobloxUsername,
		Role:            uint64(userDomain.Role),
		IsPremium:       userDomain.IsPremium,
	}
}

func UserStorageToDomain(user entities.User) *domain.User {
	return &domain.User{
		ID:              domain.UserID(user.ID),
		CreatedAt:       user.CreatedAt,
		DeletedAt:       user.CreatedAt,
		DiscordID:       domain.DiscordID(user.DiscordID),
		DiscordUsername: user.DiscordUsername,
		RobloxID:        domain.RobloxID(user.RobloxID),
		RobloxUsername:  user.RobloxUsername,
		Role:            domain.Role(user.Role),
		IsPremium:       user.IsPremium,
	}
}
