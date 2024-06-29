package main

//LongestCommonSubsequence takes two strings as input.
//It returns a longest common subsequence of the two strings.

func LongestCommonSubsequence(str1, str2 string) string {
	//n, m:= len(str1), len(str2)
	match := ""

	//Longest commmon matrix
	lcM := LCMatrix(str1, str2)

	//row and column
	r := len(lcM) - 1
	c := len(lcM[0]) - 1

	//we want to end the loop when we reach the matrix[0][0] (well)
	//starts from the end (sink)
	for r != 0 && c != 0 {
		num := lcM[r][c]
		// if going diagonal increases the score
		if lcM[r-1][c-1]-1 == num && str1[r-1] == str2[c-1] {
			match = string(str1[r-1]) + match
			r = r - 1
			c = c - 1
		} else if lcM[r-1][c] == num {
			r = r - 1
		} else if lcM[r][c-1] == num {
			c = c - 1
		} else {
			r = r - 1
			c = c - 1
		}

	}
	return match

}

func LCMatrix(str1, str2 string) [][]int {
	if len(str1) == 0 || len(str2) == 0 {
		panic("empty strings given")
	}
	fill := Matrix(len(str1)+1, len(str2)+1)
	for row := 1; row <= len(str1); row++ {
		for col := 1; col <= len(str2); col++ {
			if str1[row-1] == str2[col-1] {
				fill[row][col] = 1 + fill[row-1][col-1]
			} else {
				fill[row][col] = max2(fill[row-1][col], fill[row][col-1])
			}
		}
	}
	return fill
}
