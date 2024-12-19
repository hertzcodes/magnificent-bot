package basics

import (
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

func BotReady(logger *slog.Logger) interface{} {

	return func(bot *discordgo.Session, r *discordgo.Ready) {
		logger.Info("Logged in as %s", r.User.String())
	}
}
