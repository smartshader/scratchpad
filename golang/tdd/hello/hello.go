package hello

const helloPrefix = "hello, "

func hello(name string) string {
	if name == "" {
		name = "world"
	}

	return helloPrefix + name
}
