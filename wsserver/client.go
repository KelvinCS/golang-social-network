package wsserver

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	Id         string
	Send       chan *Message
	Contacts   *storage
	allClients *storage
	socket     *websocket.Conn
}

func newClient(id string, socket *websocket.Conn, allClients *storage) *Client {
	return &Client{
		Id:         id,
		Send:       make(chan *Message),
		Contacts:   newStorage(),
		allClients: allClients,
		socket:     socket,
	}
}

func (c *Client) Run() {
	go c.read()
	go c.write()
}

func (c *Client) read() {
	for {
		var msg Message
		err := c.socket.ReadJSON(&msg)

		if err != nil {
			break
		}

		c.allClients.SendToClient(&msg, msg.Destiny)
	}

	defer c.socket.Close()
}

func (c *Client) write() {
	for msg := range c.Send {
		err := c.socket.WriteJSON(*msg)

		if err != nil {
			break
		}
	}

	defer c.socket.Close()
}
