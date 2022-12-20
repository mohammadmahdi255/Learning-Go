package main

import "fmt"

const (
	c1       = 10.1
	c2 int64 = 11
)

//var v1 = 20
//var v2 uint64 = 12

var (
	v1        = 20
	v2 uint64 = 12
)

func main() {
	x := 10 // var x =10

	var y int

	var z float64 = 10.3

	fmt.Println(x)
	fmt.Printf("y = %d\n", y)
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Printf("v1 = %d v2 = %d\n", v1, v2)
	fmt.Printf("z = %f\n", z)
	fmt.Printf("z = %.10f\n", z)
	fmt.Printf("z = %g\n", z)
}
