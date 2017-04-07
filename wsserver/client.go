package wsserver

import "github.com/gorilla/websocket"

type Client struct {
	Id       string
	Send     chan *Message
	Contacts *storage
	socket   *websocket.Conn
}

func newClient(id string, socket *websocket.Conn) *Client {
	return &Client{
		Id:       id,
		Send:     make(chan *Message),
		Contacts: newStorage(),
		socket:   socket,
	}
}
