package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/hertzCodes/magnificent-bot/pkg/adapters/commands/types"
)

// AddCommand adds a command if active = true
// dm param enables dm permission for bot
func AddCommand(commands *[]*discordgo.ApplicationCommand, active bool, dm bool, f func(bool, bool) *types.Command) {
	if active {
		*commands = append(*commands, &f(true, true).Config)
	}
}

func RegisterCommands() []*discordgo.ApplicationCommand {
	c := make([]*discordgo.ApplicationCommand, 0)

	AddCommand(&c, true, true, types.NewPingCommand)
	// Add more commands here

	return c
}
