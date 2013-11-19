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

	for _, gen := range []struct {
		genomeName  string
		prefix      string
		identifiers []string
	}{
		{"A_gambaie", "chr", []string{"2L", "2R", "3L", "3R", "X", "U"}},
	} {
		for i := 1; i <= len(gen.identifiers); i++ {
			unmaskChr(gen.genomeName, gen.prefix, "fa", gen.identifiers[i])
		}
	}

	// for _, gen := range []struct {
	// 	genomeName string
	// 	identifier string
	// }{
	// 	{"Alpaca", "vicPac2"},
	// 	{"AmericanAlligator", "allMis1"},
	// 	{"Armadillo", "dasNov3"},
	// 	{"AtlanticCod", "gadMor1"},
	// 	{"Budgerigar", "melUnd1"},
	// 	{"Bushbaby", "otoGar3"},
	// 	{"Coelacanth", "latCha1"},
	// 	{"D_ananassae", "droAna3"},
	// 	{"D_erecta", "droEre2"},
	// 	{"D_grimshawi", "droGri2"},
	// 	{"D_mojavensis", "droMoj3"},
	// 	{"Dolphin", "turTru2"},
	// 	{"D_persimilis", "droPer1"},
	// 	{"D_pseudoobscura", "dp4"},
	// 	{"D_sechellia", "droSec1"},
	// 	{"D_virilis", "droVir3"},
	// 	{"Ferret", "musFur1"},
	// 	{"Frog", "xenTro3"},
	// 	{"Fugu", "fr3"},
	// 	{"GuineaPig", "cavPor3"},
	// 	{"Hedgehog", "eriEur1"},
	// 	{"KangarooRat", "dipOrd1"},
	// 	{"Lamprey", "petMar2"},
	// 	{"Manatee", "triMan1"},
	// 	{"MediumGroundfinch", "geoFor1"},
	// 	{"Megabat", "pteVam1"},
	// 	{"Microbat", "myoLuc2"},
	// 	{"MouseLemur", "micMur1"},
	// 	{"NakedMolerat", "hetGla2"},
	// 	{"PaintedTurtle", "chrPic1"},
	// 	{"Panda", "ailMel1"},
	// 	{"Pika", "ochPri2"},
	// 	{"RockHyrax", "proCap1"},
	// 	{"SeaHare", "aplCal1"},
	// 	{"SeaUrchin", "strPur2"},
	// 	{"Shrew", "sorAra1"},
	// 	{"Sloth", "choHof1"},
	// 	{"Squirrel", "speTri2"},
	// 	{"SquirrelMonkey", "saiBol1"},
	// 	{"Tarsier", "tarSyr1"},
	// 	{"Tenrec", "echTel2"},
	// 	{"TreeShrew", "tupBel1"},
	// 	{"Turkey", "melGal1"},
	// 	{"Wallaby", "macEug2"},
	// 	{"WhiteRhino", "cerSim1"},
	// } {
	// 	unmaskChr(gen.genomeName, "", "fa", gen.identifier)
	// }

	// unmaskChr("A_gambaie", "chr", "fa", "2L")
	// unmaskChr("A_gambaie", "chr", "fa", "2R")
	// unmaskChr("A_gambaie", "chr", "fa", "3L")
	// unmaskChr("A_gambaie", "chr", "fa", "3R")
	// unmaskChr("A_gambaie", "chr", "fa", "X")
	// unmaskChr("A_gambaie", "chr", "fa", "U")

	// unmaskChrRange("A_mellifera", "Group", "fa", 1, 16)
	// unmaskChr("A_mellifera", "Group", "fa", "Un")

	// unmaskChrRange("Baboon", "chr", "fa", 1, 20)
	// unmaskChr("Baboon", "chr", "fa", "X")
	// unmaskChr("Baboon", "chr", "fa", "Un")

	// unmaskChrRange("Btau6", "Chr", "fa", 1, 29)
	// unmaskChr("Btau6", "Chr", "fa", "X")
	// unmaskChr("Btau6", "Chr", "fa", "U")

	// unmaskChrRange("Cat", "chrA", "fa", 1, 3)
	// unmaskChrRange("Cat", "chrB", "fa", 1, 4)
	// unmaskChrRange("Cat", "chrC", "fa", 1, 2)
	// unmaskChrRange("Cat", "chrD", "fa", 1, 4)
	// unmaskChrRange("Cat", "chrE", "fa", 1, 3)
	// unmaskChrRange("Cat", "chrF", "fa", 1, 2)
	// unmaskChr("Cat", "chr", "fa", "X")
	// unmaskChr("Cat", "chr", "fa", "Un")

	// unmaskChrRange("Chicken", "chr", "fa", 1, 28)
	// unmaskChr("Chicken", "chr", "fa", "32")
	// unmaskChr("Chicken", "chr", "fa", "W")
	// unmaskChr("Chicken", "chr", "fa", "Z")
	// unmaskChr("Chicken", "chr", "fa", "Un")

	// unmaskChr("Chimp", "chr", "fa", "1")
	// unmaskChr("Chimp", "chr", "fa", "2A")
	// unmaskChr("Chimp", "chr", "fa", "2B")
	// unmaskChrRange("Chimp", "chr", "fa", 3, 22)
	// unmaskChr("Chimp", "chr", "fa", "X")
	// unmaskChr("Chimp", "chr", "fa", "Y")
	// unmaskChr("Chimp", "chr", "fa", "Un")

	// unmaskChr("D_melanogaster", "chr", "fa", "2L")
	// unmaskChr("D_melanogaster", "chr", "fa", "2R")
	// unmaskChr("D_melanogaster", "chr", "fa", "3L")
	// unmaskChr("D_melanogaster", "chr", "fa", "3R")
	// unmaskChr("D_melanogaster", "chr", "fa", "4")
	// unmaskChr("D_melanogaster", "chr", "fa", "X")
	// unmaskChr("D_melanogaster", "chr", "fa", "U")
	// unmaskChr("D_melanogaster", "chr", "fa", "Uextra")

	// unmaskChrRange("Dog", "chr", "fa", 1, 38)
	// unmaskChr("Dog", "chr", "fa", "X")
	// unmaskChr("Dog", "chr", "fa", "Un")

	// unmaskChr("D_simulans", "chr", "fa", "2L")
	// unmaskChr("D_simulans", "chr", "fa", "2R")
	// unmaskChr("D_simulans", "chr", "fa", "3L")
	// unmaskChr("D_simulans", "chr", "fa", "3R")
	// unmaskChr("D_simulans", "chr", "fa", "4")
	// unmaskChr("D_simulans", "chr", "fa", "X")
	// unmaskChr("D_simulans", "chr", "fa", "U")

	// unmaskChr("D_yakuba", "chr", "fa", "2L")
	// unmaskChr("D_yakuba", "chr", "fa", "2R")
	// unmaskChr("D_yakuba", "chr", "fa", "3L")
	// unmaskChr("D_yakuba", "chr", "fa", "3R")
	// unmaskChr("D_yakuba", "chr", "fa", "4")
	// unmaskChr("D_yakuba", "chr", "fa", "X")
	// unmaskChr("D_yakuba", "chr", "fa", "U")

	// unmaskChrRange("Elephant", "scaffold_", "fasta", 0, 2351)

	// unmaskChrRange("equCab2", "chr", "fa", 1, 31)
	// unmaskChr("equCab2", "chr", "fa", "X")
	// unmaskChr("equCab2", "chr", "fa", "Un")

	// unmaskChr("Gibbon", "chr", "fa", "1a")
	// unmaskChrRange("Gibbon", "chr", "fa", 2, 6)
	// unmaskChr("Gibbon", "chr", "fa", "7b")
	// unmaskChrRange("Gibbon", "chr", "fa", 8, 21)
	// unmaskChr("Gibbon", "chr", "fa", "22a")
	// unmaskChrRange("Gibbon", "chr", "fa", 23, 25)
	// unmaskChr("Gibbon", "chr", "fa", "X")
	// unmaskChr("Gibbon", "chr", "fa", "Un")

	// unmaskChr("Gorilla", "chr", "fa", "1")
	// unmaskChr("Gorilla", "chr", "fa", "2A")
	// unmaskChr("Gorilla", "chr", "fa", "2B")
	// unmaskChrRange("Gorilla", "chr", "fa", 3, 22)
	// unmaskChr("Gorilla", "chr", "fa", "X")
	// unmaskChr("Gorilla", "chr", "fa", "Un")

	// unmaskChrRange("Human", "chr", "fa", 1, 22)
	// unmaskChr("Human", "chr", "fa", "X")
	// unmaskChr("Human", "chr", "fa", "Y")
	// unmaskChr("Human", "chr", "fa", "Un")

	// unmaskChrRange("Lizard", "chr", "fa", 1, 6)
	// unmaskChr("Lizard", "chr", "fa", "Un")

	// unmaskChrRange("Marmoset", "chr", "fa", 1, 22)
	// unmaskChr("Marmoset", "chr", "fa", "X")
	// unmaskChr("Marmoset", "chr", "fa", "Y")
	// unmaskChr("Marmoset", "chr", "fa", "Un")

	// unmaskChrRange("Medaka", "chr", "fa", 1, 24)

	// unmaskChrRange("Mouse", "chr", "fa", 1, 19)
	// unmaskChr("Mouse", "chr", "fa", "X")
	// unmaskChr("Mouse", "chr", "fa", "Y")
	// unmaskChr("Mouse", "chr", "fa", "Un")

	// unmaskChrRange("NileTilapia", "chrLG", "fa", 1, 7)
	// unmaskChr("NileTilapia", "chrLG", "fa", "8-24")
	// unmaskChrRange("NileTilapia", "chrLG", "fa", 9, 15)
	// unmaskChr("NileTilapia", "chrLG", "fa", "16-21")
	// unmaskChrRange("NileTilapia", "chrLG", "fa", 17, 20)
	// unmaskChrRange("NileTilapia", "chrLG", "fa", 22, 23)
	// unmaskChr("NileTilapia", "chr", "fa", "Un")

	// unmaskChrRange("Opossum", "chr", "fa", 1, 8)
	// unmaskChr("Opossum", "chr", "fa", "X")
	// unmaskChr("Opossum", "chr", "fa", "Un")

	// unmaskChr("Orangutan", "chr", "fa", "1")
	// unmaskChr("Orangutan", "chr", "fa", "2a")
	// unmaskChr("Orangutan", "chr", "fa", "2b")
	// unmaskChrRange("Orangutan", "chr", "fa", 3, 22)
	// unmaskChr("Orangutan", "chr", "fa", "X")
	// unmaskChr("Orangutan", "chr", "fa", "Un")

	// unmaskChrRange("Pig", "chr", "fa", 1, 18)
	// unmaskChr("Pig", "chr", "fa", "X")
	// unmaskChr("Pig", "chr", "fa", "Y")
	// unmaskChr("Pig", "chr", "fa", "Un")

	// unmaskChrRange("Platypus", "chr", "fa", 1, 7)
	// unmaskChrRange("Platypus", "chr", "fa", 10, 12)
	// unmaskChrRange("Platypus", "chr", "fa", 14, 15)
	// unmaskChrRange("Platypus", "chr", "fa", 17, 18)
	// unmaskChr("Platypus", "chr", "fa", "20")
	// unmaskChrRange("Platypus", "chrX", "fa", 1, 3)
	// unmaskChr("Platypus", "chr", "fa", "X5")
	// unmaskChr("Platypus", "chr", "fa", "Un")

	// unmaskChrRange("Rabbit", "chr", "fa", 1, 21)
	// unmaskChr("Rabbit", "chr", "fa", "X")
	// unmaskChr("Rabbit", "chr", "fa", "Un")

	// unmaskChrRange("Rhesus", "chr", "fa", 1, 20)
	// unmaskChr("Rhesus", "chr", "fa", "X")
	// unmaskChr("Rhesus", "chr", "fa", "Un")

	// unmaskChr("SeaSquirt", "chr", "fa", "01p")
	// unmaskChr("SeaSquirt", "chr", "fa", "01q")
	// unmaskChr("SeaSquirt", "chr", "fa", "02q")
	// unmaskChr("SeaSquirt", "chr", "fa", "03p")
	// unmaskChr("SeaSquirt", "chr", "fa", "03q")
	// unmaskChr("SeaSquirt", "chr", "fa", "04q")
	// unmaskChr("SeaSquirt", "chr", "fa", "05q")
	// unmaskChr("SeaSquirt", "chr", "fa", "06q")
	// unmaskChr("SeaSquirt", "chr", "fa", "07q")
	// unmaskChr("SeaSquirt", "chr", "fa", "08q")
	// unmaskChr("SeaSquirt", "chr", "fa", "09p")
	// unmaskChr("SeaSquirt", "chr", "fa", "09q")
	// unmaskChr("SeaSquirt", "chr", "fa", "10p")
	// unmaskChr("SeaSquirt", "chr", "fa", "10q")
	// unmaskChr("SeaSquirt", "chr", "fa", "12p")
	// unmaskChr("SeaSquirt", "chr", "fa", "12q")
	// unmaskChr("SeaSquirt", "chr", "fa", "13p")
	// unmaskChr("SeaSquirt", "chr", "fa", "13q")
	// unmaskChr("SeaSquirt", "chr", "fa", "14p")
	// unmaskChr("SeaSquirt", "chr", "fa", "14q")

	// unmaskChrRange("Sheep", "chr", "fa", 1, 26)
	// unmaskChr("Sheep", "chr", "fa", "X")
	// unmaskChr("Sheep", "chr", "fa", "Un")

	// unmaskChr("Stickleback", "chr", "fa", "I")
	// unmaskChr("Stickleback", "chr", "fa", "II")
	// unmaskChr("Stickleback", "chr", "fa", "III")
	// unmaskChr("Stickleback", "chr", "fa", "IV")
	// unmaskChr("Stickleback", "chr", "fa", "V")
	// unmaskChr("Stickleback", "chr", "fa", "VI")
	// unmaskChr("Stickleback", "chr", "fa", "VII")
	// unmaskChr("Stickleback", "chr", "fa", "VIII")
	// unmaskChr("Stickleback", "chr", "fa", "IX")
	// unmaskChr("Stickleback", "chr", "fa", "X")
	// unmaskChr("Stickleback", "chr", "fa", "XI")
	// unmaskChr("Stickleback", "chr", "fa", "XII")
	// unmaskChr("Stickleback", "chr", "fa", "XIII")
	// unmaskChr("Stickleback", "chr", "fa", "XIV")
	// unmaskChr("Stickleback", "chr", "fa", "XV")
	// unmaskChr("Stickleback", "chr", "fa", "XVI")
	// unmaskChr("Stickleback", "chr", "fa", "XVII")
	// unmaskChr("Stickleback", "chr", "fa", "XVIII")
	// unmaskChr("Stickleback", "chr", "fa", "XIX")
	// unmaskChr("Stickleback", "chr", "fa", "XX")
	// unmaskChr("Stickleback", "chr", "fa", "XXI")
	// unmaskChr("Stickleback", "chr", "fa", "Un")

	// unmaskChr("TasDevil", "chr", "fa", "1.unlocalized.scaf")
	// unmaskChr("TasDevil", "chr", "fa", "2.unlocalized.scaf")
	// unmaskChr("TasDevil", "chr", "fa", "3.unlocalized.scaf")
	// unmaskChr("TasDevil", "chr", "fa", "4.unlocalized.scaf")
	// unmaskChr("TasDevil", "chr", "fa", "5.unlocalized.scaf")
	// unmaskChr("TasDevil", "chr", "fa", "6.unlocalized.scaf")
	// unmaskChr("TasDevil", "chr", "fa", "X.unlocalized.scaf")
	// unmaskChr("TasDevil", "chr", "fa", "Un")

	// unmaskChrRange("Tetraodon", "chr", "fa", 1, 21)
	// unmaskChr("Tetraodon", "chr", "fa", "Un")

	// unmaskChrRange("ZebraFinch", "chr", "fa", 1, 28)
	// unmaskChr("ZebraFinch", "chr", "fa", "1A")
	// unmaskChr("ZebraFinch", "chr", "fa", "1B")
	// unmaskChr("ZebraFinch", "chr", "fa", "4A")
	// unmaskChr("ZebraFinch", "chr", "fa", "Z")
	// unmaskChr("ZebraFinch", "chr", "fa", "Un")

	// unmaskChrRange("Zebrafish", "chr", "fa", 1, 25)
	// unmaskChr("zebrafish", "chr", "fa", "Un")

	// // unmaskChr("Alpaca", "", "fa", "vicPac2")

	// // unmaskChr("AmericanAlligator", "", "fa", "allMis1")

	// // unmaskChr("Armadillo", "", "fa", "dasNov3")

	// // unmaskChr("AtlanticCod", "", "fa", "gadMor1")

	// // unmaskChr("Budgerigar", "", "fa", "melUnd1")

	// // unmaskChr("Bushbaby", "", "fa", "otoGar3")

	// // unmaskChr("Coelacanth", "", "fa", "latCha1")

	// // unmaskChr("D_ananassae", "", "fa", "droAna3")

	// // unmaskChr("D_erecta", "", "fa", "droEre2")

	// // unmaskChr("D_grimshawi", "", "fa", "droGri2")

	// // unmaskChr("D_mojavensis", "", "fa", "droMoj3")

	// unmaskChr("Dolphin", "", "fa", "turTru2")

	// unmaskChr("D_persimilis", "", "fa", "droPer1")

	// unmaskChr("D_pseudoobscura", "", "fa", "dp4")

	// unmaskChr("D_sechellia", "", "fa", "droSec1")

	// unmaskChr("D_virilis", "", "fa", "droVir3")

	// unmaskChr("Ferret", "", "fa", "musFur1")

	// unmaskChr("Frog", "", "fa", "xenTro3")

	// unmaskChr("Fugu", "", "fa", "fr3")

	// unmaskChr("GuineaPig", "", "fa", "cavPor3")

	// unmaskChr("Hedgehog", "", "fa", "eriEur1")

	// unmaskChr("KangarooRat", "", "fa", "dipOrd1")

	// unmaskChr("Lamprey", "", "fa", "petMar2")

	// unmaskChr("Manatee", "", "fa", "triMan1")

	// unmaskChr("MediumGroundfinch", "", "fa", "geoFor1")

	// unmaskChr("Megabat", "", "fa", "pteVam1")

	// unmaskChr("Microbat", "", "fa", "myoLuc2")

	// unmaskChr("MouseLemur", "", "fa", "micMur1")

	// unmaskChr("NakedMolerat", "", "fa", "hetGla2")

	// unmaskChr("PaintedTurtle", "", "fa", "chrPic1")

	// unmaskChr("Panda", "", "fa", "ailMel1")

	// unmaskChr("Pika", "", "fa", "ochPri2")

	// unmaskChr("RockHyrax", "", "fa", "proCap1")

	// unmaskChr("SeaHare", "", "fa", "aplCal1")

	// unmaskChr("SeaUrchin", "", "fa", "strPur2")

	// unmaskChr("Shrew", "", "fa", "sorAra1")

	// unmaskChr("Sloth", "", "fa", "choHof1")

	// unmaskChr("Squirrel", "", "fa", "speTri2")

	// unmaskChr("SquirrelMonkey", "", "fa", "saiBol1")

	// unmaskChr("Tarsier", "", "fa", "tarSyr1")

	// unmaskChr("Tenrec", "", "fa", "echTel2")

	// unmaskChr("TreeShrew", "", "fa", "tupBel1")

	// unmaskChr("Turkey", "", "fa", "melGal1")

	// unmaskChr("Wallaby", "", "fa", "macEug2")

	// unmaskChr("WhiteRhino", "", "fa", "cerSim1")

}
