package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	want := 15
	got := Sum(numbers)

	if got != want {
		t.Errorf("want: %d, got: %d, given %v", want, got, numbers)
	}
}

func TestSumAll(t *testing.T) {
	numbers1 := []int{1, 2, 3, 4, 5}
	numbers2 := []int{50, 100}

	want := []int{15, 150}
	got := SumAll(numbers1, numbers2)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want: %d, got: %d, given %v and %v", want, got, numbers1, numbers2)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int, numbersToSum ...[]int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want: %d, got: %d, given %v", want, got, numbersToSum)
		}
	}

	t.Run("sum of tails of some slices", func(t *testing.T) {
		numbers1 := []int{1, 2, 3, 4, 5}
		numbers2 := []int{50, 100}

		want := []int{14, 100}
		got := SumAllTails(numbers1, numbers2)

		checkSums(t, got, want, numbers1, numbers2)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		var numbers1 []int
		numbers2 := []int{50, 100}

		want := []int{0, 100}
		got := SumAllTails(numbers1, numbers2)

		checkSums(t, got, want, numbers1, numbers2)
	})
}
