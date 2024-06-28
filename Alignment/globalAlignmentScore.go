package main

// GlobalScoreTable takes two strings and alignment penalties. It returns a 2-D array
// holding dynamic programming scores for global alignment with these penalties.
func GlobalScoreTable(str1, str2 string, match, mismatch, gap float64) [][]float64 {

	fill := Matrix(len(str1)+1, len(str2)+1)

	for row := 1; row < len(str1)+1; row++ {
		fill[row][0] = float64(row) * gap * (-1)
	}

	for col := 1; col < len(str2)+1; col++ {
		fill[0][col] = float64(col) * gap * (-1)
	}

	//Calculating the values inside the matrixs
	for row := 1; row <= len(str1); row++ {
		for col := 1; col <= len(str2); col++ {

			//Calculate values for the three directions
			up := fill[row-1][col] - gap
			left := fill[row][col-1] - gap
			diag := fill[row-1][col-1]

			//Two cases: match or mismatch
			if str1[row-1] == str2[col-1] { //match
				diag += match
			} else {
				diag -= mismatch
			}

			//Maximum score of three directions
			max3 := Max(diag, left, up)
			fill[row][col] = max3
		}
	}

	return fill
}

func Matrix(len1, len2 int) [][]float64 {
	arr := make([][]float64, len1)

	for i := 0; i < len1; i++ {
		arr[i] = make([]float64, len2)
	}
	return arr
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
