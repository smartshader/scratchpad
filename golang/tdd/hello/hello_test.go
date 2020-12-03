package hello

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		want := "hello, imad"
		got := hello("imad")
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'hello, world' when an empty string is supplied", func(t *testing.T) {
		want := "hello, world"
		got := hello("")
		assertCorrectMessage(t, got, want)
	})
}