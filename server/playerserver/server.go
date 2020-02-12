package playerserver

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	Store PlayerStore
}

// PlayerServer returns the amount of the player's wins as a string
// Example return values: "20", "0", "123"
// returns status 404 if player not found
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		p.processWin(w)
	case http.MethodGet:
		p.postScore(w, r)
	}
}

func (p *PlayerServer) postScore(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	score := p.Store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter) {
	p.Store.RecordWin("Bob")
	w.WriteHeader(http.StatusAccepted)
}
