package main

import (
	"fmt"
)

func main() {
	fmt.Println("hw0")
	fmt.Println(Permutation(25, 6))
	fmt.Println(Combination(5, 37))

}

func Permutation(n, k int) int {
	if n < k {
		return 0 // If k is greater than n, no permutation is possible.
	}
	result := 1
	for i := 0; i < k; i++ {
		result *= n - i
	}
	return result
}

func Combination(n, k int) int {
	if k > n {
		return 0 // If k is greater than n, no combination is possible.
	}
	if k > n-k { // Take advantage of symmetry, C(n, k) == C(n, n-k)
		k = n - k
	}
	result := 1
	for i := 1; i <= k; i++ {
		result = result * (n - k + i) / i //handle the multiplication and division without intermediate overflow
	}
	return result
}
