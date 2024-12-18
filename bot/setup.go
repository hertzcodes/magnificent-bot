package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/hertzCodes/magnificent-bot/config"
	"github.com/hertzCodes/magnificent-bot/bot/handlers"
	"github.com/hertzCodes/magnificent-bot/internal/app"
)

func Run(appContainer app.App, cfg config.BotConfig) error {

	bot, err := discordgo.New(cfg.Token)
	if err != nil {
		log.Fatal(err)
	}

	// registerHandlers(appContainer, bot)
	err = bot.Open()

	return err

}

func registerHandlers(appContainer app.App, bot *discordgo.Session, i *discordgo.InteractionCreate) {
	bot.AddHandler(handlers.BotReady) // CHECK LATER
}
