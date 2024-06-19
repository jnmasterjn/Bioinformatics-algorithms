package main

import (
	"fmt"
)

func main() {
	fmt.Println("perfect")
	fmt.Println(SumProperDivisors(6))
}

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

func IsPerfect(n int) bool {
	if n == SumProperDivisors(n) {
		return true
	}
	return false
}
