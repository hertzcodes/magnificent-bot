package commands

import (
	"github.com/hertzCodes/magnificent-bot/pkg/adapters/commands/types"
)

// AddCommand adds a command if active = true
// dm param enables dm permission for bot
func AddCommand(commands []*types.Command, active bool, dm bool, f func(bool, bool) *types.Command) {
	if active {
		commands = append(commands, f(true, true))
	}
}

func RegisterCommands() []*types.Command {
	c := make([]*types.Command, 1)

	AddCommand(c, true, true, types.NewPingCommand)
	// Add more commands here

	return c
}
