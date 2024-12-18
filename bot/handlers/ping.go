package handlers

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Ping(bot *discordgo.Session, i *discordgo.InteractionCreate) {

	data := i.ApplicationCommandData()

	if data.Name != "ping" {
		return
	}

	Pong(bot, i)

}

func Pong(bot *discordgo.Session, i *discordgo.InteractionCreate) {

	err := bot.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("🏓 Pong **%s** ", bot.HeartbeatLatency().Round(time.Millisecond)),
		},
	})

	if err != nil {
		log.Println(err)
		return
	}
}
