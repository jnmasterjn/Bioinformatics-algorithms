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
		//Making the diversity matrix
		//Use sort.Strings()
	
		//Decides which metric to use
		
	
		//Sample array
		//To loop through the 2D array, could we use allMaps[samples[index]], so the order of the 2d array would be correct
		samples := make([]string,0) //make rows first, this stores the key, the sample name
		for sName := range allMaps { //range over the keys to get the sample names
			samples = append(samples, sName) //append the sample names to the array, append(theslice,what to append)
		}
		sort.Strings(samples) //sort the name of sample in alpabet order

		//2D array, the distance matrix
		bDiv := make([][]float64, len(allMaps)) //make columns, this stores the value of distance
		for value := range bDiv { 
			bDiv[value] = make([]float64, len(allMaps))
		}
		
		//Loop through the matrix
		for i:= 0; i<len(samples); i++{
			for j:=0; j<len(samples);j++{
				bDiv[i][j] = CheckDistanceMetrics(distanceMetric, allMaps[samples[i]], allMaps[samples[j]])
			}
		}
		return samples, bDiv
	}
	
	func CheckDistanceMetrics(distMetric string, map1, map2 map[string]int)float64 { 
		//check input to know which method to calculate, could be jaccard or bray-curtis
		switch distMetric{
			case "Jaccard":
				return JaccardDistance(map1, map2)
			case "Bray-Curtis":
				return BrayCurtisDistance(map1, map2)
			default :
				panic("Unsupported string")
		}
		/*if distMetric == "Jaccard" {
			return JaccardDistance(map1, map2)
		} else if distMetric == "BrayCurtis"{
			return BrayCurtisDistance(map1, map2)
		}
		return 0*/
	}
	
	func JaccardDistance(map1 map[string]int, map2 map[string]int) float64 {
		c := SumOfMinima(map1, map2)
		t := SumOfMaxima(map1, map2)
		j := 1 - (float64(c) / float64(t))
		return j
	}
	
	func BrayCurtisDistance(map1 map[string]int, map2 map[string]int) float64 {
		c := SumOfMinima(map1, map2)
		s1 := SampleTotal(map1)
		s2 := SampleTotal(map2)
	
		//don't allow the situation in which we have zero richness.
		if s1 == 0 || s2 == 0 {
			panic("Error: sample given to BrayCurtisDistance() has no positive values.")
		}
	
		av := Average(float64(s1), float64(s2))
		return 1 - (float64(c) / av)
	}
	
	/*Code Challenge: Implement a function BetaDiversityMatrix(allMaps, distanceMetric).
	
	Input: a map allMaps whose keys are strings corresponding to sample IDs and whose values are frequency maps (each such frequency map corresponds to a single sample and maps strings to integers), followed by a string distanceMetric. 
	Return: a sorted list sampleNames of the sample IDs, as well as a 2-D array D such that D[i][j] is the distance from samples[i] to samples[j] using the distance metric indicated by distanceMetric.
	Note: distanceMetric should only be able to have one of two values: "Bray-Curtis" or "Jaccard".
	
	Hint: You may find helpful the command sort.Strings() from the sort package, which takes a slice of strings as input and sorts the strings in the slice.
	
	Functions that are implemented behind the scenes for you (in addition to subroutines that they call):
	
	BrayCurtisDistance(): takes as input two maps of strings to integers, and returns the Bray-Curtis distance between them.
	JaccardDistance(): takes as input two maps of strings to integers, and returns the Jaccard distance between them.

