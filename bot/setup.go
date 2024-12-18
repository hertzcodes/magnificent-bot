package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/hertzCodes/magnificent-bot/config"
	"github.com/hertzCodes/magnificent-bot/internal/app"
)

func Run(appContainer app.App, cfg config.BotConfig) error {

	bot, err := discordgo.New(cfg.Token)
	if err != nil {
		log.Fatal(err)
	}

	err = bot.Open()

	return err

}

func RegisterCommands(bot *discordgo.Session, i *discordgo.InteractionCreate) {
	
}
