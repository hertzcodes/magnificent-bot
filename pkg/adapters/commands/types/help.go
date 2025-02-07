package types

import (
	"github.com/bwmarrin/discordgo"
)

func NewHelpCommand(dm bool) *discordgo.ApplicationCommand {

	var (
		p int64 = discordgo.PermissionSendMessages
	)

	return &discordgo.ApplicationCommand{
		Name:                     "help",
		DMPermission:             &dm,
		DefaultMemberPermissions: &p,
		Type:                     discordgo.ChatApplicationCommand,
		Description:              "shows bot's guide for commands",
	}
}
