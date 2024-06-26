package main

//SimpsonsMap takes a map mapping sample IDs to frequency maps.package solution
//It returns a map of sample IDs to Simpson's indices for each sample.
func SimpsonsMap(allMaps map[string](map[string]int)) map[string]float64 {
	s := make(map[string]float64)

	//range through sample IDs and set appropriate richness value
	for sampleName, freqMap := range allMaps {
		s[sampleName] = SimpsonsIndex(freqMap)
	}

	return s
}
