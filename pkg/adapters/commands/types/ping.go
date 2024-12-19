package types

import (
	"github.com/bwmarrin/discordgo"
)

func NewPingCommand(dm bool) *Command {

	var (
		p int64 = discordgo.PermissionReadMessageHistory + discordgo.PermissionSendMessages
	)

	return (&Command{
		Config: discordgo.ApplicationCommand{
			Name:                     "ping",
			Description:              "shows bot's latency",
			DefaultMemberPermissions: &p,
			DMPermission:             &dm,
			Type:                     discordgo.ChatApplicationCommand,
		},
	})
}
