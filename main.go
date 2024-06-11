package main

import (
	"log"

	"github.com/mastodon-backend/config"
	mastodonpkg "github.com/mastodon-backend/pkg/mastodon"
	"github.com/mastodon-backend/websockets"
)

func main() {
	config := config.Load()

	mastodonpkg.StartMastodonApp(config)

	if err := websockets.StartWebSocket(config); err != nil {
		log.Println("ListenAndServe: ", err)
	}
}
