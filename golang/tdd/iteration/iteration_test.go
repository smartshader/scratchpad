package iteration

import (
	"fmt"
	"testing"
)

func BenchmarkRepeatChar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatChar("a", 10)
	}
}

func ExampleRepeatChar() {
	repeated := RepeatChar("f", 3)
	fmt.Println(repeated)
	// Output: fff
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