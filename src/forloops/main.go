package main

import (
	"fmt"
)

// for initialization; condition; postcondition

func main() {
	fmt.Println("for loops")
	fmt.Println(Factorial(4))
	fmt.Println(SumEven(3))

}

func Factorial(n int) int {

	p := 1

	for i := 1; i <= n; i++ {
		p = p * i
	}
	return p
}

func SumEven(k int) int {

	var sum int

	for i := 0; i <= k; i++ {
		if i%2 == 0 {
			sum += i
		}
	}

	return sum

}
