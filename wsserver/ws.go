package wsserver

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

const (
	ReadBufferSize  = 1024
	WriteBufferSize = 512
)

type WS struct {
	clients  *storage
	upgrader websocket.Upgrader
}

func New() *WS {
	return &WS{
		clients: newStorage(),
		upgrader: websocket.Upgrader{
			ReadBufferSize:  ReadBufferSize,
			WriteBufferSize: WriteBufferSize,
		},
	}
}

func (w *WS) EchoHandler(c echo.Context) error {
	socket, err := w.upgrader.Upgrade(c.Response(), c.Request(), nil)

	if err != nil {
		log.Println(err)
		return err
	}

	id := c.Param("id")
	client := newClient(id, socket, w.clients)

	client.Run()
	w.clients.Register(id, client)

	return err
}

//Função que recebe um callback a ser executado quando se tenta enviar
//mensagens para um socket fechado
