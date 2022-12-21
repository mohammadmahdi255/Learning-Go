package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s = "1230"
	var a = 10

	i, _ := strconv.ParseInt(s, 10, 64)

	a = int(i) + a

	fmt.Println(i)
	fmt.Println(a)

	fmt.Println(strconv.Atoi("123"))
}
