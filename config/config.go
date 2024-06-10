package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort   string
	ClientID     string
	ClientSecret string
	AccessToken  string
}

func Load() Config {
	godotenv.Load()
	return Config{
		ServerPort:   os.Getenv("SERVER_PORT"),
		ClientID:     os.Getenv("MASTODON_CLIENT_ID"),
		ClientSecret: os.Getenv("MASTODON_CLIENT_SECRET"),
		AccessToken:  os.Getenv("MASTODON_ACCESS_TOKEN"),
	}
}
