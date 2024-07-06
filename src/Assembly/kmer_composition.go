package main

//KmerComposition returns the k-mer composition (all k-mer substrings) of a given genome.

/*Code Challenge: Implement a function KmerComposition(text, k).

Input: a string text and an integer k.
Return: an array containing all k-mer substrings of text, including repeated strings.
Note: The strings may appear in your array in any order.

Sample Input 1:

1
ATCG
Sample Output 1:

A C G T
Sample Input 2:

4
ATCG
Sample Output 2:

ATCG
Sample Input 3:

2
CCCCCCC
Sample Output 3:

CC CC CC CC CC CC*/

func KmerComposition(genome string, k int) []string {
	store := make([]string, 0, len(genome)) //make slice
	for i := 0; i < len(genome)-k+1; i++ {  //n-k+1
		store = append(store[i : i+k]) //start from i, end before i+k, 剛好會是k-mer
	}
	return store
}
