package main

//LocalAlignment takes two strings, along with match, mismatch, and gap scores.
//It returns a maximum score local alignment of the strings corresponding to these penalties.
/*Code Challenge: Implement a function LocalAlignment(str1, str2, match, mismatch, gap ).

Input: two strings, str1 and str2 , and three integers, match, the match reward, mismatch, the mismatch penalty, and gap , the indel penalty.
Return: a maximum-scoring local alignment of str1 and str2 (having type Alignment) using match as a match reward, mismatch as a mismatch penalty, and gap as a gap penalty, followed by four integers: start1, end1, start2, and end2; the first two integers correspond to the start and end of the alignment in str1, and the second two integers correspond to the start and end of the alignment in str2.
Functions that are implemented behind the scenes for you:

LocalScoreTable(): takes as input two strings and three scoring parameters, and returns a matrix of values corresponding to the local alignment dynamic programming table for the two strings.
Test case documentation (you must be logged into CMU account to view)

Sample Input 1:

1 1 2
GAGA
GAT
Sample Output 1:

A sample correct alignment, aligning indices 0 to 2 in the first string and 0 to 2 in the second string, for this problem is:
GA
GA
Your alignment should have the same score as the sample alignment: 2
Sample Input 2:

3 3 1
AGC
ATC
Sample Output 2:

A sample correct alignment, aligning indices 0 to 3 in the first string and 0 to 3 in the second string, for this problem is:
A-GC
AT-C
Your alignment should have the same score as the sample alignment: 4
Sample Input 3:

1 1 1
AT
AG
Sample Output 3:

A sample correct alignment, aligning indices 0 to 1 in the first string and 0 to 1 in the second string, for this problem is:
A
A
Your alignment should have the same score as the sample alignment: 1
Sample Input 4:

1 1 1
TAACG
ACGTG
Sample Output 4:

A sample correct alignment, aligning indices 2 to 5 in the first string and 0 to 3 in the second string, for this problem is:
ACG
ACG
Your alignment should have the same score as the sample alignment: 3
Sample Input 5:

3 2 1
CAGAGATGGCCG
ACG
Sample Output 5:

A sample correct alignment, aligning indices 10 to 12 in the first string and 1 to 3 in the second string, for this problem is:
CG
CG
Your alignment should have the same score as the sample alignment: 6
Sample Input 6:

2 3 1
CTT
AGCATAAAGCATT
Sample Output 6:

A sample correct alignment, aligning indices 0 to 3 in the first string and 9 to 13 in the second string, for this problem is:
C-TT
CATT
Your alignment should have the same score as the sample alignment: 5
*/
func LocalAlignment(str1, str2 string, match, mismatch, gap float64) (Alignment, int, int, int, int) {
	//Get the score table
	table := LocalScoreTable(str1, str2, match, mismatch, gap)

	var a Alignment

	maxScore := 0.0
	//We need the position in order to find the location of the end of the match
	trackRow := 0
	trackCol := 0
	//find max score
	for r := 0; r < len(table); r++ {
		for c := 0; c < len(table[0]); c++ {
			if table[r][c] > maxScore {
				maxScore = table[r][c]
				trackRow = r
				trackCol = c
			}
		}
	}

	ra := trackRow
	rb := trackCol
	//backtrack, add string to a
	for table[trackRow][trackCol] >= 0 && trackRow > 0 && trackCol > 0 {
		//In case of a match
		// Change r to trackRow!!!!
		if str1[trackRow-1] == str2[trackCol-1] && table[trackRow][trackCol] == table[trackRow-1][trackCol-1]+match {
			a[0] = string(str1[trackRow-1]) + a[0]
			a[1] = string(str2[trackCol-1]) + a[1]
			trackRow--
			trackCol--
		} else if str1[trackRow-1] != str2[trackCol-1] && table[trackRow][trackCol] == table[trackRow-1][trackCol-1]-mismatch {
			a[0] = string(str1[trackRow-1]) + a[0]
			a[1] = string(str2[trackCol-1]) + a[1]
			trackRow--
			trackCol--
		} else if table[trackRow][trackCol] == table[trackRow-1][trackCol]-gap {
			a[0] = string(str1[trackRow-1]) + a[0]
			a[1] = "-" + a[1]
			trackRow--
		} else if table[trackRow][trackCol] == table[trackRow][trackCol-1]-gap {
			a[0] = "-" + a[0]
			a[1] = string(str2[trackCol-1]) + a[1]
			trackCol--
		} else {
			break
		}
	}
	return a, trackRow, ra, trackCol, rb
}

//I renamed these:
//	maxRow -> trackRow
//	maxCol -> trackCol
