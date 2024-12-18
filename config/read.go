package config

import (
	"encoding/json"
	"log"
	"os"
)

func ReadConfig(configPath string) (Config, error) {
	var c Config
	data, err := os.ReadFile(configPath)

	if err != nil {
		return c, err
	}

	return c, json.Unmarshal(data, &c)
}

func MustReadConfig(configPath string) Config {
	c, err := ReadConfig(configPath)

	if err != nil {
		log.Fatal(err)
	}
	return c
}
