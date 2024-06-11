package mastodon

import (
	"context"
	"log"

	"github.com/mastodon-backend/config"
	"github.com/mattn/go-mastodon"
)

func StartMastodonApp(config config.Config) {

	// Configure Mastodon client
	client := NewClient(&mastodon.Config{
		Server:       "https://mastodon.social",
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		AccessToken:  config.AccessToken,
	})

	// Start streaming
	go StartStreaming(client)
}

func StartStreaming(client *MastodonClient) {
	streamCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := client.StreamingPublic(streamCtx, false)

	if err != nil {
		log.Println(err)
	}
}
