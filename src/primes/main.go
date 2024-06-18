package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Primes")
	fmt.Print(IsPrime(3))
	fmt.Print(TrivialPrimeFinder(4))
}

func TrivialPrimeFinder(n int) []bool {
	primeBooleans := make([]bool, n+1)

	for p := 2; p <= n; p++ {
		primeBooleans[p] = IsPrime(p)
	}
	return primeBooleans
}

func IsPrime(p int) bool {
	if p == 0 || p == 1 {
		return false
	}
	for k := 2; float64(k) <= math.Sqrt(float64(p)); k++ {
		if p%k == 0 {
			return false
		}
	}
	return true
}

func SieveOfEratosthenes(n int) []bool {
	primeBooleans := make([]bool, n+1)

	// everyone starts as false by default
	// now, set everything from 2 onward equal to true
	for p := 2; p <= n; p++ {
		primeBooleans[p] = true
	}

	for p := 2; float64(p) <= math.Sqrt(float64(n)); p++ {
		// is p prime? If so, cross off its multiples
		if primeBooleans[p] == true {
			primeBooleans = CrossOffMultiples(primeBooleans, p)
		}
	}

	return primeBooleans
}

func CrossOffMultiples(primeBooleans []bool, p int) []bool {
	n := len(primeBooleans) - 1

	// consider every multiple of p, starting with 2p, and "cross it off" by setting its corresponding entry of the slice to false
	for k := 2 * p; k <= n; k += p {
		primeBooleans[k] = false
	}

	return primeBooleans
}

func ListPrimes(n int) []int {
	// first, create a slice of length 0
	primeList := make([]int, 0)

	primeBooleans := SieveOfEratosthenes(n)
	for p := range primeBooleans {
		if primeBooleans[p] {
			// append p to our list
			primeList = append(primeList, p)
		}
	}

	return primeList
}
