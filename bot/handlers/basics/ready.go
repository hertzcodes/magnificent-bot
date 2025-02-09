package basics

import (
	"fmt"
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

func BotReady(logger *slog.Logger) interface{} {

	return func(bot *discordgo.Session, r *discordgo.Ready) {
		logger.Info(fmt.Sprintf("Logged in as %s", r.User.String()))
	}
}
