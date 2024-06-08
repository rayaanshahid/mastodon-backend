package main

import (
	"context"
	"fmt"
	"log"

	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/mattn/go-mastodon"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Configure Mastodon client
	client := mastodon.NewClient(&mastodon.Config{
		Server:       "https://mastodon.social",
		ClientID:     "Ef9b9eHQ5jLuh7OVqcUop9RfTjGNejmBa0TKpq4B_bQ",
		ClientSecret: "eWSNU1fuPL2tYuffSLO6nLcCHPlFlnlRIB1sLXajzaA",
		AccessToken:  os.Getenv("MASTODON_ACCESS_TOKEN"),
	})

	log.Println("Server started on :8000")
	go func() {

		// Create a context for the streaming API
		streamCtx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ch, err := client.StreamingPublic(streamCtx, false)

		for event := range ch {
			switch e := event.(type) {
			case *mastodon.UpdateEvent:
				fmt.Printf("New status: %s\n", e.Status.Content)
			case *mastodon.DeleteEvent:
				fmt.Printf("Status deleted: %s\n", e.ID)
			case *mastodon.NotificationEvent:
				fmt.Printf("Notification: %s\n", e.Notification.Type)
			default:
				fmt.Printf("Unknown event: %T\n", e)
			}
		}

		if err != nil {
			log.Fatal(err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()

}
