package main

import "sort"

//BetaDiversityMatrix takes a map of frequency maps along with a distance metric
//("Bray-Curtis" or "Jaccard") as input.
//It returns a slice of strings corresponding to the sorted names of the keys
//in the map, along with a matrix of distances whose (i,j)-th
//element is the distance between the i-th and j-th samples using the input metric.
// Input: a collection of frequency maps samples and a distance metric
// Output: a list of sample names and a distance matrix where D_i,j is the distance between
// sample_i and sample_j according to the given distance metric
func BetaDiversityMatrix(allMaps map[string](map[string]int), distanceMetric string) ([]string, [][]float64) {

	sampleNames := make([]string, 0, len(allMaps))

	numSamples := len(allMaps)

	// first, grab all the sample names and sort them

	for name := range allMaps {
		sampleNames = append(sampleNames, name)
	}

	// sort the names
	sort.Strings(sampleNames)

	//now, form the distance matrix
	mtx := InitializeSquareMatrix(numSamples)

	// the main diagonal values are all zero by default, which is good!

	// range through matrix and set all our values
	for i := 0; i < numSamples; i++ {
		for j := i + 1; j < numSamples; j++ {
			//let's grab the maps that we are interested in
			//we want the samples with names sampleNames[i] and sampleNames[j]
			map1 := allMaps[sampleNames[i]]
			map2 := allMaps[sampleNames[j]]

			//now, compute distance
			var dist float64

			//check which metric we are using
			if distanceMetric == "Bray-Curtis" {
				dist = BrayCurtisDistance(map1, map2)
			} else if distanceMetric == "Jaccard" {
				dist = JaccardDistance(map1, map2)
			} else {
				panic("Error: invalid distance metric given to BetaDiversityMatrix")
			}

			// set two values of mtx that are both equal to dist based on symmetry
			mtx[i][j] = dist
			mtx[j][i] = dist
		}
	}
	return sampleNames, mtx
}

//InitializeSquareMatrix takes an integer n as input.
//It returns an n x n matrix of float64 variables.
func InitializeSquareMatrix(n int) [][]float64 {
	//first, make the slice
	mtx := make([][]float64, n)
	for i := 0; i < n; i++ {
		// in here, we make the rows
		mtx[i] = make([]float64, n)
	}

	return mtx
}
