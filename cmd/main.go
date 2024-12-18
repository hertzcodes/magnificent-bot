package main

import (
	"flag"
	"os"

	"github.com/hertzCodes/magnificent-bot/config"
	"github.com/hertzCodes/magnificent-bot/internal/app"
)

var configPath = flag.String("config", "config.json", "application configuration")

func main() {
	flag.Parse()

	if path := os.Getenv("CONFIG_PATH"); len(path) > 0 {
		*configPath = path
	}

	c := config.MustReadConfig(*configPath)

	appContainer := app.NewMustApp(c)

}
