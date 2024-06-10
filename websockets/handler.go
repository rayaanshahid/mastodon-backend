package websockets

import (
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/mastodon-backend/utils"
)

var clients = struct {
	mutex     sync.Mutex
	clientMap map[*websocket.Conn]bool
}{clientMap: make(map[*websocket.Conn]bool)}

var upgrader = websocket.Upgrader{}
var allowedOrigins = make(map[string]bool)

func CheckOrigin(r *http.Request) bool {
	origin := r.Header.Get("Origin")
	return allowedOrigins[origin]
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	allowedOrigins = utils.AddAllowedOrigins()
	upgrader.CheckOrigin = CheckOrigin
	websocketConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	} else {
		defer websocketConn.Close()
		clients.mutex.Lock()
		clients.clientMap[websocketConn] = true
		clients.mutex.Unlock()
		log.Println("connection added! current number of connections: " + strconv.Itoa((len(clients.clientMap))))
		ReceiveMessages(websocketConn)
	}
}
