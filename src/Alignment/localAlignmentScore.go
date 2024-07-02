package main

//LocalScoreTable takes two strings and alignment penalties. It returns a 2-D array
//holding dynamic programming scores for local alignment with these penalties.

/*
Code Challenge: Implement a function LocalScoreTable(str1, str2, match, mismatch, gap).

Input: two strings, str1 and str2 , and three decimals, match, mismatch, and gap.
Return: the score table for local alignment of str1 and str2, using match as a match reward, mismatch as a mismatch penalty, and gap as a gap penalty.
Test case documentation (you must be logged into CMU account to view)

Sample Input 1:

1 1 2
GAGA
GAT
Sample Output 1:

0 0 0 0
0 1 0 0
0 0 2 0
0 1 0 1
0 0 2 0
Sample Input 2:

3 3 1
AGC
ATC
Sample Output 2:

0 0 0 0
0 3 2 1
0 2 1 0
0 1 0 4
Sample Input 3:

1 1 1
AT
AG
Sample Output 3:

0 0 0
0 1 0
0 0 0
Sample Input 4:

1 1 1
TAACG
ACGTG
Sample Output 4:

0 0 0 0 0 0
0 0 0 0 1 0
0 1 0 0 0 0
0 1 0 0 0 0
0 0 2 1 0 0
0 0 1 3 2 1
Sample Input 5:

3 2 1
CAGAGATGGCCG
ACG
Sample Output 5:

0 0 0 0
0 0 3 2
0 3 2 1
0 2 1 5
0 3 2 4
0 2 1 5
0 3 2 4
0 2 1 3
0 1 0 4
0 0 0 3
0 0 3 2
0 0 3 2
0 0 2 6
*/
func LocalScoreTable(str1, str2 string, match, mismatch, gap float64) [][]float64 {
	//Make the table
	table := Matrix(len(str1)+1, len(str2)+1)

	//Set the values of the matrix

	for row := 1; row <= len(str1); row++ {
		for col := 1; col <= len(str2); col++ {
			up := table[row-1][col] - gap
			left := table[row][col-1] - gap
			diag := table[row-1][col-1]

			//Two cases: match or mismatch
			if str1[row-1] == str2[col-1] { //match
				diag += match
			} else {
				diag -= mismatch
			}

			//Maximum score of three directions
			max4 := Max(diag, left, up, 0.0)
			table[row][col] = max4

		}
	}
	return table
}
