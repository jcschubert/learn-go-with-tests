package main

import (
	"log"
	"net/http"

	"github.com/jcschubert/learn-go-with-tests/server/server"
	"github.com/jcschubert/learn-go-with-tests/server/store"
)

func main() {
	server := server.NewPlayerServer(store.NewInMemoryPlayerStore())
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("Could not listen on port 5000: %v", err)
	}
}
