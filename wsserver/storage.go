package wsserver

import (
	"fmt"
	"log"
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

func (s *storage) SendToClient(message *Message, clientId string) error {
	var err error

	s.mutex.Lock()
	if client, ok := s.clients[clientId]; ok {
		client.Send <- message

	} else {
		err = fmt.Errorf("Client does not exist")
		log.Println(err)
	}
	s.mutex.Unlock()

	return err
}

func (s *storage) GetClientById(clientId string) *Client {
	s.mutex.Lock()
	client := s.clients[clientId]
	s.mutex.Unlock()

	return client
}
