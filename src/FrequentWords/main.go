package main

import (
	"fmt"
)

func main() {
	fmt.Println("frequent words")
}

func FindFrequentWords(text string, k int) []int {
	freqPattern := make([]string, 0)

	freqMap := FrequencyTable(text, k) // key and value like dictionary in python

	//once having freq table, find the max value
	max := MaxMapValue(freqMap)

	//pattern stores the key, val stores the value
	for pattern, val := range freqMap {
		if val == max {
			freqPattern = append(freqPattern, pattern)
		}

	}
	return freqPattern()

}

func FrequencyTable(text string, k int) map[string]int { //key->map; value->int
	freqMap := make(map[string]int)
	n := len(text)

	//range over all substring of len(k)
	for i := 0; i < n-k+1; i++ {
		pattern := text[i : i+k] //grab current value

		//update the value of freqmap w this pattern w len(k)

	}
}

// find max value in the map
func MaxMapValue(dict map[string]int) int {
	m := 0
	FirstTimeThrough := true

	//range over map and update m when a larger value appear
	for _, val := range dict {
		if FirstTimeThrough || val > m {
			m = val
		}
		FirstTimeThrough = false
	}

	return m
}
