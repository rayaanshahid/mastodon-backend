package main

import (
	"context"
	"log"
	"net/http"

	"github.com/mastodon-backend/clients"
	"github.com/mastodon-backend/config"
	"github.com/mastodon-backend/websockets"
	"github.com/mattn/go-mastodon"
)

func main() {
	configuration := config.Load()

	// Configure Mastodon client
	client := clients.NewClient(&mastodon.Config{
		Server:       "https://mastodon.social",
		ClientID:     configuration.ClientID,
		ClientSecret: configuration.ClientSecret,
		AccessToken:  configuration.AccessToken,
	})

	// Set up Websocket
	go websockets.PublishMessage()

	http.HandleFunc("/ws", websockets.HandleConnections)

	// Call the mastodon streaming API
	go func() {
		streamCtx, cancel := context.WithCancel(context.Background())
		defer cancel()

		err := client.StreamingPublic(streamCtx, false)

		if err != nil {
			log.Println(err)
		}
	}()

	// Start the server
	log.Println("Server started on :" + configuration.ServerPort)
	err := http.ListenAndServe(":"+configuration.ServerPort, nil)
	if err != nil {
		log.Println("ListenAndServe: ", err)
	}
}
