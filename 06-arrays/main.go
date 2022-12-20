package main

import "fmt"

func main() {

	var a [10]int
	var b [4]int = [...]int{1, 2, 3, 4}
	var c = [4]int{1, 4, 5, 8}
	d := [4]int{1}

	for index, value := range a {
		fmt.Printf("a[%d] = %d\n", index, value)
	}

	fmt.Println()

	for index, value := range b {
		fmt.Printf("b[%d] = %d\n", index, value)
	}

	fmt.Println()

	// a = c
	// compile error, cannot use c (type [4]int) as type [10]int in assignment

	fmt.Println("c: ", c)
	fmt.Println("d: ", d)

	b = c

	fmt.Println("b: ", b)

	b[0] = 10

	fmt.Println("after change b", c)

	//  n := 10
	// var c [n]int
	//  compile error, arrays' length must be constant
}
