package app

import (
	"log"

	"github.com/hertzCodes/magnificent-bot/pkg/postgres"
	"github.com/hertzCodes/magnificent-bot/config"
	"gorm.io/gorm"
)

type app struct {
	db  *gorm.DB
	cfg config.Config
}

func (a *app) DB() *gorm.DB {
	return a.db
}

func (a *app) Config() config.Config {
	return a.cfg
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

func NewApp(cfg config.Config) (App, error) {
	a := &app{
		cfg: cfg,
	}

	if err := a.setDB(); err != nil {
		return nil, err
	}

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
