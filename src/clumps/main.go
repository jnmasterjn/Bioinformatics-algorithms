package main

import (
	"fmt"
	"io"
	"log"
	"net/http" //accesssing urls
	//need to read and write from files
	//for log files
)

func main() {
	fmt.Println("clumps")

	url := "https://bioinformaticsalgorithms.com/data/realdatasets/Replication/E_coli.txt"

	resp, err := http.Get(url)

	//panic if there's an error
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Recieved Non-OK status: %v", resp.Status)
	}

	GenomeSymbol, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(len(GenomeSymbol))

	//close connection
	resp.Body.Close()

}

// k: The length of the k-mers (substrings) to search for.
// L: The length of the window (a specific segment of the string) within which to find these repeated k-mers.
// t: The minimum number of times a k-mer must appear within a window of length L to be considered a clump.

func FindClumps(text string, k, L, t int) []string {
	n := len(text)
	foundPatterns := make([]string, 0)
	patternSet := make(map[string]bool) // Using a map to track unique patterns efficiently

	// Ranging over possible windows
	for i := 0; i <= n-L; i++ {
		window := text[i : i+L]
		freqMap := FrequencyTable(window, k)

		// Check frequency within the window
		for pattern, count := range freqMap {
			if count >= t {
				if _, exists := patternSet[pattern]; !exists {
					foundPatterns = append(foundPatterns, pattern)
					patternSet[pattern] = true // Mark this pattern as found
				}
			}
		}
	}
	return foundPatterns
}

func FrequencyTable(text string, k int) map[string]int {
	freqMap := make(map[string]int)
	n := len(text)

	// Range over all substrings of length k
	for i := 0; i <= n-k; i++ {
		pattern := text[i : i+k]
		freqMap[pattern]++
	}
	return freqMap
}
