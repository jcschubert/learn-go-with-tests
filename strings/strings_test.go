package strings

import (
	"strings"
	"testing"
)

func TestStringCompare(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("lexicographical, first = second", func(t *testing.T) {
		got := strings.Compare("a", "a")
		want := 0
		assertCorrectMessage(t, got, want)
	})

	t.Run("lexicographical comparison, first < second", func(t *testing.T) {
		got := strings.Compare("a", "z")
		want := -1
		assertCorrectMessage(t, got, want)
	})

	t.Run("lexicographical comparison, second > first", func(t *testing.T) {
		got := strings.Compare("z", "a")
		want := 1
		assertCorrectMessage(t, got, want)
	})

}
