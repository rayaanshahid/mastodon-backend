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
	// check origin : allows all incomming connections for now (for tests to run)
	// Should have a list of not-allowed connections to filter
	origin := r.Header.Get("Origin")
	return allowedOrigins[origin]
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	allowedOrigins = utils.AddAllowedOrigins(r.Header.Get("Origin"))
	upgrader.CheckOrigin = CheckOrigin
	websocketConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	} else {
		defer websocketConn.Close()

		// lock for the shared resource
		clients.mutex.Lock()
		clients.clientMap[websocketConn] = true
		clients.mutex.Unlock()
		log.Println("connection added! current number of connections: " + strconv.Itoa((len(clients.clientMap))))

		//Start to read messages on the connection
		ReceiveMessages(websocketConn)
	}
}
