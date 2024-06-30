package main

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
