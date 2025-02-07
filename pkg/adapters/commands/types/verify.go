package types

import "github.com/bwmarrin/discordgo"

func NewVerifyCommand(dm bool) *discordgo.ApplicationCommand {

	var (
		p int64 = discordgo.PermissionSendMessages
	)

	return &discordgo.ApplicationCommand{
		Name:                     "verify",
		DMPermission:             &dm,
		DefaultMemberPermissions: &p,
		Type:                     discordgo.ChatApplicationCommand,
		Description:              "verify your roblox username",
		Options: []*discordgo.ApplicationCommandOption{
			&discordgo.ApplicationCommandOption{
				Type:         discordgo.ApplicationCommandOptionString,
				Name:         "username",
				Description:  "your roblox username",
				Required:     true,
				Autocomplete: false,
			},
		},
	}

}
