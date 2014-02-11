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
		// {"Alpaca", "vicPac2"},
		// {"AmericanAlligator", "allMis1"},
		// {"Armadillo", "dasNov3"},
		// {"AtlanticCod", "gadMor1"},
		// {"Baboon", "papAnu2"},
		// {"Budgerigar", "melUnd1"},
		// {"Bushbaby", "otoGar3"},
		// {"Cat", "felCat5"},
		// {"Chicken", "galGal4"},
		// {"Chimp", "panTro4"},
		// {"Coelacanth", "latCha1"},
		// {"Cow", "bosTau6"},
		// {"D_ananassae", "droAna3"},
		// {"D_erecta", "droEre2"},
		// {"D_grimshawi", "droGri2"},
		// {"D_mojavensis", "droMoj3"},
		// {"Dog", "canFam3"},
		// {"Dolphin", "turTru2"},
		// {"D_persimilis", "droPer1"},
		// {"D_pseudoobscura", "dp4"},
		// {"D_sechellia", "droSec1"},
		// {"D_virilis", "droVir3"},
		// {"Elephant", "loxAfr3"},
		// {"Ferret", "musFur1"},
		// {"Frog", "xenTro3"},
		// {"Fugu", "fr3"},
		// {"Gibbon", "nomLeu3"},
		// {"Gorilla", "gorGor3"},
		// {"GuineaPig", "cavPor3"},
		// {"Hedgehog", "eriEur1"},
		// {"KangarooRat", "dipOrd1"},
		// {"Lamprey", "petMar2"},
		// {"Lizard", "anoCar2"},
		// {"Manatee", "triMan1"},
		// {"Marmoset", "calJac3"},
		// {"Medaka", "oryLat2"},
		// {"MediumGroundfinch", "geoFor1"},
		// {"Megabat", "pteVam1"},
		// {"Microbat", "myoLuc2"},
		// {"MouseLemur", "micMur1"},
		// {"NakedMolerat", "hetGla2"},
		// {"NileTilapia", "oreNil2"},
		// {"PaintedTurtle", "chrPic1"},
		// {"Panda", "ailMel1"},
		// {"Pig", "susScr3"},
		// {"Pika", "ochPri2"},
		// {"Platypus", "ornAna1"},
		// {"Rabbit", "oryCun2"},
		// {"Rat", "rn5"},
		// {"Rhesus", "rheMac3"},
		// {"RockHyrax", "proCap1"},
		// {"SeaHare", "aplCal1"},
		// {"SeaUrchin", "strPur2"},
		// {"Sheep", "oviAri3"},
		// {"Shrew", "sorAra1"},
		// {"Sloth", "choHof1"},
		// {"Squirrel", "speTri2"},
		// {"SquirrelMonkey", "saiBol1"},
		// {"Stickleback", "gasAcu1"},
		// {"Tarsier", "tarSyr1"},
		// {"TasDevil", "sarHar1"},
		// {"Tenrec", "echTel2"},
		// {"TreeShrew", "tupBel1"},
		// {"Turkey", "melGal1"},
		// {"Wallaby", "macEug2"},
		// {"WhiteRhino", "cerSim1"},
		// {"ZebraFinch", "taeGut1"},
		// {"Zebrafish", "danRer7"},
		// {"CapeElephantShrew", "EleEdw1.0"},
		// {"Aardvark", "OryAfe1.0"},
		// {"CapeGoldenMole", "ChrAsi1.0"},
		// {"AmericanPika", "OchPri3.0"},
		// {"Mouse", "GRCm38.p2"},
		// {"GoldenHamster", "MesAur1.0"},
		// {"PrairieVole", "MicOch1.0"},
		// {"LesserEgyptianJerboa", "JacJac1.0"},
		// {"Chinchilla", "ChiLan1.0"},
		// {"Rat", "Rnor_5.0"},
		// {"Degu", "OctDeg1.0"},
		// {"ChineseHamster", "CriGri_1.0"},
		// {"DeerMouse", "Pman_1.0"},
		// {"GreenMonkey", "Chlorocebus_sabeus_1.0"},
		// {"CrabeatingMacaque", "Macaca_fascicularis_5.0"},
		// {"Bonobo", "panpan1"},
		// {"SumatranOrangutan", "P_pygmaeus_2.0.2"},
		// {"Human", "GRCh38"},
		// {"Tarsier", "Tarsius_syrichta_2.0.1"},
		// {"Marmoset", "Callithrix_jacchus_3.2"},
		// {"ChineseTreeShrew", "TupChi_1.0"},
		// {"WeddellSeal", "LepWed1.0"},
		// {"Walrus", "Oros_1.0"},
		// {"SiberianTiger", "PanTig1.0"},
		// {"TibetanAntelope", "PHO1.0"},
		// {"KillerWhale", "Oorc_1.1"},
		// {"Goat", "CHIR_1.0"},
		// {"Yak", "BosGru_v2.0"},
		// {"MinkeWhale", "BalAcu1.0"},
		// {"WaterBuffalo", "UMD_CASPUR_WB_2.0"},
		// {"SpermWhale", "Physeter_macrocephalus_2.0.2"},
		// {"Baiji", "Lipotes_vexillifer_v1"},
		// {"BactrianCamel", "CB1"},
		// {"MouseearedBat", "ASM32734v1"},
		// {"BlackFlyingFox", "ASM32557v1"},
		// {"BigBrownBat", "EptFus1.0"},
		// {"StrawcolouredFruitBat", "ASM46528v1"},
		// {"GreaterFalseVampireBat", "ASM46534v1"},
		// {"ParnellsMustachedBat", "ASM46540v1"},
		// {"Hedgehog", "EriEur2.0"},
		// {"CommonShrew", "SorAra2.0"},
		// {"StarnosedMole", "ConCri1.0"},
		// {"BurmesePython", "Python_molurus_bivittatus_5.0.2"},
		// {"SpinySoftshellTurtle", "ASM38561v1"},
		// {"GreenSeaTurtle", "CheMyd_1.0"},
		// {"ChineseSoftshellTurtle", "PelSin_1.0"},
		// {"ChineseAlligator", "ASM45574v1"},
		// {"ScarletMacaw", "SMACv1.1"},
		// {"WhitethroatedSparrow", "Zonotrichia_albicollis_1.0.1"},
		// {"Mallard", "BGI_duck_1.0"},
		// {"RockDove", "Cliv_1.0"},
		// {"SakerFalcon", "F_cherrug_v1.0"},
		// {"PeregrineFalcon", "F_peregrinus_v1.0"},
		// {"PuertoRicanParrot", "AV1"},
		// {"GroundTit", "PseHum1.0"},
		// {"AtlanticCanary", "SCA1"},
		// {"JapaneseQuail", "Coja_1.0"},
		// {"CollaredFlycatcher", "FicAlb1.5"},
		// {"WesternClawedFrog", "Xtropicalis_v7"},
		// {"AustralianGhostshark", "Callorhinchus_milii_6.1.3"},
		// {"YellowbellyPufferfish", "version_1_of_Takifugu_flavidus_genome"},
		// {"MexicanTetra", "Astyanax_mexicanus_1.0.2"},
		// {"Medaka", "ASM31367v1"},
		// {"ZebraMbuna", "MetZeb1.1"},
		// {"SpottedGar", "LepOcu1"},
		// {"SouthernPlatyfish", "Xiphophorus_maculatus_4.4.2"},
		// {"AstatotilapiaBurtoni", "AstBur1.0"},
		// {"NeolamprologusBrichardi", "NeoBri1.0"},
		// {"FlameBackCichlid", "PunNye1.0"},
		// {"ArticLamprey", "LetJap1.0"},
		{"GreaterHorseshoeBat", "ASM46549v1"},
		{"BrandtsBat", "ASM41265v1"},
	} {
		unmaskChr(gen.genomeName, "", "fa", gen.identifier)
	}
}
