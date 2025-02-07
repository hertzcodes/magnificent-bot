package basics

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/hertzCodes/magnificent-bot/bot/handlers"
	"github.com/hertzCodes/magnificent-bot/pkg/discord"
)

func Ping(logger *slog.Logger) interface{} {

	return func(bot *discordgo.Session, i *discordgo.InteractionCreate) {

		data := i.ApplicationCommandData()

		if data.Name != "ping" {
			return
		}

		msg := fmt.Sprintf("🏓 Pong **%s!** ", bot.HeartbeatLatency().Round(time.Millisecond))

		err := discord.SendChannelMessage(msg, bot, i)

		if err != nil {
			logger.Error(err.Error(), "command", "ping", "user", handlers.GetUserDiscordID(i.Interaction))
		}

	}
}
