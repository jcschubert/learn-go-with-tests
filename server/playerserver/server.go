package playerserver

import (
	"fmt"
	"net/http"
	"strings"
)

// PlayerServer returns the amount of the player's wins as a string
// Example return values: "20", "0", "123"
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Fprintf(w, GetPlayerScore(player))
}

// GetPlayerScore returns the score for the provided player name.
// If there is no such player, it returns an empty string.
func GetPlayerScore(name string) string {
	if name == "Pepper" {
		return "20"
	}

	if name == "Floyd" {
		return "10"
	}

	return ""
}
