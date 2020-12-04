package iteration

func RepeatChar(char string, iter int) (repeated string) {
	for i := 0; i < iter; i++ {
		repeated = repeated + char
	}

	return
}