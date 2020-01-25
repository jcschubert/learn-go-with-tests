package maps

import "testing"

func assertString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "this is just a test"}
	got := dict.Search("test")
	want := "this is just a test"
	assertString(t, got, want)
}
