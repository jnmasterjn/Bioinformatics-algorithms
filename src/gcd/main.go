package main

import (
	"fmt"
)

// TrivialGCD(a,b)
//     d ← 1
//     m ← Min2(a,b)
//     for every integer p between 1 and m
//         if p is a divisor of both a and b
//             d ← p
//     return d

func main() {
	fmt.Println("hi")
}

func Min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TrivialGCD(a, b int) int {
	d := 1
	m := Min2(a, b)
	for i 

}
