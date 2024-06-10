package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"github.com/mastodon-backend/websockets"
	"github.com/mattn/go-mastodon"
)

var testClient *mastodon.Client

func TestWebSocketConnection(t *testing.T) {
	// Start test server
	server := httptest.NewServer(http.HandlerFunc(websockets.HandleConnections))
	defer server.Close()

	// Start the connection
	url := "ws" + server.URL[len("http"):] + "/ws"
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("dial error: %v", err)
	}
	defer ws.Close()

	//test the message send on the connection
	message := mastodon.Status{Content: "test"}
	if err := ws.WriteJSON(message); err != nil {
		t.Fatalf("write error: %v", err)
	}
}

func TestMain(t *testing.M) {
	// Load .env file
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("Error loading .env file", err)
	}
	log.Println(os.Getenv("MASTODON_ACCESS_TOKEN"))
	// Configure Mastodon client
	testClient = mastodon.NewClient(&mastodon.Config{
		Server:       "https://mastodon.social",
		AccessToken:  os.Getenv("MASTODON_ACCESS_TOKEN"),
		ClientID:     os.Getenv("MASTODON_CLIENT_ID"),
		ClientSecret: os.Getenv("MASTODON_CLIENT_SECRET"),
	})

	go websockets.PublishMessage()

	// Start test server
	server := httptest.NewServer(http.HandlerFunc(websockets.HandleConnections))
	defer server.Close()

	// Run tests
	os.Exit(t.Run())
}
