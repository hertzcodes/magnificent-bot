package basics

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Ping(logger *slog.Logger) interface{} {

	return func(bot *discordgo.Session, i *discordgo.InteractionCreate) {

		data := i.ApplicationCommandData()

		if data.Name != "ping" {
			return
		}

		err := Pong(bot, i)

		if err != nil {
			logger.Error(err.Error(), "command", "ping", "user", data.TargetID)
		}

	}
}
func Pong(bot *discordgo.Session, i *discordgo.InteractionCreate) error {

	err := bot.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("🏓 Pong **%s!** ", bot.HeartbeatLatency().Round(time.Millisecond)),
		},
	})

	return err
}
