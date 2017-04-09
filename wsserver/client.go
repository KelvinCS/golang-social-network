package wsserver

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	Id         string
	Send       chan *Message
	Contacts   *storage
	allClients *storage
	online     bool
	socket     *websocket.Conn
}

func newClient(id string, socket *websocket.Conn, allClients *storage) *Client {
	return &Client{
		Id:         id,
		online:     true,
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
			fmt.Println(err.Error())
			break
		}
		fmt.Println(msg)

		c.allClients.SendToClient(&msg, msg.Destiny)
		fmt.Println("FIM!")
	}

	defer c.socket.Close()
}

//TODO: Não fechar mais o canal por aqui, o canal precisa ser fechado pelo escritor
func (c *Client) write() {
	defer c.socket.Close()

	for {
		select {
		case msg := <-c.allClients.pendingMessages[c.Id]:
			err := c.socket.WriteJSON(*msg)

			if err != nil {
				c.online = false
				c.allClients.SaveMessage(msg)
			}

		case msg := <-c.Send:
			err := c.socket.WriteJSON(*msg)

			if err != nil {
				c.online = false
				c.allClients.SaveMessage(msg)
			}
		}
	}

	//c.Send = c.allClients.pendingMessages[c.Id]

}
