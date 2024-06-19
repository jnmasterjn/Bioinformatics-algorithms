package main

import (
	"fmt"
)

func main() {
	fmt.Println("working w arrays")
	fmt.Println(FibonacciArray(4))
	fmt.Println(DividesAll([]int{2, 4, 6, 8}, 2))
	fmt.Println(MaxIntegerArray([]int{3, 4, 6, 2}))
	fmt.Println(SumInteger(5, 7, 4, 3, 77, 23, 6, 7))
	fmt.Println(SumIntegerr(5, 7, 4, 3, 77, 23, 6, 7))
	fmt.Println(GCDArray([]int{3, 6, 9}))
}

func FibonacciArray(n int) []int {
	if n <= 0 {
		return []int{1}
	}
	if n == 1 {
		return []int{1, 1}
	}

	fib := []int{1, 1}

	for i := 2; i <= n; i++ {
		next := fib[i-1] + fib[i-2] //calculates the sum
		fib = append(fib, next)
	}
	return fib
}

func DividesAll(a []int, d int) bool {
	if d == 0 {
		return false
	}
	if d == 1 {
		return true
	}
	for i := 0; i < len(a); i++ {
		if a[i]%d != 0 {
			return false // If any element is not divisible by d, return false.
		}
	}
	return true // Return true only if all elements are divisible by d
}

func MaxIntegerArray(list []int) int {
	if len(list) == 0 {
		panic("the list is empty")
	}
	max := list[0]
	for i := 1; i < len(list); i++ {
		if list[i] > max {
			max = list[i]
		}
	}
	return max
}

func MaxIntegers(numbers ...int) int {
	if len(numbers) == 0 {
		panic("the list is empty")
	}
	max := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > max {
			max = numbers[i]
		}
	}
	return max
}

func SumInteger(numbers ...int) int {
	sum := 0
	for i := 0; i <= len(numbers)-1; i++ {
		sum += numbers[i]
	}
	return sum
}

// using the range syntax stuff
func SumIntegerr(numbers ...int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}

func MinIntegerArray(list []int) int {
	if len(list) == 0 {
		panic("the list is empty")
	}
	min := list[0]
	for i := 1; i < len(list); i++ {
		if list[i] < min {
			min = list[i]
		}
	}
	return min
}

func GCDArray(a []int) int {
	if len(a) == 0 {
		return 0
	}
	min := MinIntegerArray(a)
	for divisor := min; divisor > 0; divisor-- {
		if DividesAll(a, divisor) {
			return divisor
		}
	}
	return 1
}
