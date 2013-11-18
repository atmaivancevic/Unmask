package main

import (
	"fmt"
	"os"
)

func main() {

	// // go to the Marmoset directory
	// os.Chdir("/scratch/atmaGenomes/Marmoset")

	// // unmask chr1-22,X,Y,Un
	// unmask("chr1.fa", "chr1_unmasked.fa")
	// unmask("chr2.fa", "chr2_unmasked.fa")
	// unmask("chr3.fa", "chr3_unmasked.fa")
	// unmask("chr4.fa", "chr4_unmasked.fa")
	// unmask("chr5.fa", "chr5_unmasked.fa")
	// unmask("chr6.fa", "chr6_unmasked.fa")
	// unmask("chr7.fa", "chr7_unmasked.fa")
	// unmask("chr8.fa", "chr8_unmasked.fa")
	// unmask("chr9.fa", "chr9_unmasked.fa")
	// unmask("chr10.fa", "chr10_unmasked.fa")
	// unmask("chr11.fa", "chr11_unmasked.fa")
	// unmask("chr12.fa", "chr12_unmasked.fa")
	// unmask("chr13.fa", "chr13_unmasked.fa")
	// unmask("chr14.fa", "chr14_unmasked.fa")
	// unmask("chr15.fa", "chr15_unmasked.fa")
	// unmask("chr16.fa", "chr16_unmasked.fa")
	// unmask("chr17.fa", "chr17_unmasked.fa")
	// unmask("chr18.fa", "chr18_unmasked.fa")
	// unmask("chr19.fa", "chr19_unmasked.fa")
	// unmask("chr20.fa", "chr20_unmasked.fa")
	// unmask("chr21.fa", "chr21_unmasked.fa")
	// unmask("chr22.fa", "chr22_unmasked.fa")

	// unmask("chrX.fa", "chrX_unmasked.fa")
	// unmask("chrY.fa", "chrY_unmasked.fa")
	// unmask("chrUn.fa", "chrUn_unmasked.fa")

	// go to the Medaka directory
	// os.Chdir("/scratch/atmaGenomes/Medaka")

	// // unmask chr1-24
	// unmask("chr1.fa", "chr1_unmasked.fa")
	// unmask("chr2.fa", "chr2_unmasked.fa")
	// unmask("chr3.fa", "chr3_unmasked.fa")
	// unmask("chr4.fa", "chr4_unmasked.fa")
	// unmask("chr5.fa", "chr5_unmasked.fa")
	// unmask("chr6.fa", "chr6_unmasked.fa")
	// unmask("chr7.fa", "chr7_unmasked.fa")
	// unmask("chr8.fa", "chr8_unmasked.fa")
	// unmask("chr9.fa", "chr9_unmasked.fa")
	// unmask("chr10.fa", "chr10_unmasked.fa")
	// unmask("chr11.fa", "chr11_unmasked.fa")
	// unmask("chr12.fa", "chr12_unmasked.fa")
	// unmask("chr13.fa", "chr13_unmasked.fa")
	// unmask("chr14.fa", "chr14_unmasked.fa")
	// unmask("chr15.fa", "chr15_unmasked.fa")
	// unmask("chr16.fa", "chr16_unmasked.fa")
	// unmask("chr17.fa", "chr17_unmasked.fa")
	// unmask("chr18.fa", "chr18_unmasked.fa")
	// unmask("chr19.fa", "chr19_unmasked.fa")
	// unmask("chr20.fa", "chr20_unmasked.fa")
	// unmask("chr21.fa", "chr21_unmasked.fa")
	// unmask("chr22.fa", "chr22_unmasked.fa")

	// unmask("chrX.fa", "chrX_unmasked.fa")
	// unmask("chrY.fa", "chrY_unmasked.fa")
	// unmask("chrUn.fa", "chrUn_unmasked.fa")

	for i := 1; i <= 3; i++ {
		// unmask("test"+(i)+"fa", "test_unmasked"+(i)+"fa")
		fmt.Println("test" + (i) + ".fa")
	}
}
