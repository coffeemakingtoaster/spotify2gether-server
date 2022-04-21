package socket

import (
	"errors"
	"net/http"

	"github.com/gorilla/websocket"
	"golang.org/x/text/message"
	"spotify2gether.server/pkg/helpers"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

//Upgrade http to websocket connection
func upgradeConnectionToWS(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, errors.New("Something went wrong")
	}
	return ws, nil

}

func HandleInitialConnection(w http.ResponseWriter, r *http.Request) *websocket.Conn {
	var socketConn, err = upgradeConnectionToWS(w, r)
	//if connection was not http => give error
	if err != nil {
		helpers.ConnectionNotWebsocketError(w, r)
		return nil
	}
	return socketConn

}

func (c *websocket.Conn) WsSender(message byte[] ) error {
	c.WriteMessage(0, message)
	return nil
}
