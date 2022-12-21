package main

import "fmt"

func swap(a, b int) {
	a, b = b, a
}

func swapWithPointer(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	x, y := 10, 1
	swap(x, y)
	fmt.Printf("after \"swap(%d, %d)\": %d, %d\n", x, y, x, y)

	swapWithPointer(&x, &y)
	fmt.Printf("after \"swapWithPointer(%d, %d)\": %d, %d\n", y, x, x, y)
	fmt.Println(x, y)

	var a int
	a = 10
	fmt.Printf("%d\n", a)

	b := new(int)
	*b = 10
	fmt.Printf("b = %p, *b = %d\n", b, *b)

	var c *int
	fmt.Printf("c: %v %p\n", c, c)

}
