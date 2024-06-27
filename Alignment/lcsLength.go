/*Code Challenge: Implement a function LCSlength(str1, str2).

Input: two strings, str1 and str2.
Return: the length of a longest common subsequence of the two strings. If no common subsequence exists, your function should return 0.
Test case documentation (you must be logged into CMU account to view)

Sample Input 1:

GACT
ATG
Sample Output 1:

2
Sample Input 2:

ACTGAG
GACTGG
Sample Output 2:

5
Sample Input 3:

AC
AC
Sample Output 3:

2
Sample Input 4:

GGGGT
CCCCT
Sample Output 4:

1
Sample Input 5:

TCCCC
TGGGG
Sample Output 5:

1
Sample Input 6:

AA
CGTGGAT
Sample Output 6:

1
Sample Input 7:

GGTGACGT
CT
Sample Output 7:

2
Sample Input 8:

yes
no
Sample Output 8:
0
Sample Input 9:

ACCGTCTTAGCGATCAACACATTTAACAACGCGCCGCACCCCCCGTCAAACGAGCTTTTGGGCTCTTGTCCTTTTACAAGCTTCACGACGCATACAGCCTTGATCAACGGTTTGATCTGTCTCCCTTCAGCTGGCTTTAAAGGACATACATATGAAGGCCTTAATAAGGTCCGGGAACTCCACATATTCGGTACTGGGCAAACCCCATGAACCACCTCAACATGAAGAGTCCGAGGACTCTCACGATCCACCAATGCAGATCGGAACTGTGCGATCGCGTAATGAGCCGAGTACTTGGTTTGTGTTTAGGTTATGGGGGCCGGGAGCCGGTTCAATATAAGGAAGTAGTTGCAGATTAGTTGTTGCGAACGGTCATAAATTTGATGGGTAAACGTGAACTTAACAAACCGTGATAGCTAATCCTATGCATCCCTTACGTGGATCGACTCGAGTACCCAGGTGAACCGACTACTTGATAACCGGAAATCGCGGTATAAAAGCGCTCACGGTCAGGAGATATACCTCCAAGCAGTAGTCTTTCTGAGCCTAGAGTAGTAAATTACAGGGACGATGTCTTTTACCGAGGCAACATTTTATTGAGAATCACATGAGGCACAGGTAAAGGCGACATCACGATCGAGATCAACCCCTACTTGTTCAAAACATTGAGAACCAGCTCTGTTTTGGAACCTAGAAAGATAACGCATCCGCTTGATATTCCACGGCTTGTCCCTCTTGTGCGGTCCATCTATCGGAGTTTCCTCCGATACGACCCGCAATGTTTCCAGGCGTACGGTACTTTATGAATACACTCGCGCTGTAACCTGTTATGTGAAACACACACGACAGAGCTTCGCGTGGGCCCAGCGACCCGGTAATACTACATCACCGCACACGACCTCGAGCAGTCTTTGCCGGCGTCCGTAAGTAGTCTAAAGTTGTGTTGATGCTTGGGGTTAAAGCTAAATCGTCCGCAGAATACGACTCTCATCCCAAT
ACCCGCACGCGCTTTGGTCTAGATTCTAGCTCCAACTTGCCTGCTAGATACTCTGTTAAAAGATGGTTTTACAACCCCCTCCTCTGTCCCTGGGGTATTATATAATACGTCGGATAGTCAGGTACAAATACAAGTGGGTGGGAATACTTTTCCTCGGATCCTAGACCACGGATTACTGCGTGGTTGACAAGAGTCGGCCCGGAGGGAAACGTGAAGGTTAGTGCAATTAAAGTCTCTAATGTGAAGCCTCCGCGAAGCGAGGAGTTTCTGAGATCGAGTACTATTTAGAGTTCGAAATCACGGCTTAACCTCACTGCCACGCATAACTTGCCGGCAATCCAGTTTTGCAACGATACTTAATTTGTGCAGCTCATCTTTGCTGTCCAGAAATAGAGCTAGTCGATCTCATCTTGCGGGTAGCCAGAAGTCCTACCGTCTCCTCCATGTAGCTTAAAAATTTCGGTGAGGATCAAAAATGATAAACGTGACAGGTAAGCTCCTACGTCTATCCTATGACCCCCGCGGCAGAATAGGTTGGTAGTGTTAGTGCGTGAGCTGGTAGAATAGAGCACACTTAGGGAAACGGGAACCGTTATGTAGGGCTGCGACACACAAAAAAGTGTTCGTTGGTAAGCTGCCTCTCCACTAAACAGGATTTCTCTGGATGATCCCATCGAAGCAAGTTACGCACCACGCCGAGGCGGACCCTGGTACTAGCTGCCCCCCCCTTTATGGGGCGCTCGTACATCAAGATGATCGCGGACTCAACCTGATTACGAGTTGTCCAAGTAGTCCAGGGTAAGAGAAACTGGAGAGA
Sample Output 9:

573
*/

package main

//LCSLength takes two strings as input. It returns the length of a longest common
//subsequence of the two strings.

func LCSLength(str1, str2 string) int {
	fill := Matrix(len(str1)+1, len(str2)+1)
	for row := 1; row <= len(str1); row++ {
		for col := 1; col <= len(str2); col++ {
			if str1[row-1] == str2[col-1] { //whether top left is the same for both string
				fill[row][col] = 1 + fill[row-1][col-1] //matched, so plus 1
			} else {
				fill[row][col] = max2(fill[row-1][col], fill[row][col-1]) //if not matched, find the bigger sum
			}
		}
	}
	return fill[len(str1)][len(str2)]
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Matrix(len1, len2 int) [][]int {
	arr := make([][]int, len1)

	for i := 0; i < len1; i++ {
		arr[i] = make([]int, len2)
	}
	return arr
}
