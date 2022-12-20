package main

import "fmt"

func main() {

	s1 := make([]int, 10)

	s1[0] = 10
	s1[1] = 20
	s1[2] = 30

	fmt.Printf("s1, %+v, len(s1): %d, cap(s1): %d\n", s1, len(s1), cap(s1))

	s1 = append(s1, 10)
	fmt.Printf("s1, %+v, len(s1): %d, cap(s1): %d\n", s1, len(s1), cap(s1))

	for index, value := range s1 {
		fmt.Printf("s1[%d]: %d\n", index, value)
	}
	fmt.Println()

	s2 := make([]int, 0, 10)

	fmt.Printf("s2, %+v, len(s2): %d, cap(s2): %d\n", s2, len(s2), cap(s2))

	s1[0] = 123

	s2 = append(s2, 1)
	s2 = append(s2, 2)
	s2 = append(s2, 3)
	s2 = append(s2, 4)
	s2 = append(s2, 5)
	s2 = append(s2, 6)
	s2 = append(s2, 7)
	s2 = append(s2, 8)
	s2 = append(s2, 9)
	s2 = append(s2, 10)

	fmt.Printf("s2, %+v, len(s2): %d, cap(s2): %d\n", s2, len(s2), cap(s2))
}
