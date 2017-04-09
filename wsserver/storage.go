package wsserver

import (
	"fmt"
	"sync"
)

type storage struct {
	clients map[string]*Client
	mutex   sync.Mutex
}

func newStorage() *storage {
	return &storage{
		clients: make(map[string]*Client),
	}
}

func (s *storage) Register(key string, client *Client) {
	s.mutex.Lock()
	s.clients[key] = client
	s.mutex.Unlock()

	fmt.Println(s.clients)
}

func (s *storage) SendToClient(message *Message, clientId string) {
	s.mutex.Lock()
	s.clients[clientId].Send <- message
	s.mutex.Unlock()
}

func (s *storage) GetClientById(clientId string) *Client {
	s.mutex.Lock()
	client := s.clients[clientId]
	s.mutex.Unlock()

	return client
}
