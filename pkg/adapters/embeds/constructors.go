package embeds

import (
	"github.com/bwmarrin/discordgo"
	"github.com/hertzCodes/magnificent-bot/pkg/adapters/embeds/types"
)

func NewHelpEmbed() *discordgo.MessageEmbed {
	return types.HelpEmbed
}
