package wsserver

import (
	"fmt"

	"sync"
)

type storage struct {
	pendingMessages map[string]chan *Message
	clients         map[string]*Client
	mutex           sync.Mutex
}

func newStorage() *storage {
	return &storage{
		pendingMessages: make(map[string]chan *Message),
		clients:         make(map[string]*Client),
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
		if client.online {
			select {

			case client.Send <- message:

			default:
				s.SaveMessage(message)
				delete(s.clients, client.Id)
			}

		} else {
			close(client.Send)
			delete(s.clients, client.Id)
		}

	} else {
		s.SaveMessage(message)

	}

	s.mutex.Unlock()

	return err
}

func (s *storage) SaveMessage(message *Message) {
	if _, ok := s.pendingMessages[message.Destiny]; !ok {

		s.pendingMessages[message.Destiny] = make(chan *Message)

	}
	go func() { s.pendingMessages[message.Destiny] <- message }()

}

func (s *storage) GetClientById(clientId string) *Client {
	s.mutex.Lock()
	client := s.clients[clientId]
	s.mutex.Unlock()

	return client
}
