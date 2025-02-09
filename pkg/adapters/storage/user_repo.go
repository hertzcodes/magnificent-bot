package storage

import (
	"github.com/hertzCodes/magnificent-bot/internal/user/domain"
	"github.com/hertzCodes/magnificent-bot/internal/user/port"
	"github.com/hertzCodes/magnificent-bot/pkg/adapters/storage/entities"
	"github.com/hertzCodes/magnificent-bot/pkg/adapters/storage/mapper"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) port.Repo {
	return &userRepo{db}
}

func (r *userRepo) Create(userdomain domain.User) (domain.UserID, error) {
	user := mapper.UserDomainToStorage(userdomain)
	return domain.UserID(user.ID), r.db.Table("users").Create(user).Error
}

func (r *userRepo) Update(userdomain domain.User) error {
	user := mapper.UserDomainToStorage(userdomain)

	q := r.db.Table("users").Where("discord_id = ?", user.DiscordID)

	err := q.Updates(user).Error

	return err
}

func (r *userRepo) GetByDiscordID(id domain.DiscordID) (*domain.User, error) {
	var user entities.User

	q := r.db.Table("users").Where("discord_id = ?", id)

	err := q.First(&user).Error

	if err != nil {
		return nil, err
	}

	return mapper.UserStorageToDomain(user), nil

}
