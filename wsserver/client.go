package wsserver

import "github.com/gorilla/websocket"

type Client struct {
	Id     string
	Send   chan *Message
	socket *websocket.Conn
}
