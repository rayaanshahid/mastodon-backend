package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/mastodon-backend/clients"
	"github.com/mastodon-backend/websockets"
	"github.com/mattn/go-mastodon"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Configure Mastodon client
	client := clients.NewClient(&mastodon.Config{
		Server:       "https://mastodon.social",
		ClientID:     os.Getenv("MASTODON_CLIENT_ID"),
		ClientSecret: os.Getenv("MASTODON_CLIENT_SECRET"),
		AccessToken:  os.Getenv("MASTODON_ACCESS_TOKEN"),
	})

	go websockets.PublishMessage()

	http.HandleFunc("/ws", websockets.HandleConnections)

	go func() {

		// Create a context for the streaming API
		streamCtx, cancel := context.WithCancel(context.Background())
		defer cancel()

		err := client.StreamingPublic(streamCtx, false)

		if err != nil {
			log.Println(err)
		}
	}()

	// Start the server
	err = http.ListenAndServe(":8000", nil)
	log.Println("Server started on :8000")
	if err != nil {
		log.Println("ListenAndServe: ", err)
	}
}
