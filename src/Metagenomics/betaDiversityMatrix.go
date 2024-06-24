package main

//BetaDiversityMatrix takes a map of frequency maps along with a distance metric
//("Bray-Curtis" or "Jaccard") as input.
//It returns a slice of strings corresponding to the sorted names of the keys
//in the map, along with a matrix of distances whose (i,j)-th
//element is the distance between the i-th and j-th samples using the input metric.
// Input: a collection of frequency maps samples and a distance metric
// Output: a list of sample names and a distance matrix where D_i,j is the distance between
// sample_i and sample_j according to the given distance metric
func BetaDiversityMatrix(allMaps map[string](map[string]int), distanceMetric string) ([]string, [][]float64) {

}
