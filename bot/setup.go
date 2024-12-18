package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/hertzCodes/magnificent-bot/bot/handlers"
	"github.com/hertzCodes/magnificent-bot/config"
	"github.com/hertzCodes/magnificent-bot/internal/app"
)

func Run(appContainer app.App, cfg config.BotConfig) *discordgo.Session {

	bot, err := discordgo.New(cfg.Token)
	if err != nil {
		log.Fatal(err)
	}

	registerHandlers(appContainer, bot)
	_, err = bot.ApplicationCommandBulkOverwrite(cfg.AppID, "1205824707874656257", appContainer.Commands())
	if err != nil {
		log.Fatal("failed ", err)
	}

	err = bot.Open()

	return bot
}

func registerHandlers(appContainer app.App, bot *discordgo.Session) {
	bot.AddHandler(handlers.Ping)
	bot.AddHandler(handlers.BotReady)
}
