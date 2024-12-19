package discord

import "github.com/bwmarrin/discordgo"

func SendChannelMessage(text string, bot *discordgo.Session, i *discordgo.InteractionCreate) error {

	err := bot.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			AllowedMentions: &discordgo.MessageAllowedMentions{RepliedUser: true},
			Content:         text,
		},
	})

	return err

}

func SendChannelMessageDeferred(text string, bot *discordgo.Session, i *discordgo.InteractionCreate) error {

	err := bot.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: text,
		},
	})

	return err

}

func SendChannelMessageEmbed(text string, embeds []*discordgo.MessageEmbed, bot *discordgo.Session, i *discordgo.InteractionCreate) error {

	err := bot.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			AllowedMentions: &discordgo.MessageAllowedMentions{RepliedUser: true},
			Content:         text,
			Embeds:          embeds,
		},
	})

	return err

}

func SendChannelMessageEmbedDeferred(text string, embeds []*discordgo.MessageEmbed, bot *discordgo.Session, i *discordgo.InteractionCreate) error {

	err := bot.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content:         text,
			Embeds:          embeds,
		},
	})

	return err

}