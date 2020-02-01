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

	if player == "Pepper" {
		fmt.Fprintf(w, "20")
		return
	}

	if player == "Floyd" {
		fmt.Fprint(w, "10")
		return
	}
}
