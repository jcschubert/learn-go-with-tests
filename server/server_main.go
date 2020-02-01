package main

import (
	"log"
	"net/http"

	"github.com/jcschubert/learn-go-with-tests/server/playerserver"
)

func main() {
	handler := http.HandlerFunc(playerserver.PlayerServer)

	if err := http.ListenAndServe(":5000", handler); err != nil {
		log.Fatalf("Could not listen on port 5000%v", err)
	}
}
