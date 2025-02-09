package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/hertzCodes/magnificent-bot/bot/handlers/basics"
	"github.com/hertzCodes/magnificent-bot/config"
	"github.com/hertzCodes/magnificent-bot/internal/app"
)

func Run(appContainer app.App, cfg config.BotConfig) *discordgo.Session {

	bot, err := discordgo.New(cfg.Token)
	if err != nil {
		log.Fatal(err)
	}

	registerHandlers(appContainer, bot)
	_, err = bot.ApplicationCommandBulkOverwrite(cfg.AppID, "", appContainer.Commands().Global)
	if err != nil {
		log.Fatal("[GLOBAL COMMANDS OVERWRITE]", err)
	}

	for _, cmd := range appContainer.Commands().Admin {
		_, err = bot.ApplicationCommandCreate(cfg.AppID, cfg.GuildID, cmd)

		if err != nil {
			appContainer.Logger().Error(err.Error(), "command", cmd.Name)
		}
	}

	err = bot.Open()

	return bot
}

func registerHandlers(appContainer app.App, bot *discordgo.Session) {
	bot.AddHandler(basics.Ping(appContainer.Logger()))
	bot.AddHandler(basics.BotReady(appContainer.Logger()))
	bot.AddHandler(basics.Help(appContainer.Logger()))
}
