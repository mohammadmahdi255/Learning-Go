package main

import "fmt"

func main() {
	var s2 string = "Mohammad Mahdi Nemati Haravani"
	s3 := "سلام دنیا"

	fmt.Println([]byte(s3))

	fmt.Printf("%d (len(s3))\n", len(s3))
	fmt.Println(s2)
	fmt.Println(s3[1])
	fmt.Printf("%c\n", s3[1])

	for _, c := range s3 {
		fmt.Printf("%c ", c)
	}

	fmt.Println()
}
