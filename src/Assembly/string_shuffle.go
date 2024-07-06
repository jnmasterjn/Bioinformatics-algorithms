package main

import (
	"math/rand"
)

//ShuffleStrings takes a collection of strings patterns as input.
//It returns a random shuffle of the strings.

func ShuffleStrings(patterns []string) []string {

	random_generator := rand.Perm(len(patterns)) //store index
	shuffle := make([]string, len(patterns))     //store value

	for i := 0; i < len(random_generator); i++ {
		for u := 0; u < len(patterns); u++ { //nested loop to find all posible positions
			if random_generator[i] == u {
				shuffle[i] = patterns[u]
			}
		}
	}

	return shuffle
}
