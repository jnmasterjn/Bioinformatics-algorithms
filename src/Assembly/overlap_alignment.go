package main

//ALL PENALTIES POSITIVE

//ScoreOverlapAlignment takes two strings along with match, mismatch, and gap penalties.
//It returns the maximum score of an overlap alignment in which str1 is overlapped with str2.
//Assume we are overlapping a suffix of str1 with a prefix of str2.

// Sample Input 1:
// 1 1 2
// GAGA
// GAT

// Sample Output 1:
// 2

// Sample Input 2:
// 1 1 1
// CCAT
// AT

// Sample Output 2:
// 2

// Sample Input 3:
// 1 5 1
// CAT
// GAT

// Sample Output 3:
// 1

func ScoreOverlapAlignment(str1, str2 string, match, mismatch, gap float64) float64 {
	matrix := Matrix(len(str1)+1, len(str2)+1) //make a matrix

	//set first row into negative numbers, first row col is all 0
	//whole matrix start from row 1 col 1, not row 0 col 0

	for col := 1; col < len(str2)+1; col++ {
		matrix[0][col] = float64(col) * gap * (-1)
	}

	for row := 1; row <= len(str1); row++ {
		for col := 1; col <= len(str2); col++ {
			up := matrix[row-1][col] - gap
			left := matrix[row][col-1] - gap
			diag := matrix[row-1][col-1]

			//Two cases: match or mismatch
			if str1[row-1] == str2[col-1] { //match
				diag += match
			} else {
				diag -= mismatch
			}

			//Maximum score of three directions
			max3 := Max(diag, left, up)
			matrix[row][col] = max3
		}
	}
	score := 0.0

	for i := 0; i < len(matrix[0]); i++ { //range over last row
		if matrix[len(matrix)-1][i] > score { //range over last row every col
			score = matrix[len(matrix)-1][i]
		}
	}
	return score
}

func Max(x ...float64) float64 {
	var max float64
	for i, val := range x {
		if i == 0 || val > max {
			max = val
		}
	}
	return max
}

func Matrix(len1, len2 int) [][]float64 {
	arr := make([][]float64, len1)

	for i := 0; i < len1; i++ {
		arr[i] = make([]float64, len2)
	}
	return arr
}
