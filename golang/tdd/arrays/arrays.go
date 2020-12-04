package arrays

func Sum(numbers []int) (sum int) {
	for _, value := range numbers {
		sum += value
	}

	return
}

func SumAll(numbersToSum ...[]int) (sum []int) {
	for _, a := range numbersToSum {
		sum = append(sum, Sum(a))
	}

	return
}

func SumAllTails(numbersToSum ...[]int) (sum []int) {
	for _, a := range numbersToSum {
		if len(a) == 0 {
			sum = append(sum, 0)
		} else {
			sum = append(sum, Sum(a[1:]))
		}
	}

	return
}
