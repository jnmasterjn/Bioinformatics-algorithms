package main

// EditDistance takes two strings as input. It returns the Levenshtein distance
// between the two strings; that is, the minimum number of substitutions, insertions, and deletions
// needed to transform one string into the other.
func EditDistance(str1, str2 string) int {
	dist := Matrix(len(str1)+1, len(str2)+1)

	for row := 1; row < len(str1)+1; row++ {
		dist[row][0] = row
	}

	for col := 1; col < len(str2)+1; col++ {
		dist[0][col] = col
	}

	for row := 1; row <= len(str1); row++ {
		for col := 1; col <= len(str2); col++ {

			//If the string's letters are the same, then the diagonal value will not increase
			if str1[row-1] == str2[col-1] {
				dist[row][col] = dist[row-1][col-1]
			} else {
				//🌪️🌪️🌪️🌪️🌪️🌪️🌪️🌪️🌪️🌪️🌪️🌪️🌪️🌪️🌪️
				//Otherwise, find the minimum of the diagonal, left index, and upwards index
				dist[row][col] = min(dist[row-1][col], dist[row][col-1], dist[row-1][col-1]) + 1
			}
		}
	}

	return dist[len(dist)-1][len(dist[0]-1)]
}

func Matrix(len1, len2 int) [][]int {
	arr := make([][]int, len1)

	for i := 0; i < len1; i++ {
		arr[i] = make([]int, len2)
	}
	return arr
}

func min(ur ...int) int {
	if len(ur) < 1 {
		panic("Input slice length is too little")
	}
	var least int
	for i, val := range ur {
		if i == 0 || val < least {
			least = val
		}
	}
	return least
}
