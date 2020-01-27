package racer

import (
	"fmt"
	"net/http"
	"time"
)

const tenSecondTimeout = 10 * time.Second

// Racer takes two urls a and b and returns the one with faster response time.
// If neiter url responds faster than 10 seconds, an error will be returned.
func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

// ConfigurableRacer takes two urls and a timeout and returns the url that responds faster
// After timeout, it will return an error
func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
	}
	return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
