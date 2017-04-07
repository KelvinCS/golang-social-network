package wsserver

import "sync"

type storage struct {
	clients map[string]*Client
	mutex sync.Mutex
}

func newStorage() *storage {
	return &storage{
		clients: make(map[string]*Client),
	}
}

func(s *storage) Register(client *Client, key string) {
	s.mutex.Lock()
	s.clients[key] = client
	s.mutex.Unlock()
}