package main

import (
	"log"
	"net/http"

	"github.com/jcschubert/learn-go-with-tests/server/playerserver"
)

type InMemoryPlayerStore struct{}

func (s *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	server := &playerserver.PlayerServer{&InMemoryPlayerStore{}}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("Could not listen on port 5000%v", err)
	}
}
