package racer

import (
	"net/http"
	"time"
)

// Racer takes two urls a and b and returns the one with faster response time.
func Racer(a, b string) string {
	if responseTime(a) < responseTime(b) {
		return a
	}

	return b
}

func responseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
