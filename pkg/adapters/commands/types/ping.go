package types

import (
	"github.com/bwmarrin/discordgo"
)

func NewPingCommand(active bool, dm bool) *Command {

	var (
		p int64 = discordgo.PermissionReadMessageHistory + discordgo.PermissionSendMessages
	)

	return (&Command{
		Active: active,
		Config: discordgo.ApplicationCommand{
			Name:                     "ping",
			Description:              "shows bot's ping",
			DefaultMemberPermissions: &p,
			DMPermission:             &dm,
			Type:                     discordgo.ChatApplicationCommand,
		},
	})
}
