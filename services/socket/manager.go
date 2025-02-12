package socket

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type Manager struct {
	Client map[string]*websocket.Conn
}

var upgrader = websocket.Upgrader{
	CheckOrigin:     checkOrigin,
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func checkOrigin(r *http.Request) bool {
	return true
}

func (m *Manager) HandleConn(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	fmt.Println(key)
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Upgrade error:", err)
		return
	}
	// defer conn.Close()

	m.Client[key] = conn
	fmt.Println("Client connected")
	conn.WriteMessage(websocket.TextMessage, []byte("Connected to server"))
}
