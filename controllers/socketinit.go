package controllers

import (
	"net/http"

	"github.com/gorilla/websocket"
	"main.go/services/socket"
)

var Manager *socket.Manager

func SocketInit() {

	Manager = &socket.Manager{
		Client: make(map[string]*websocket.Conn),
	}

}

func HandleWS(w http.ResponseWriter, r *http.Request) {
	Manager.HandleConn(w, r)
}
