package main

import (
	"fmt"
	"os"
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

func main() {

	// array for genomes with a single genome file (using unmaskChr)
	for _, gen := range []struct {
		genomeName string
		identifier string
	}{
		{"Alpaca", "vicPac2"},
		{"AmericanAlligator", "allMis1"},
		{"Armadillo", "dasNov3"},
		{"AtlanticCod", "gadMor1"},
		{"Baboon", "papAnu2"},
		{"Budgerigar", "melUnd1"},
		{"Bushbaby", "otoGar3"},
		{"Cat", "felCat5"},
		{"Chicken", "galGal4"},
		{"Chimp", "panTro4"},
		{"Coelacanth", "latCha1"},
		{"Cow", "bosTau6"},
		{"D_ananassae", "droAna3"},
		{"D_erecta", "droEre2"},
		{"D_grimshawi", "droGri2"},
		{"D_mojavensis", "droMoj3"},
		{"Dog", "canFam3"},
		{"Dolphin", "turTru2"},
		{"D_persimilis", "droPer1"},
		{"D_pseudoobscura", "dp4"},
		{"D_sechellia", "droSec1"},
		{"D_virilis", "droVir3"},
		{"Elephant", "loxAfr3"},
		{"Ferret", "musFur1"},
		{"Frog", "xenTro3"},
		{"Fugu", "fr3"},
		{"Gibbon", "nomLeu3"},
		{"Gorilla", "gorGor3"},
		{"GuineaPig", "cavPor3"},
		{"Hedgehog", "eriEur1"},
		{"KangarooRat", "dipOrd1"},
		{"Lamprey", "petMar2"},
		{"Lizard", "anoCar2"},
		{"Manatee", "triMan1"},
		{"Marmoset", "calJac3"},
		{"Medaka", "oryLat2"},
		{"MediumGroundfinch", "geoFor1"},
		{"Megabat", "pteVam1"},
		{"Microbat", "myoLuc2"},
		{"MouseLemur", "micMur1"},
		{"NakedMolerat", "hetGla2"},
		{"NileTilapia", "oreNil2"},
		{"PaintedTurtle", "chrPic1"},
		{"Panda", "ailMel1"},
		{"Pig", "susScr3"},
		{"Pika", "ochPri2"},
		{"Platypus", "ornAna1"},
		{"Rabbit", "oryCun2"},
		{"Rat", "rn5"},
		{"Rhesus", "rheMac3"},
		{"RockHyrax", "proCap1"},
		{"SeaHare", "aplCal1"},
		{"SeaUrchin", "strPur2"},
		{"Sheep", "oviAri3"},
		{"Shrew", "sorAra1"},
		{"Sloth", "choHof1"},
		{"Squirrel", "speTri2"},
		{"SquirrelMonkey", "saiBol1"},
		{"Stickleback", "gasAcu1"},
		{"Tarsier", "tarSyr1"},
		{"TasDevil", "sarHar1"},
		{"Tenrec", "echTel2"},
		{"TreeShrew", "tupBel1"},
		{"Turkey", "melGal1"},
		{"Wallaby", "macEug2"},
		{"WhiteRhino", "cerSim1"},
		{"ZebraFinch", "taeGut1"},
		{"Zebrafish", "danRer7"},
	} {
		unmaskChr(gen.genomeName, "", "fa", gen.identifier)
	}
}
