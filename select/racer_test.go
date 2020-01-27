package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compare speed of servers, returning the url of the fastest", func(t *testing.T) {
		slowHTTP := makeDelayedServer(20 * time.Millisecond)
		defer slowHTTP.Close()

		fastHTTP := makeDelayedServer(0 * time.Millisecond)
		defer fastHTTP.Close()

		got, _ := Racer(slowHTTP.URL, fastHTTP.URL)

		if got != fastHTTP.URL {
			t.Errorf("got %q, want %q", got, fastHTTP.URL)
		}
	})
	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		serverA := makeDelayedServer(11 * time.Second)
		defer serverA.Close()

		serverB := makeDelayedServer(12 * time.Second)
		defer serverB.Close()

		_, err := Racer(serverA.URL, serverB.URL)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}

	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
