package hello

import "testing"

func TestHello(t *testing.T) {
	want := "hello, imad"
	got := hello("imad")

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}