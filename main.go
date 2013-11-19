package main

import (
	"fmt"
	"os"
	"strconv"
)

const genomeDir = "/scratch/atmaGenomes/"

// this runs unmask on the file [prefix][identifier].[extension] within the specific genome directory [genomeDir]/[genomeName]

// prefix = chr/scaffold (only exists if the genome has been divided into seperate, smaller files)
// e.g. prefix for Mouse is "chr"
// identifier = the unique identifier for each chromosome/scaffold/genome
// e.g. the identifiers for Mouse are 1-19, X, Y, Un (since the Mouse genome is seperated into chromosomes)
// but the identifier for Frog is xentro3 (since the Frog genome is in a single file )
// extension = fa or fasta

// run unmask on a single identifier
// note: although called unmaskChr, also used for unmasking full genomes, if no seperate chr files available
func unmaskChr(genomeName, prefix, extension, identifier string) {

	// move to the appropriate genome directory
	os.Chdir(genomeDir + genomeName)

	inputFileName := prefix + identifier + "." + extension
	outputFileName := prefix + identifier + "_unmasked." + extension

	// if the unmasked file does not exist, create it
	if _, err := os.Stat(outputFileName); os.IsNotExist(err) {
		unmask(inputFileName, outputFileName)
		fmt.Println("Processed: " + genomeName + "/" + prefix + identifier)
	} else {
		fmt.Println(genomeName + "/" + prefix + identifier + " already unmasked; skipping.")
	}
}

// run unmask on a range of identifiers (numbers)
func unmaskChrRange(genomeName, prefix, extension string, startIdentifier, endIdentifier int) {
	for i := startIdentifier; i <= endIdentifier; i++ {
		unmaskChr(genomeName, prefix, extension, strconv.Itoa(i))
	}
}

