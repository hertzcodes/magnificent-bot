package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/hertzCodes/magnificent-bot/pkg/adapters/commands/types"
)

// AddCommand adds a command if active = true
// dm param enables dm permission for bot
func AddCommand(commands *[]*discordgo.ApplicationCommand, active bool, dm bool, f func(bool) *discordgo.ApplicationCommand) {
	if active {
		*commands = append(*commands, f(dm))
	}
}

func RegisterCommandsAdmin() []*discordgo.ApplicationCommand {
	return nil
}

func RegisterCommandsGlobal() []*discordgo.ApplicationCommand {
	c := make([]*discordgo.ApplicationCommand, 0)

	AddCommand(&c, true, true, types.NewPingCommand)
	AddCommand(&c, true, true, types.NewHelpCommand)
	AddCommand(&c, true , true, types.NewVerifyCommand)
	// Add more commands here

	return c
}
