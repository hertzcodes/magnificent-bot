package types

import "github.com/bwmarrin/discordgo"

var HelpEmbed *discordgo.MessageEmbed = &discordgo.MessageEmbed{
	Title:       "Commands",
	Color:       15626743,
	Description: "List of slash commands:",
	Fields: []*discordgo.MessageEmbedField{
		&discordgo.MessageEmbedField{
			Name:  "/ping",
			Value: "See the bot's response time",
		},

		&discordgo.MessageEmbedField{
			Name:  "/value `[game]` `[item-name]` `[options]`",
			Value: "See an item's value",
		},

		&discordgo.MessageEmbedField{
			Name:  "/verify `[username]`",
			Value: "Verify your roblox username",
		},
	},
}
