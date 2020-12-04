package iteration

import "testing"

func BenchmarkRepeatChar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatChar("a")
	}
}

func TestRepeatChar(t *testing.T) {
	want := "ccccc"
	got := RepeatChar("c")

	if got != want {
		t.Errorf("want: %q, but got: %q", want, got)
	}
}