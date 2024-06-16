package main

import (
	"fmt"
)

// there is no while loops in golang, they use the word "for"
func main() {
	fmt.Println("this is while loop")
	fmt.Println(Factorial(5))
	// fmt.Println(Factorial(10))
	// fmt.Print(Factorial(15))
	fmt.Println(SumFirstNIntegers(6))
}

// input n and return n!
// n*(n-1)*(n-2)...*2*1
func Factorial(n int) int {
	if n < 0 {
		panic("insert positive number please")
	}

	p := 1 //store products
	i := 1 //counters

	for i <= n {
		p = p * i
		i += 1 //start from 1*2*3...until meet n then loop stops
	}
	return p
}

// n + (n-1) + (n-2) until n = 0

func SumFirstNIntegers(n int) int {
	if n < 0 {
		panic("please insert postive integer")
	}

	sum := 0 //store final answer
	i := 1

	for i <= n {
		sum = sum + i //sum+=i
		i += 1
	}

	return sum
}
