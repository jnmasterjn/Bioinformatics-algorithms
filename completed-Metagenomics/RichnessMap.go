package main

//RichnessMap takes a map of frequency maps as input.  It returns a map
//whose values are the richness of each sample.
func RichnessMap(allMaps map[string](map[string]int)) map[string]int {
	r := make(map[string]int)

	//range through sample IDs and set appropriate richness value
	for sampleName, freqMap := range allMaps {
		r[sampleName] = Richness(freqMap)
	}

	return r
}
