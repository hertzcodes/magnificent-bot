package app

import (
	"github.com/hertzCodes/magnificent-bot/config"
	"github.com/hertzCodes/magnificent-bot/pkg/adapters/commands/types"
	"gorm.io/gorm"
)

type App interface {
	DB() *gorm.DB
	Config() config.Config
	Commands() []*types.Command
}
