package websockets

import (
	"log"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/mattn/go-mastodon"
)

var Broadcast = make(chan mastodon.Status)

func ReceiveMessages(websocketConn *websocket.Conn) {
	for {
		_, _, err := websocketConn.ReadMessage()
		if err != nil {
			log.Printf("error: %v", err)
			clients.mutex.Lock()
			delete(clients.clientMap, websocketConn)
			clients.mutex.Unlock()
			log.Println("connection deleted! current number of connections: " + strconv.Itoa((len(clients.m))))
			break
		}
	}
}

func PublishMessage() {
	for {
		msg := <-Broadcast
		for client := range clients.clientMap {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients.clientMap, client)
			}
		}
	}
}
