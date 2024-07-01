package main

import "fmt"

// func main() {
// 	fmt.Println("Sequence alignment!")
// }

func Hemoglobin() {
	ZebraFish := ReadFASTAFile("Data/Hemoglobin/Danrio_rerio_hemoglobin.fasta")

	Cow := ReadFASTAFile("Data/Hemoglobin/Bos_taurus_hemoglobin.fasta")

	Gorilla := ReadFASTAFile("Data/Hemoglobin/Gorilla_gorilla_hemoglobin.fasta")

	Human := ReadFASTAFile("Data/Hemoglobin/Homo_sapiens_hemoglobin.fasta")

	match := 1.0
	mismatch := 1.0
	gap := 3.0
}

func SARSAlignment() {

	//read in 2 files
	fmt.Println("Reading in genome files.")
	sars1 := ReadFASTAFile("Data/Coronaviruses/SARA-CoV_genome.fasta")
	sars2 := ReadFASTAFile("Data/Coronaviruses/SARA-CoV-2_genome.fasta")

	fmt.Println("Files read. Now, aligning genomes")

	GlobalAlignment(sars1, sars2, 1, 1, 2)

	WriteAlignmentToFASTA(sarsAlignment, "Output/sarasGenoneAlignment.txt")

	fmt.Println("Alignment written to files")
}
