package wsserver

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

const (

)

type ws struct {
	storage  *storage
	upgrader websocket.Upgrader
}

func New() *ws {
	return &ws{
		storage: newStorage(),
	}
}

func (w *ws) EchoHandler(c echo.Context) {

}
