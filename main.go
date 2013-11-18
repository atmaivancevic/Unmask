package main

import (
	"fmt"
	"os"
	"strconv"
)

const genomeDir = "/scratch/atmaGenomes/"

// this runs unmask on the file [prefix][identifier].[extension] within the specific genome directory genomeDir/genomeName
// prefix = chr/scaffold (for genomes where the files are seperated into chr/scaffolds);
// or scientific name (for genomes with a single genome files)
func unmaskChr(genomeName, prefix, extension, identifier string) {

	os.Chdir(genomeDir + genomeName)

	inputFileName := prefix + identifier + "." + extension
	outputFileName := prefix + identifier + "_unmasked." + extension

	// if the unmasked file does not exist, create it
	if _, err := os.Stat(outputFileName); os.IsNotExist(err) {
		unmask(inputFileName, outputFileName)
		fmt.Println("Processed: " + genomeName + "/" + prefix + identifier)
	} else {
		fmt.Println(prefix + identifier + " already unmasked; skipping.")
	}

}

// run unmask on a range of values
func unmaskChrRange(genomeName, prefix, extension string, startIdentifier, endIdentifier int) {
	for i := startIdentifier; i <= endIdentifier; i++ {
		unmaskChr(genomeName, prefix, extension, strconv.Itoa(i))
	}
}

func main() {

	unmaskChrRange("Medaka", "chr", "fa", 1, 24)

	unmaskChrRange("Mouse", "chr", "fa", 1, 19)
	unmaskChr("Mouse", "chr", "fa", "X")
	unmaskChr("Mouse", "chr", "fa", "Y")
	unmaskChr("Mouse", "chr", "fa", "Un")

}
