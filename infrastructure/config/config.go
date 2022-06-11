package config

import (
	"os"

	"github.com/joho/godotenv"
)

type SlackConfig struct {
	SlackWebhook string `env:"SLACKWEBHOOK"`
	BotToken     string `env:"BOT_TOKEN"`
	ChId         string `env:"CH_ID"`
}

func GetSlackEnv() (SlackConfig, error) {
	godotenv.Load(".env")
	cfg := SlackConfig{
		SlackWebhook: os.Getenv("SLACKWEBHOOK"),
		BotToken:     os.Getenv("BOT_TOKEN"),
		ChId:         os.Getenv("CH_ID"),
	}
	return cfg, nil
}
