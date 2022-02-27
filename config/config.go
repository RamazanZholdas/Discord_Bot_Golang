package config

import (
	"encoding/json"
	"os"
)

var (
	Token     string
	BotPrefix string

	config *Config
)

type Config struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

func ReadConfig() error {
	byteSlice, err := os.ReadFile("./config.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteSlice, &config)
	if err != nil {
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil
}
