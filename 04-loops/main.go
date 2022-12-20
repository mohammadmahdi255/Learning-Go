package main

import "fmt"

func main() {
	var n int

	fmt.Print("Please enter n: ")
	fmt.Scanf("%d", &n)

	fmt.Printf("is %d prime? %t\n", n, IsPrime(n))

	fmt.Printf("%dth prime is %d\n", n, NthPrime(n))

	fmt.Println("gcd", 10, 20, "is", GCD(10, 20))
	fmt.Println("gcd", 11, 15, "is", GCD(11, 15))

}

func GCD(n, m int) int {
	for n%m != 0 {
		n, m = m, n%m
	}

	return m
}

func NthPrime(n int) int {
	i := 2
	counter := 0

	for {
		if IsPrime(i) {
			counter++

			if counter == n {
				return i
			}
		}

		i++
	}
}

func IsPrime(n int) bool {

	if n == 1 {
		return false
	}

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
