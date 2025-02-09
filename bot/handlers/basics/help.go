package basics

import (
	"log/slog"

	"github.com/bwmarrin/discordgo"
	"github.com/hertzCodes/magnificent-bot/bot/handlers"
	"github.com/hertzCodes/magnificent-bot/pkg/adapters/embeds"
	"github.com/hertzCodes/magnificent-bot/pkg/discord"
)

func Help(logger *slog.Logger) interface{} {
	return func(bot *discordgo.Session, i *discordgo.InteractionCreate) {

		data := i.ApplicationCommandData()

		if data.Name != "help" {
			return
		}

		embeds := []*discordgo.MessageEmbed{
			embeds.NewHelpEmbed(),
		}

		err := discord.SendChannelMessageEmbed("", embeds, bot, i)

		if err != nil {
			logger.Error(err.Error(), "command", "help", "user", handlers.GetUserDiscordID(i.Interaction))
		}

	}
}
