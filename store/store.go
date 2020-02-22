package store

import (
	"sync"
)

type Store struct {
	// using channels instead of mutex
	urisCh      chan string
	runOnce     *sync.Once
	RequestURIs map[string]uint
}

func NewStore() *Store {
	return &Store{
		urisCh:      make(chan string),
		runOnce:     &sync.Once{},
		RequestURIs: make(map[string]uint),
	}
}

func (s *Store) Save(uri string) {
	if uri != "" {
		s.urisCh <- uri
	}
}

func (s *Store) runUriConsumer() {
	for {
		select {
		case uri := <-s.urisCh:
			v, exists := s.RequestURIs[uri]
			if !exists {
				s.RequestURIs[uri] = 1
			} else {
				s.RequestURIs[uri] = v + 1
			}
		}
	}
}

func (s *Store) Run() *Store {
	s.runOnce.Do(func() {
		// start Consumers
		go s.runUriConsumer()
	})
	return s
}
