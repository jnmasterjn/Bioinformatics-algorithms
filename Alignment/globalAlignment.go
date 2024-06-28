package main

type Alignment [2]string

// GlobalAlignment takes two strings, along with match, mismatch, and gap scores.
// It returns a maximum score global alignment of the strings corresponding to these penalties.
func GlobalAlignment(str1, str2 string, match, mismatch, gap float64) Alignment {
	var a Alignment

	//The final strings
	ms1 := ""
	ms2 := ""

	//Make the score table
	lcM := GlobalScoreTable(str1, str2, match, mismatch, gap)

	//row and column
	r := len(lcM) - 1
	c := len(lcM[0]) - 1

	//we want to end the loop when we reach the matrix[0][0] (well)
	//starts from the end (sink)
	for r > 0 && c > 0 {
		num := lcM[r][c]

		//This part needs to be modified
		// if going diagonal increases the score, matched
		if lcM[r-1][c-1]+match == num && str1[r-1] == str2[c-1] {
			//Add the letters
			ms1 = string(str1[r-1]) + ms1
			ms2 = string(str2[c-1]) + ms2
			r--
			c--
			//add letter to both ms1 and ms2 if match
		} else if lcM[r-1][c]-gap == num {
			ms1 = string(str1[r-1]) + ms1
			ms2 = "-" + ms2
			r--
			//add dash to ms1, add letter to ms2
		} else if lcM[r][c-1]-gap == num {
			ms2 = string(str2[c-1]) + ms2
			ms1 = "-" + ms1
			c--
		} else { //mismatch
			//add the different letters to corresponding message
			ms1 = string(str1[r-1]) + ms1
			ms2 = string(str2[c-1]) + ms2
			r--
			c--

		}

	}

	//Compensating for the dashes at the beginning 😭 😭😭😭😭😭
	for r > 0 {
		ms1 = "-" + ms1
		ms2 = string(str2[r]) + ms2
		r--
	}
	for c > 0 {
		ms2 = "-" + ms2
		ms1 = string(str1[c]) + ms1
		c--
	}

	//Setting the alignment variable
	a[0] = ms1
	a[1] = ms2

	return a
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
