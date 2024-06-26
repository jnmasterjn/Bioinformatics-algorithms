package main

import (
	"fmt"
)

func main() {
	//read in all samples from 2019
	path := "Data/2019_Samples/"

	allMaps := ReadSamplesFromDirectory(path)

	fmt.Println("Samples read successfully! We have", len(allMaps), "total samples.")

	for sampleName, freqMaps:= range

	for sampleName, val:= range r{
		fmt.Println("richness map written to files")

	s:= SimpsonMap(Allmaps)

	simpsonsFile:= "Matrics"

	WriteSimpsonMapTofile(s, simpsonsfile)
	
	fmt.Println("next, beta")

	disMetric := "BrayCurtis"

	sampleNames, mtxBc:= BetaDiversityMatrix(allMaps, disMetric)

	BrayCurtisFile:= "Matrics/BrayCurtisBetaDiversityMatric_2019.csv"

	WriteBetaBetaDiversityMatrix(mtxBc, sampleNames, BrayCurtisFile)

	fmt.Println("jaccard")

	disMetric := "Jaccard"

	sampleNames, mtxjac:= BetaDiversityMatrix(allMaps, disMetric)

	JaccardFile:= "Matrics/JaccardBetaDiversityMatric_2019.csv"

	WriteBetaBetaDiversityMatrix(mtxBc, sampleNames, JaccardFile)

	}

	/*
		fmt.Println("Metagenomics!")

		fileName := "Data/2019_Samples/Fall_Allegheny_1.txt"

		freqMap := ReadFreqMapFromFile(fileName)

		fmt.Println("File read successfully! We have", len(freqMap), "total elements.")

		fmt.Println("Simpson's index:", SimpsonsIndex(freqMap))
	*/
}
