package types

import "github.com/bwmarrin/discordgo"

func NewHelpCommand(active bool, dm bool) *Command {

	var (
		p int64 = discordgo.PermissionSendMessages
	)

	return &Command{
		Active: active,
		Config: discordgo.ApplicationCommand{
			Name:                     "help",
			DMPermission:             &dm,
			DefaultMemberPermissions: &p,
			Type:                     discordgo.ChatApplicationCommand,
			Description:              "shows bot's guide for commands",
		},
	}
}
