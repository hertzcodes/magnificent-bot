package basics

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func BotReady(bot *discordgo.Session, r *discordgo.Ready) {
	log.Printf("Logged in as %s", r.User.String())
}
