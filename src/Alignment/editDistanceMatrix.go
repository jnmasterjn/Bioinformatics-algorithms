package main

//EditDistanceMatrix takes as input a slice of strings patterns.
//It returns a matrix whose (i,j)th entry is the edit distance between
//the i-th and j-th strings in patterns.
/*Code Challenge: Implement a function EditDistanceMatrix(patterns).

Input: an array of strings, patterns.
Return: a matrix E such that E[i][j] is the edit distance between the i-th string and the j-th strings in patterns.
Functions that are implemented behind the scenes for you:

EditDistance(): takes as input two strings, and returns the edit distance between the two strings.

Sample Output 4:

0 2 2 2
2 0 2 2
2 2 0 2
2 2 2 0
Sample Input 5:

AC CT TG GC AG GA GAT CAT ACG GAC
Sample Output 5:

0 2 2 1 1 2 2 2 1 1
2 0 2 2 2 2 2 1 2 3
2 2 0 2 1 2 3 3 2 3
1 2 2 0 2 1 2 3 2 1
1 2 1 2 0 2 2 2 1 2
2 2 2 1 2 0 1 2 3 1
2 2 3 2 2 1 0 1 3 1
2 1 3 3 2 2 1 0 3 2
1 2 2 2 1 3 3 3 0 2
1 3 3 1 2 1 1 2 2 0 */

func EditDistanceMatrix(patterns []string) [][]int {

	//make a matrix storing the values of editDistances
	stor := Matrix(len(array), len(array))
	//nested for loops calculating distance between different array[] values
	for i := 0; i < len(array); i++ {

		for j := 0; j < len(array); j++ {
			stor[i][j] = EditDistance(array[i], array[j])

		}

	}
	return stor
}
