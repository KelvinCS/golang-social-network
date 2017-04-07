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

func (w *ws) EchoHandler(c echo.Context) {
	socket, err := w.upgrader.Upgrade(c.Response(), c.Request(), nil)

	if err != nil {
		log.Println(err)
	}

	client := newClient("Kelvin", socket)

	w.storage.Register("Kelvin", client)
}
