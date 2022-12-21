package main

import "fmt"

func genericSum[v int16 | int32 | int64 | float32 | float64](n, m v) v {
	return m + n
}

func main() {
	var x float64 = 10
	y := 3.4
	fmt.Println(genericSum(x, y))
}
