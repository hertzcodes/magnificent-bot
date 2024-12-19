package app

import (
	"log"
	"log/slog"

	"github.com/bwmarrin/discordgo"
	"github.com/hertzCodes/magnificent-bot/config"
	"github.com/hertzCodes/magnificent-bot/pkg/adapters/commands"
	"github.com/hertzCodes/magnificent-bot/pkg/logger"
	"github.com/hertzCodes/magnificent-bot/pkg/postgres"
	"gorm.io/gorm"
)

type app struct {
	db       *gorm.DB
	cfg      config.Config
	logger   *slog.Logger
	commands Commands
}

type Commands struct {
	Global []*discordgo.ApplicationCommand
	Admin  []*discordgo.ApplicationCommand
}

func (a *app) DB() *gorm.DB {
	return a.db
}

func (a *app) Config() config.Config {
	return a.cfg
}

func (a *app) Commands() Commands {
	return a.commands
}

func (a *app) Logger() *slog.Logger {
	return a.logger
}

func (a *app) setDB() error {
	db, err := postgres.NewPsqlGormConnection(

		postgres.DBConnOptions{
			User:     a.cfg.DB.User,
			Password: a.cfg.DB.Password,
			Host:     a.cfg.DB.Host,
			Port:     a.cfg.DB.Port,
			DBName:   a.cfg.DB.Database,
			Schema:   a.cfg.DB.Schema,
		})

	if err != nil {
		return err
	}

	a.db = db
	return nil
}

func (a *app) setCommands() {

	a.commands = Commands{
		Global: commands.RegisterCommandsGlobal(),
		Admin:  commands.RegisterCommandsAdmin(),
	}
}

func NewApp(cfg config.Config) (App, error) {
	a := &app{
		cfg:    cfg,
		logger: logger.NewLogger(),
	}

	if err := a.setDB(); err != nil {
		return nil, err
	}

	a.setCommands()
	// ... other initialization code ...

	return a, nil

}

func NewMustApp(cfg config.Config) App {
	app, err := NewApp(cfg)
	if err != nil {
		log.Fatal(err)
	}
	return app

}
