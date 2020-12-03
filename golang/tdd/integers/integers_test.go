package integers

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestAdd(t *testing.T) {
	want := 4
	got := Add(2, 2)

	if got != want {
		t.Errorf("got: '%d', want '%d'", got, want)
	}
}