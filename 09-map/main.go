package main

import "fmt"

func main() {

	var m map[int]string
	m = make(map[int]string)

	m[10] = "Hello"
	m[1380] = "Mahdi"

	for key, value := range m {
		fmt.Println(key, value)
	}

	var v string

	v = m[10]
	fmt.Printf("m[%d] = %s\n", 10, v)

	v = m[250]
	fmt.Printf("m[%d] = %s\n", 250, v)

	s := make(map[int]bool)

	for {
		var n int
		fmt.Scanf("%d\n", &n)

		if n == 0 {
			break
		}

		if s[n] {
			fmt.Printf("%d is already exist\n", n)
		} else {
			s[n] = true
			fmt.Printf("%d saved\n", n)
		}
	}

	opinions := make(map[string]bool)

	opinions["mahdi"] = true
	opinions["parham"] = true

	opinion, ok := opinions["amin"]
	if ok {
		fmt.Printf("opinions[%s] = %t", "amin", opinion)
	} else {
		fmt.Printf("%s exist in opinions", "amin")
	}

}
