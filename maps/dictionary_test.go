package maps

import "testing"

func assertString(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("expected to get an error.")
	}
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertDefinition(t *testing.T, dict Dictionary, word, definition string) {
	t.Helper()

	got, err := dict.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
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
		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	dict := Dictionary{}
	word := "test"
	definition := "this is just a test"
	dict.Add(word, definition)

	assertDefinition(t, dict, word, definition)
}
