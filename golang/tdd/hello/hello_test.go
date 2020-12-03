package hello

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying Hello to people", func(t *testing.T) {
		want := "Hello, imad"
		got := Hello("imad", "")
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, world' when an empty string is supplied", func(t *testing.T) {
		want := "Hello, world"
		got := Hello("", "")
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		want := "hola, Elodie"
		got := Hello("Elodie", "Spanish")
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		want := "Bonjour, imad"
		got := Hello("imad", "French")
		assertCorrectMessage(t, got, want)
	})
}