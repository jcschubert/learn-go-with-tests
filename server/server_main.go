package main

import (
	"log"
	"net/http"

	"github.com/jcschubert/learn-go-with-tests/server/playerserver"
)

func main() {
	server := &playerserver.PlayerServer{playerserver.NewInMemoryPlayerStore()}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("Could not listen on port 5000%v", err)
	}
}
