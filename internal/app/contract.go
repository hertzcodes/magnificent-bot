package app

import (
	"github.com/bwmarrin/discordgo"
	"github.com/hertzCodes/magnificent-bot/config"
	"gorm.io/gorm"
)

type App interface {
	DB() *gorm.DB
	Config() config.Config
	Commands() []*discordgo.ApplicationCommand
}
