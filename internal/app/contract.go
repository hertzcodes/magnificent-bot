package app

import (
	"log/slog"

	"github.com/hertzCodes/magnificent-bot/config"
	"gorm.io/gorm"
)

type App interface {
	DB() *gorm.DB
	Config() config.Config
	Commands() Commands
	Logger() *slog.Logger
}
