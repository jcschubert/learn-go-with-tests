package maps

import "testing"

func assertString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func TestSearch(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		dict := Dictionary{"test": "this is just a test"}
		got, _ := dict.Search("test")
		want := "this is just a test"
		assertString(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		dict := Dictionary{"test": "this is just a test"}
		_, err := dict.Search("not there")
		want := "could not find the word you were looking for"
		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertString(t, err.Error(), want)
	})
}
