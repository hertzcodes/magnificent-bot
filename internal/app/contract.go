package app

import (
	"gorm.io/gorm"
	"github.com/hertzCodes/magnificent-bot/config"
)

type App interface {
	DB() *gorm.DB
	Config() config.Config
}
