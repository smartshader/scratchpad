package iteration

import "testing"

func BenchmarkRepeatChar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatChar("a", 10)
	}
}

func TestRepeatChar(t *testing.T) {
	t.Run("repeating character", func(t *testing.T) {
		want := "cccccccccc"
		got := RepeatChar("c", 10)

		if got != want {
			t.Errorf("want: %q, but got: %q", want, got)
		}
	})
}