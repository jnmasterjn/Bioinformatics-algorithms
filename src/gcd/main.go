package main

import (
	"fmt"
	"log"
	"time"
)

// TrivialGCD(a,b)
//     d ← 1
//     m ← Min2(a,b)
//     for every integer p between 1 and m
//         if p is a divisor of both a and b
//             d ← p
//     return d

func main() {
	fmt.Println("find GCD")
	fmt.Println(TrivialGCD(63, 42))
	fmt.Println(EuclidGCD(42, 63))

	x := 192847289
	y := 24875428

	start := time.Now() //start stop watch
	TrivialGCD(x, y)
	elasped := time.Since(start)              //stop the watch
	log.Printf("TrivialGCD took %s", elasped) //print the time
}

func Min2(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func TrivialGCD(a, b int) int {
	d := 1
	m := Min2(a, b)
	for p := 1; p <= m; p++ {
		// p is a divisor of a, if a % p == 0
		if a%p == 0 {
			//now check if p is divisor of b
			if b%p == 0 {
				//if the loop came here, we know p is divisor of both a, b
				d = p
			}
		}
	}
	return d
}

// EuclidGCD(a,b)
//     while a ≠ b
//         if a > b
//             a ← a − b
//         else
//             b ← b − a
//     return a

func EuclidGCD(a, b int) int {
	for a != b {
		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}
	return a
}
