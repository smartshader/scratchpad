package arrays

import "testing"

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	want := 15
	got := Sum(numbers)

	if got != want {
		t.Errorf("want: %d, got: %d, given %v", want, got, numbers)
	}
}
