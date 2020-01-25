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
	dict.Add("test", "this is just a test")

	want := "this is just a test"
	got, err := dict.Search("test")

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if want != got {
		t.Errorf("got %q want %q", got, want)
	}
}
