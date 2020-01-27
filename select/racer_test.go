package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowHTTP := makeDelayedServer(20 * time.Millisecond)
	defer slowHTTP.Close()

	fastHTTP := makeDelayedServer(0 * time.Millisecond)
	defer fastHTTP.Close()

	got := Racer(slowHTTP.URL, fastHTTP.URL)

	if got != fastHTTP.URL {
		t.Errorf("got %q, want %q", got, fastHTTP.URL)
	}
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
