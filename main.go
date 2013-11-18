package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var genomeName, inputFileName, outputFileName string

	// go to the Medaka directory
	genomeName = "Medaka"
	os.Chdir("/scratch/atmaGenomes/" + genomeName)

	// unmask chr1-24
	for i := 1; i <= 24; i++ {
		inputFileName = "chr" + strconv.Itoa(i) + ".fa"
		outputFileName = "chr" + strconv.Itoa(i) + "_unmasked.fa"
		unmask(inputFileName, outputFileName)
		fmt.Println("Processed: " genomeName + "/" + inputFileName)
	}

	// go to the Mouse directory
	genomeName = "Mouse"
	os.Chdir("/scratch/atmaGenomes/" + genomeName)

	// unmask chr1-19, X, Y, Un
	for i := 1; i <= 19; i++ {
		inputFileName = "chr" + strconv.Itoa(i) + ".fa"
		outputFileName = "chr" + strconv.Itoa(i) + "_unmasked.fa"
		unmask(inputFileName, outputFileName)
		fmt.Println("Processed: " genomeName + "/" + inputFileName)
	}
	
	inputFileName = "chrX.fa"
	outputFileName = "chrX_unmasked.fa"
	unmask(inputFileName, outputFileName)
	fmt.Println("Processed: " genomeName + "/" + inputFileName)

	inputFileName = "chrY.fa"
	outputFileName = "chrY_unmasked.fa"
	unmask(inputFileName, outputFileName)
	fmt.Println("Processed: " genomeName + "/" + inputFileName)

	inputFileName = "chrUn.fa"
	outputFileName = "chrUn_unmasked.fa"
	unmask(inputFileName, outputFileName)
	fmt.Println("Processed: " genomeName + "/" + inputFileName)

}
