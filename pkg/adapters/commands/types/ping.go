package types

import (
	"github.com/bwmarrin/discordgo"
)

func NewPingCommand(dm bool) *discordgo.ApplicationCommand {

	var (
		p int64 = discordgo.PermissionReadMessageHistory + discordgo.PermissionSendMessages
	)

	return &discordgo.ApplicationCommand{
		Name:                     "ping",
		Description:              "shows bot's latency",
		DefaultMemberPermissions: &p,
		DMPermission:             &dm,
		Type:                     discordgo.ChatApplicationCommand,
	}
}
