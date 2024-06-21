package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("string")
	fmt.Println(strconv.Itoa(45))

	pi, err2 := strconv.ParseFloat("3.14", 64)
	if err2 != nil {
		panic(err2)
	}

	fmt.Println(pi)

	dna := "ACCGAT"
	fmt.Println(ReverseComplement(dna))
}

func ReverseComplement(pattern string) string {
	return Reverse(Complement(pattern))
}

func Complement(dna string) string {
	dna2 := make([]byte, len(dna))
	for i, symbol := range dna {
		switch symbol {
		case 'A':
			dna2[i] = 'T'
		case 'C':
			dna2[i] = 'G'
		case 'G':
			dna2[i] = 'C'
		case 'T':
			dna2[i] = 'A'
		default:
			panic("invalid input")
		}
	}
	return string(dna2)
}

func Reverse(pattern string) string {
	rev := make([]byte, len(pattern))
	n := len(pattern)

	for i := range pattern {
		rev[i] = pattern[n-1-i]
	}
	return string(rev)
}
