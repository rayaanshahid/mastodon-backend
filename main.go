package main

import (
	"context"
	"fmt"
	"log"

	//"net/http"
	"os"
	"sync"

	//"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"github.com/mattn/go-mastodon"
)

//var clients = make(map[*websocket.Conn]bool)

//var broadcast = make(chan mastodon.Status)

//var upgrader = websocket.Upgrader{}

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

	//go handleMessages()

	//http.HandleFunc("/ws", handleConnections)

	// Start the server
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
				//broadcast <- *e.Status
			case *mastodon.DeleteEvent:
				fmt.Printf("Status deleted: %s\n", e.ID)
			case *mastodon.NotificationEvent:
				fmt.Printf("Notification: %s\n", e.Notification.Type)
			default:
				fmt.Printf("Unknown event: %T\n", e)
			}
		}
		/*switch event := event.(type) {
		case *mastodon.UpdateEvent:
			broadcast <- *event.Status
		}*/

		if err != nil {
			log.Fatal(err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()

	/*err = http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}*/
}

/*func handleConnections(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()
	clients[ws] = true

	for {
		var msg mastodon.Status
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}*/
