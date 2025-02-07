package handlers

import (
	"github.com/bwmarrin/discordgo"
)

func GetUserDiscordID(i *discordgo.Interaction) string {
	user := i.Member.User.ID

	if user == "" {
		user = i.User.ID
	}

	return user

}
