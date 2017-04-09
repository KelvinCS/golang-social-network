package wsserver

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

const (
	ReadBufferSize  = 1024
	WriteBufferSize = 512
)

type WS struct {
	storage  *storage
	upgrader websocket.Upgrader
}

func New() *WS {
	return &WS{
		storage: newStorage(),
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

	fmt.Println("Cliente conectado")
	fmt.Println(c.Request().Method)

	id := c.Param("id")
	fmt.Println(id)
	client := newClient(id, socket, w.storage)

	w.storage.Register(id, client)

	return err
}
