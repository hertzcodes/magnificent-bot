package app

import (
	"log/slog"

	"github.com/bwmarrin/discordgo"
	"github.com/hertzCodes/magnificent-bot/config"
	"gorm.io/gorm"
)

type App interface {
	DB() *gorm.DB
	Config() config.Config
	Commands() []*discordgo.ApplicationCommand
	Logger() *slog.Logger
}
