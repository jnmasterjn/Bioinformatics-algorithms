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
	fmt.Println("the skew array ")

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

func SkewArray(genome string) []int {
	n := len(genome)

	array := make([]int, n+1)
	array[0] = 0

	for i := 1; i < n+1; i++ {
		array[i] = array[i-1] + Skew(genome[i-1])

	}

	return array
}

func Skew(symbol byte) int {
	if symbol == 'G' {
		return 1
	} else if symbol == 'C' {
		return -1
	}
	return 0
}
