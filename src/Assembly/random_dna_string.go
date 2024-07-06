package main

import (
	"math/rand"
)

// GenerateRandomGenome takes a parameter length and returns
// a random DNA string of this length where each nucleotide has equal probability.
func GenerateRandomGenome(length int) string {
	genome := ""

	for i := 0; i < length+1; i++ {
		random := rand.Intn(4)

		if random == 0 {
			genome += "A"
		} else if random == 1 {
			genome += "T"
		} else if random == 2 {
			genome += "C"
		} else {
			genome += "G"
		}
	}
	return genome
}
