package main

import (
	"os"
)

func main() {
	os.Chdir("/scratch/atmaGenomes/Marmoset")
	unmask("test.fa", "test_unmasked.fa")
}
