package main

// LocalScoreTable takes two strings and alignment penalties. It returns a 2-D array
// holding dynamic programming scores for local alignment with these penalties.
func LocalScoreTable(str1, str2 string, match, mismatch, gap float64) [][]float64 {
	fill := Matrix(len(str1)+1, len(str2)+1)

	for row := 1; row < len(str1)+1; row++ {
		fill[row][0] = float64(row) * gap * (-1)
	} //vertical, fill in edge

	for col := 1; col < len(str2)+1; col++ {
		fill[0][col] = float64(col) * gap * (-1)
	} //horizontal, fill the edge

	//Calculating the values inside the matrixs
	for row := 1; row <= len(str1); row++ { //start from 1 because 0 is already filled
		for col := 1; col <= len(str2); col++ {

			//Calculate values for the three directions
			up := fill[row-1][col] - gap
			left := fill[row][col-1] - gap
			diag := fill[row-1][col-1]

			//Two cases: match or mismatch for diag
			if str1[row-1] == str2[col-1] { //match
				diag += match
			} else {
				diag -= mismatch
			}

			//Maximum score of three directions
			max3 := Max(diag, left, up, 0.0)
			fill[row][col] = max3
		}
	}

	return fill
}
