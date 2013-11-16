package main

import (
	"code.google.com/p/biogo/alphabet"
	"code.google.com/p/biogo/io/seqio"
	"code.google.com/p/biogo/io/seqio/fasta"
	"code.google.com/p/biogo/seq/linear"

	"bufio"
	"fmt"
	"os"
	"strings"
)

// use this function to unmask (capitalize) nucleotide sequences
func unmask(inputFileName, outputFileName string) {
	// open inputFile in read-only mode
	inputFile, inputError := os.Open(inputFileName)
	if inputError != nil {
		fmt.Println("An error occurred opening the input file.")
		return
	}
	defer inputFile.Close()

	// open outputFile in write-only mode
	outputFile, outputError := os.Create(outputFileName)
	if outputError != nil {
		fmt.Println("An error occurred with file creation.")
		return
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)

	sc := seqio.NewScanner(fasta.NewReader(inputFile, linear.NewSeq("", nil, alphabet.DNA)))
	for sc.Next() {
		s := sc.Seq().(*linear.Seq)
		for i := range s.Seq {
			s.Seq[i] &^= ' '
		}
		// get rid of pipes in the seq ID
		s.ID = strings.Replace(s.ID, "|", "", -1)
		// write header
		outputWriter.WriteString(">" + s.ID + "\n")
		// write FASTA sequence
		outputWriter.WriteString(s.String() + "\n")
	}
	if sc.Error() != nil {
		panic(sc.Error())
	}

	// write the buffer completely to the file
	outputWriter.Flush()
}
