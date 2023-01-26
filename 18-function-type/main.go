package main

import "fmt"

func logBeforeRun(f func(int) int) func(int) int {
	return func(i int) int {
		fmt.Println("Hello, I am log")
		r := f(i)
		fmt.Printf("Result: %d\n", r)
		return r
	}
}

func main() {
	f := func(x int) int {
		return x * x
	}

	fmt.Println(f(1))
	fmt.Println(f(2))

	functionMap := make(map[string]func(int) int)

	functionMap["sqrt"] = f

	logBeforeRun(f)(10)
}
