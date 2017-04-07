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

type ws struct {
	storage  *storage
	upgrader websocket.Upgrader
}

func New() *ws {
	return &ws{
		storage: newStorage(),
		upgrader: websocket.Upgrader{
			ReadBufferSize:  ReadBufferSize,
			WriteBufferSize: WriteBufferSize,
		},
	}
}

func (w *ws) EchoHandler(c echo.Context) error {
	socket, err := w.upgrader.Upgrade(c.Response(), c.Request(), nil)

	if err != nil {
		log.Println(err)
	}

	id := c.FormValue("id")
	client := newClient(id, socket)

	w.storage.Register(id, client)

	return err
}
