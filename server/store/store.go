package store

import (
	"encoding/json"
	"io"
	"sync"
)

type InMemoryPlayerStore struct {
	mutex *sync.Mutex
	store map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		mutex: &sync.Mutex{},
		store: map[string]int{},
	}
}

func (s *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return s.store[name]
}

func (s *InMemoryPlayerStore) RecordWin(name string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.store[name]++
}

func (s *InMemoryPlayerStore) GetLeague() []Player {
	var league []Player
	for name, wins := range s.store {
		league = append(league, Player{name, wins})
	}
	return league
}

type FileSystemStore struct {
	Database io.Reader
}

func (f *FileSystemStore) GetLeague() []Player {
	var league []Player
	json.NewDecoder(f.Database).Decode(&league)
	return league
}
