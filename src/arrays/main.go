package main

import (
	"fmt"
)

func main() {
	fmt.Println("arrays")

	var list [6]int

	list[0] = 8 //first index set to 8
	fmt.Println(list)
	fmt.Println(len(list))
	fmt.Println(FactorialArray(5))

	//slices -> array that allows u to work w collection of data
	//declare a slice

	var a []int //value: nil
	n := 4

	//slices must be made
	a = make([]int, n)

	//setting values of slices is similar to setting values of arrays
	a[0] = 1
	a[2] = 4
	fmt.Println(a)

	//one line decoration
	b := make([]int, 5)
	fmt.Println(b)
}

//slices and arrays difference
//arrays are modified in new memory location
//arrays are modified as a pointer, to the new memory

// Factorial Array Function
func FactorialArray(n int) []int {
	if n < 0 {
		panic("Error: please insert postive integer")
	}
	values := make([]int, n+1)

	for i := range values {
		values[i] = Factorial(i)
	}

	return values
}

func Factorial(n int) int {
	p := 1

	for i := 1; i <= n; i++ {
		p = p * i
	}
	return p
}

//Min Integer Array Function

func MinIntegerArray(list []int) int {
	if len(list) == 0 {
		panic("empty slice given")
	}

	min := list[0]

	for i := 0; i <= len(list)-1; i++ {
		if list[i] < min {
			min = list[i]
		}
	}
	return min
}

func MinIntegers(numbers ...int) int {
	if len(numbers) == 0 {
		panic("No numbers provided")
	}

	var minimun int

	for i, current := range numbers {
		if i == 0 || current < minimun {
			minimun = current
		}
	}
	return minimun
}
