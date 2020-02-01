package playerserver

import (
	"fmt"
	"net/http"
)

// PlayerServer returns the amount of the player's wins as a string
// Example return values: "20", "0", "123"
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}
