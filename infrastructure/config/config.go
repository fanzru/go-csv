package config

import (
	"os"

	"github.com/joho/godotenv"
)

type SlackConfig struct {
	SlackWebhook string `env:"SLACKWEBHOOK"`
}

func GetSlackEnv() (SlackConfig, error) {
	godotenv.Load(".env")
	cfg := SlackConfig{
		SlackWebhook: os.Getenv("SLACKWEBHOOK"),
	}
	return cfg, nil
}
