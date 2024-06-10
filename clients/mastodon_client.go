package clients

import (
	"context"
	"log"

	"github.com/mastodon-backend/websockets"
	"github.com/mattn/go-mastodon"
)

type MastodonClient struct {
	Client *mastodon.Client
}

func NewClient(config *mastodon.Config) *MastodonClient {
	client := mastodon.NewClient(config)
	return &MastodonClient{client}
}

func (c *MastodonClient) StreamingPublic(ctx context.Context, public bool) error {
	ch, err := c.Client.StreamingPublic(ctx, false)
	if err != nil {
		log.Println("Error in StreamingPublic:", err)
		return err
	}
	for event := range ch {
		switch e := event.(type) {
		case *mastodon.UpdateEvent:
			websockets.Broadcast <- *e.Status
		case *mastodon.DeleteEvent:
			//log.Printf("Status deleted: %s\n", e.ID)
		case *mastodon.NotificationEvent:
			//log.Printf("Notification: %s\n", e.Notification.Type)
		default:
			//log.Printf("Unknown event: %T\n", e)
		}
	}
	return nil
}
