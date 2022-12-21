package main

import (
	"fmt"
)

func swap(a, b int) (int, int) {
	return b, a
}

func complexFunction(a int) (int, string) {
	c := a + 10

	return c, fmt.Sprintf("%X", a)
}

func main() {
	a, b := swap(10, 2)

	fmt.Println(a)
	fmt.Println(b)

	_, c := swap(20, 1)
	fmt.Println(c)

	// compiler error: assignment mismatch: 1 variable but swap returns 2 values
	// d := swap(10, 20)

	fmt.Println(complexFunction(10))
}
