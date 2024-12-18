package types

import "github.com/bwmarrin/discordgo"

type Command struct {
	Active bool
	Config discordgo.ApplicationCommand
}