func main() {

	// arrays for genomes that are seperated into chromosomes/scaffolds

	// array for cases using unmaskChrRange
	for _, gen := range []struct {
		genomeName      string
		prefix          string
		startIdentifier int
		endIdentifier   int
	}{
		{"A_mellifera", "Group", 1, 16},
		{"Baboon", "chr", 1, 20},
		{"Btau6", "Chr", 1, 29},
		{"Cat", "chrA", 1, 3},
		{"Cat", "chrB", 1, 4},
		{"Cat", "chrC", 1, 2},
		{"Cat", "chrD", 1, 4},
		{"Cat", "chrE", 1, 3},
		{"Cat", "chrF", 1, 2},
		{"Chicken", "chr", 1, 28},
		{"Chimp", "chr", 3, 22},
		{"Dog", "chr", 1, 38},
		{"equCab2", "chr", 1, 31},
		{"Gibbon", "chr", 2, 6},
		{"Gibbon", "chr", 8, 21},
		{"Gibbon", "chr", 23, 25},
		{"Gorilla", "chr", 3, 22},
		{"Human", "chr", 1, 22},
		{"Lizard", "chr", 1, 6},
		{"Marmoset", "chr", 1, 22},
		{"Medaka", "chr", 1, 24},
		{"Mouse", "chr", 1, 19},
		{"NileTilapia", "chrLG", 1, 7},
		{"NileTilapia", "chrLG", 9, 15},
		{"NileTilapia", "chrLG", 17, 20},
		{"NileTilapia", "chrLG", 22, 23},
		{"Opossum", "chr", 1, 8},
		{"Orangutan", "chr", 3, 22},
		{"Pig", "chr", 1, 18},
		{"Platypus", "chr", 1, 7},
		{"Platypus", "chr", 10, 12},
		{"Platypus", "chr", 14, 15},
		{"Platypus", "chr", 17, 18},
		{"Platypus", "chrX", 1, 3},
		{"Rabbit", "chr", 1, 21},
		{"Rhesus", "chr", 1, 20},
		{"Sheep", "chr", 1, 26},
		{"Tetraodon", "chr", 1, 21},
		{"ZebraFinch", "chr", 1, 28},
		{"Zebrafish", "chr", 1, 25},
	} {
		unmaskChrRange(gen.genomeName, gen.prefix, "fa", gen.startIdentifier, gen.endIdentifier)
	}

	// exception of the above array (extension is fasta, not fa)
	unmaskChrRange("Elephant", "scaffold_", "fasta", 0, 2351)

	// array for cases using unmaskChr
	for _, gen := range []struct {
		genomeName  string
		prefix      string
		identifiers []string
	}{
		{"A_gambaie", "chr", []string{"2L", "2R", "3L", "3R", "X", "U"}},
		{"A_mellifera", "Group", []string{"Un"}},
		{"Baboon", "chr", []string{"X", "Un"}},
		{"Btau6", "Chr", []string{"X", "Un"}},
		{"Cat", "chr", []string{"X", "Un"}},
		{"Chicken", "chr", []string{"32", "W", "Z", "Un"}},
		{"Chimp", "chr", []string{"1", "2A", "2B", "X", "Y", "Un"}},
		{"D_melanogaster", "chr", []string{"2L", "2R", "3L", "3R", "4", "X", "U", "Uextra"}},
		{"Dog", "chr", []string{"X", "Un"}},
		{"D_simulans", "chr", []string{"2L", "2R", "3L", "3R", "4", "X", "U"}},
		{"D_yakuba", "chr", []string{"2L", "2R", "3L", "3R", "4", "X", "U"}},
		{"equCab2", "chr", []string{"X", "Un"}},
		{"Gibbon", "chr", []string{"1a", "7b", "22a", "X", "Un"}},
		{"Gorilla", "chr", []string{"1", "2A", "2B", "X", "Un"}},
		{"Human", "chr", []string{"X", "Y", "Un"}},
		{"Lizard", "chr", []string{"Un"}},
		{"Marmoset", "chr", []string{"X", "Y", "Un"}},
		{"Mouse", "chr", []string{"X", "Y", "Un"}},
		{"NileTilapia", "chrLG", []string{"8-24", "16-21", "Un"}},
		{"Opossum", "chr", []string{"X", "Un"}},
		{"Orangutan", "chr", []string{"1", "2a", "2b", "X", "Un"}},
		{"Pig", "chr", []string{"X", "Y", "Un"}},
		{"Platypus", "chr", []string{"20", "X5", "Un"}},
		{"Rabbit", "chr", []string{"X", "Un"}},
		{"Rhesus", "chr", []string{"X", "Un"}},
		{"SeaSquirt", "chr", []string{"01p", "01q", "02q", "03p", "03q", "04q", "05q", "06q", "07q", "08q", "09p", "09q", "10p", "10q", "12p", "12q", "13p", "13q", "14p", "14q"}},
		{"Sheep", "chr", []string{"X", "Un"}},
		{"Stickleback", "chr", []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX", "XXI", "Un"}},
		{"TasDevil", "chr", []string{"1.unlocalized.scaf", "2.unlocalized.scaf", "3.unlocalized.scaf", "4.unlocalized.scaf", "5.unlocalized.scaf", "6.unlocalized.scaf", "X.unlocalized.scaf", "Un"}},
		{"Tetraodon", "chr", []string{"Un"}},
		{"ZebraFinch", "chr", []string{"1A", "1B", "4A", "Z", "Un"}},
		{"Zebrafish", "chr", []string{"Un"}},
	} {
		for i := 0; i < len(gen.identifiers); i++ {
			unmaskChr(gen.genomeName, gen.prefix, "fa", gen.identifiers[i])
		}
	}

	// array for genomes with a single genome file (using unmaskChr)
	for _, gen := range []struct {
		genomeName string
		identifier string
	}{
		{"Alpaca", "vicPac2"},
		{"AmericanAlligator", "allMis1"},
		{"Armadillo", "dasNov3"},
		{"AtlanticCod", "gadMor1"},
		{"Budgerigar", "melUnd1"},
		{"Bushbaby", "otoGar3"},
		{"Coelacanth", "latCha1"},
		{"D_ananassae", "droAna3"},
		{"D_erecta", "droEre2"},
		{"D_grimshawi", "droGri2"},
		{"D_mojavensis", "droMoj3"},
		{"Dolphin", "turTru2"},
		{"D_persimilis", "droPer1"},
		{"D_pseudoobscura", "dp4"},
		{"D_sechellia", "droSec1"},
		{"D_virilis", "droVir3"},
		{"Ferret", "musFur1"},
		{"Frog", "xenTro3"},
		{"Fugu", "fr3"},
		{"GuineaPig", "cavPor3"},
		{"Hedgehog", "eriEur1"},
		{"KangarooRat", "dipOrd1"},
		{"Lamprey", "petMar2"},
		{"Manatee", "triMan1"},
		{"MediumGroundfinch", "geoFor1"},
		{"Megabat", "pteVam1"},
		{"Microbat", "myoLuc2"},
		{"MouseLemur", "micMur1"},
		{"NakedMolerat", "hetGla2"},
		{"PaintedTurtle", "chrPic1"},
		{"Panda", "ailMel1"},
		{"Pika", "ochPri2"},
		{"RockHyrax", "proCap1"},
		{"SeaHare", "aplCal1"},
		{"SeaUrchin", "strPur2"},
		{"Shrew", "sorAra1"},
		{"Sloth", "choHof1"},
		{"Squirrel", "speTri2"},
		{"SquirrelMonkey", "saiBol1"},
		{"Tarsier", "tarSyr1"},
		{"Tenrec", "echTel2"},
		{"TreeShrew", "tupBel1"},
		{"Turkey", "melGal1"},
		{"Wallaby", "macEug2"},
		{"WhiteRhino", "cerSim1"},
	} {
		unmaskChr(gen.genomeName, "", "fa", gen.identifier)
	}
}
