package main

import (
	"fmt"
)

func main() {
	fmt.Println(Min2(4, 6), "is smaller")
	fmt.Println(SameSign(-4, 2))
	fmt.Println(WhichIsGreater(55555, -222222))
	fmt.Println(PositiveDifference(55, -22))

}

func Min2(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func WhichIsGreater(x, y int) int {
	if x == y {
		return 0
	} else if x > y {
		return 1
	} else {
		return -1
	}
}

// return true when two integer have same signs
// 0 is both negative and positve
// bitwise and &
// logic and &&
func SameSign(x, y int) bool {
	return (x >= 0 && y >= 0) || (x <= 0 && y <= 0)
}

// take in two integer and return their absolute difference
func PositiveDifference(a, b int) int {
	Abs := func(c int) int {
		if c < 0 {
			return -c
		} else {
			return c
		}
	}
	a = Abs(a)
	b = Abs(b)

	if a > b {
		return a - b
	} else {
		return b - a
	}
}
