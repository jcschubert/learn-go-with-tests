package store

import (
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

type FileSystemPlayerStore struct {
	Database io.ReadSeeker
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	f.Database.Seek(0, 0)
	league, _ := NewLeague(f.Database)
	return league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	var wins int

	for _, player := range f.GetLeague() {
		if player.Name == name {
			wins = player.Wins
			break
		}
	}

	return wins
}
