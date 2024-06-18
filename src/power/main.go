package main

import (
	"fmt"
)

func main() {
	fmt.Println("Power")
}

// 'b' is the times that 'a' gonna multiply itself
func Power(a, b int) int {
	if b == 0 {
		return 1
	}

	if b < 0 { //if b is negative
		a = 1 / a
		b = -b
	}

	answer := 1

	for i := 0; i < b; i++ {
		answer *= a
	}
	return answer
}

//10 -> 1, 2, 5, 10

func SumProperDivisors(n int) int {
	if n == 0 || n < 0 {
		panic("please insert an integer that is positive and not zero")
	}

	sum := 0
	for i := 1; i <= n-1; i++ {
		if n%i == 0 {
			sum = sum + i
		}
	}
	return sum
}
