package main

import (
	"fmt"
	"os"
)

const genomeDir = "/data01/Genomes/Protostomia/"

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
		// {"Aardvark", "OryAfe1.0"}, //Vertebrates
		// {"Alpaca", "vicPac2"},
		// {"AmazonMolly", "Poecilia_formosa_5.1.2"},
		// {"AmericanAlligator", "allMis1"},
		// {"AmericanPika", "OchPri3.0"},
		// {"ArcticLamprey", "LetJap1.0"},
		// {"Armadillo", "dasNov3"},
		// {"AstatotilapiaBurtoni", "AstBur1.0"},
		// {"AtlanticCanary", "SCA1"},
		// {"AtlanticCod", "gadMor1"},
		// {"AustralianGhostshark", "Callorhinchus_milii_6.1.3"},
		// {"Baboon", "papAnu2"},
		// {"BactrianCamel", "CB1"},
		// {"Baiji", "Lipotes_vexillifer_v1"},
		// {"BigBrownBat", "EptFus1.0"},
		// {"BlackCormorant", "ASM70892v1"},
		// {"BlackFlyingFox", "ASM32557v1"},
		// {"BlindMolerat", "S.galili_v1.0"},
		// {"Bonobo", "panpan1"},
		// {"BrandtsBat", "ASM41265v1"},
		// {"Budgerigar", "melUnd1"},
		// {"BurmesePython", "Python_molurus_bivittatus_5.0.2"},
		// {"Bushbaby", "otoGar3"},
		// {"CapeElephantShrew", "EleEdw1.0"},
		// {"CapeGoldenMole", "ChrAsi1.0"},
		// {"Cat", "felCat5"},
		// {"Chicken", "galGal4"},
		// {"Chimp", "panTro4"},
		// {"Chinchilla", "ChiLan1.0"},
		// {"ChineseAlligator", "ASM45574v1"},
		// {"ChineseHamster", "CriGri_1.0"},
		// {"ChineseSoftshellTurtle", "PelSin_1.0"},
		// {"ChineseTreeShrew", "TupChi_1.0"},
		// {"Coelacanth", "latCha1"},
		// {"CollaredFlycatcher", "FicAlb1.5"},
		// {"CommonShrew", "SorAra2.0"},
		// {"Cow", "bosTau6"},
		// {"CrabeatingMacaque", "Macaca_fascicularis_5.0"},
		// {"C_semi", "Cse_v1.0"},
		// {"DeerMouse", "Pman_1.0"},
		// {"Degu", "OctDeg1.0"},
		// {"Dog", "canFam3"},
		// {"Dolphin", "turTru2"},
		// {"Echidna", "Tachyglossus"},
		// {"Elephant", "LAv4"},
		// {"ElephantShark", "calMil1"},
		// {"Ferret", "musFur1"},
		// {"FlagRockfish", "SRub1.0"},
		// {"FlameBackCichlid", "PunNye1.0"},
		// {"Fugu", "fr3"},
		// {"Gibbon", "nomLeu3"},
		// {"Goat", "CHIR_1.0"},
		// {"GoldenHamster", "MesAur1.0"},
		// {"Gorilla", "gorGor3"},
		// {"GreaterFalseVampireBat", "ASM46534v1"},
		// {"GreaterHorseshoeBat", "ASM46549v1"},
		// {"GreenMonkey", "Chlorocebus_sabeus_1.0"},
		// {"GreenSeaTurtle", "CheMyd_1.0"},
		// {"GroundTit", "PseHum1.0"},
		// {"GuineaPig", "cavPor3"},
		// {"Hedgehog", "EriEur2.0"},
		// {"Horse", "equCab2"},
		// {"Human", "GRCh38"},
		// {"JapaneseQuail", "Coja_1.0"},
		// {"KangarooRat", "dipOrd1"},
		// {"KillerWhale", "Oorc_1.1"},
		// {"Lamprey", "petMar2"},
		// {"LesserEgyptianJerboa", "JacJac1.0"},
		// {"Lizard", "anoCar2"},
		// {"Mallard", "BGI_duck_1.0"},
		// {"Manatee", "triMan1"},
		// {"Marmoset", "calJac3"},
		// {"Medaka", "ASM31367v1"},
		// {"MediumGroundfinch", "geoFor1"},
		// {"Megabat", "pteVam1"},
		// {"MexicanTetra", "Astyanax_mexicanus_1.0.2"},
		// {"Microbat", "myoLuc2"},
		// {"MinkeWhale", "BalAcu1.0"},
		// {"Mouse", "GRCm38.p2"},
		// {"MouseearedBat", "ASM32734v1"},
		// {"MouseLemur", "micMur1"},
		// {"Mummichog", "Fundulus_heteroclitus_3.0.2"},
		// {"NakedMolerat", "hetGla2"},
		// {"NeolamprologusBrichardi", "NeoBri1.0"},
		// {"NileTilapia", "oreNil2"},
		// {"Opossum", "monDom5"},
		// {"PaintedTurtle", "Chrysemys_picta_bellii_3.0.3"},
		// {"Panda", "ailMel1"},
		// {"ParnellsMustachedBat", "ASM46540v1"},
		// {"PeregrineFalcon", "F_peregrinus_v1.0"},
		// {"Pig", "susScr3"},
		// {"Platypus", "ornAna1"},
		// {"PrairieVole", "MicOch1.0"},
		// {"PuertoRicanParrot", "AV1"},
		// {"Rabbit", "oryCun2"},
		// {"Rat", "rn5"},
		// {"Rhesus", "rheMac3"},
		// {"RockDove", "Cliv_1.0"},
		// {"RockHyrax", "proCap1"},
		// {"SakerFalcon", "F_cherrug_v1.0"},
		// {"ScarletMacaw", "SMACv1.1"},
		// {"SeaHare", "aplCal1"},
		// {"Sheep", "oviAri3"},
		// {"SiberianTiger", "PanTig1.0"},
		// {"Sloth", "choHof1"},
		// {"SouthernPlatyfish", "Xiphophorus_maculatus_4.4.2"},
		// {"SpermWhale", "Physeter_macrocephalus_2.0.2"},
		// {"SpinySoftshellTurtle", "ASM38561v1"},
		// {"SpottedGar", "LepOcu1"},
		// {"Squirrel", "speTri2"},
		// {"SquirrelMonkey", "saiBol1"},
		// {"StarnosedMole", "ConCri1.0"},
		// {"Stickleback", "gasAcu1"},
		// {"StrawcolouredFruitBat", "ASM46528v1"},
		// {"SumatranOrangutan", "P_pygmaeus_2.0.2"},
		// {"Tarsier", "Tarsius_syrichta_2.0.1"},
		// {"TasDevil", "sarHar1"},
		// {"Tenrec", "echTel2"},
		// {"Tetraodon", "tetNig2"},
		// {"TibetanAntelope", "PHO1.0"},
		// {"TigerRockfish", "Snig1.0"},
		// {"TreeShrew", "tupBel1"},
		// {"Turkey", "melGal1"},
		// {"Wallaby", "macEug2"},
		// {"Walrus", "Oros_1.0"},
		// {"WaterBuffalo", "UMD_CASPUR_WB_2.0"},
		// {"WeddellSeal", "LepWed1.0"},
		// {"WesternClawedFrog", "Xtropicalis_v7"},
		// {"WhiteRhino", "cerSim1"},
		// {"WhitethroatedSparrow", "Zonotrichia_albicollis_1.0.1"},
		// {"Yak", "BosGru_v2.0"},
		// {"YellowbellyPufferfish", "version_1_of_Takifugu_flavidus_genome"},
		// {"ZebraFinch", "taeGut1"},
		// {"Zebrafish", "danRer7"},
		// {"ZebraMbuna", "MetZeb1.1"},
		// {"Zebu", "Bos_indicus_1.0"},
		// {"Egret", "ASM68718v1"}, //Avian
		// {"BarnOwl", "ASM68720v1"},
		// {"Flamingo", "ASM68726v1"},
		// {"Tropicbird", "ASM68728v1"},
		// {"Pelican", "ASM68737v1"},
		// {"Seriema", "ASM69053v1"},
		// {"Mousebird", "ASM690715v1"},
		// {"Sunbittern", "ASM69077v1"},
		// {"Fulmar", "ASM69083v1"},
		// {"Loon", "ASM69087v1"},
		// {"WhitetailedEagle", "ASM69140v1"},
		// {"CuckooRoller", "ASM69178v1"},
		// {"BeeEater", "ASM69184v1"},
		// {"AmericanCrow", "ASM69197v1"},
		// {"Manakin", "ASM69201v1"},
		// {"Hoatzin", "ASM69207v1"},
		// {"MacQueensBustard", "ASM69519v1"},
		// {"Mesite", "ASM69576v1"},
		// {"Rifleman", "ASM69581v1"},
		// {"Kea", "ASM69687v1"},
		// {"Ostrich", "ASM69896v1"},
		// {"Woodpecker", "ASM69900v1"},
		// {"AnnasHummingbird", "ASM69908v1"},
		// {"AdeliePenguin", "ASM69910v1"},
		// {"EmperorPenguin", "ASM69914v1"},
		// {"Sandgrouse", "ASM69924v1"},
		// {"Grebe", "ASM69954v1"},
		// {"TurkeyVulture", "ASM69994v1"},
		// {"ChuckWillsWidow", "ASM70074v1"},
		// {"Trogon", "ASM70340v1"},
		// {"Tinamou", "ASM70537v2"},
		// {"Killdeer", "ASM70802v2"},
		// {"Ibis", "ASM70822v1"},
		// {"CommonCuckoo", "ASM70932v1"},
		// {"Turaco", "ASM70936v1"},
		// {"GreyCrane", "ASM70989v1"},
		// {"Hornbill", "ASM71030v1"},
		// {"BaldEagle", "Haliaeetus_leucocephalus_4.0"},
		// {"HoodedCrow", "Hooded_Crow_genome"},
		// {"ChimneySwift", "ChaPel_1.0"},
		// {"GoldenEagle", "Aquila_chrysaetos_1.0.2"},
		// {"BlackGrouse", "tetTet1"},
		// {"Bobwhite", "NB1.1"},
		// {"C_elegans", "WBcel235"}, //Protostomia
		{"C_briggsae", "ASM455v1"},
		// {"MengenillaMoldrzyki", "Memo_1.0"},
		// {"A_darlingi", "A_darlingi_v1"},
		// {"GlossinaMorsitans", "ASM107743v1"},
		// {"AchipteriaColeoptrata", "SM98876v1"},
		// {"AedesAegypti", "AaegL2"},
		// {"AgrilusPlanipennis", "Apla_1.0"},
		// {"GermanCockroach", "Bger_1.0"},
		// {"BajaCaliforniaBarkScorpion", "Cexi_1.0"},
		// {"WheatStemSawfly", "Ccin1"},
		// {"CerapachysBiroi", "CerBir1.0"},
		// {"BedBug", "Clec_1.0"},
		// {"CopidosomaFloridanum", "Cflo_1.0"},
		// {"SouthernHouseMosquito", "CulPip1.0"},
		// {"HouseDustMite", "Dfarinae1.0"},
		// {"EurytemoraAffinis", "Eaff_1.0"},
		// {"FrankliniellaOccidentalis", "Focc_1.0"},
		// {"GlossinaFuscipes", "Glossina_fuscipes_3.0.2"},
		// {"GlossinaBrevipalpis", "Glossina_brevipalpis_1.0.3"},
		// {"GlossinaPallidipes", "Glossina_pallidipes_1.0.3"},
		// {"GlossinaAusteni", "Glossina_austeni_1.0.3"},
		// {"HyalellaAzteca", "Hazt_1.0"},
		// {"HypochthoniusRufulus", "ASM98884v1"},
		// {"CastorBeanTick", "ASM97304v1"},
		// {"DeerTick", "JCVI_ISG_i3_1.0"},
		// {"WesternBlackWidow", "Lhes_1.0"},
		// {"Caddisfly", "Llun_1.0"},
		// {"MelitaeaCinxia", "MelCinx1.0"},
		// {"ChineseScorpion", "M_martensii_Version_1"},
		// {"MicroplitisDemolitor", "Mdem1"},
		// {"N_vitrip", "Nvit_2.1"},
		// {"NilaparvataLugens", "NilLug1.0"},
		// {"MilkweedBug", "Ofas_1.0"},
		// {"TaurusScarab", "Otau_1.0"},
		// {"OrussusAbietinus", "Oabi_1.0"},
		// {"PachypsyllaVenusta", "Pven_1.0"},
		// {"BodyLouse", "JCVI_LOUSE_1.0"},
		// {"PlatynothrusPeltifer", "ASM98890v1"},
		// {"CattleTick", "CCG_Rmi_1.0"},
		// {"DogMite", "SarSca1.0"},
		// {"SteganacarusMagnus", "ASM98888v1"},
		// {"SocialSpider", "Stegodyphus_mimosarum_v1"},
		// {"RedFlourBeetle", "Tcas_3.0"},
		// {"TrichogrammaPretiosum", "Tpre_1.0"},
		// {"EusocialTermite", "ZooNev1.0"},
		// {"Microworm", "Pred3"},
		// {"PigRoundworm", "AscSuum_1.0"},
		// {"BeneficialNematode", "Heterorhabditis_bacteriophora_7.0"},
		// {"PorkWorm", "Trichinella_spiralis_3.7.1"},
		// {"C_11", "Caenorhabditis_sp11_JU1373_3.0.1"},
		// {"C_brenneri", "C_brenneri_6.0.1b"},
		// {"C_angaria", "ps1010rel4"},
		// {"C_japonica", "C_japonica_7.0.1"},
		// {"HumanHookworm", "N_americanus_v1"},
		// {"S_monti", "S_monti_v1"},
		// {"RedDeerNematode", "EEL001"},
		// {"O_vol", "OVOC001"},
		// {"WireWorm", "HCON"},
		// {"AsianBeetle", "Agla_1.0"},
		// {"D_pseudo", "Dpse_3.0"},
		// {"Housefly", "Musca_domestica_2.0.2"},
		// {"Dragonfly", "Lful_1.0"},
		// {"D_miranda", "DroMir_2.2"},
		// {"MountainPineBeetle", "DendPond_male_1.0"},
		// {"A_quad", "Anop_quad_QUAD4_A_V1"},
		// {"A_funestus", "Anop_fune_FUMOZ_V1"},
		// {"A_epiro", "Anop_epir_epiroticus2_V1"},
		// {"A_alb", "Anop_albi_ALBI9_A_V1"},
		// {"A_dirus", "Anop_diru_WRAIR2_V1"},
		// {"A_christyi", "Anop_chri_ACHKN1017_V1"},
		// {"A_arab", "Anop_arab_DONG5_A_V1"},
		// {"A_gambiae", "anoGam1"},
		// {"A_minimus", "Anop_mini_MINIMUS1_V1"},
		// {"Medfly", "Ccap_1.0"},
		// {"TurnipSawfly", "Aros_1.0"},
		// {"D_eugra", "Deug_2.0"},
		// {"D_rho", "Drho_2.0"},
		// {"D_bipect", "Dbip_2.0"},
		// {"D_taka", "Dtak_2.0"},
		// {"D_kikka", "Dkik_2.0"},
		// {"D_elegans", "Dele_2.0"},
		// {"D_bia", "Dbia_2.0"},
		// {"D_ficus", "Dfic_2.0"},
		// {"DiamondbackMoth", "DBM_FJ_V1.1"},
		// {"PostmanButterfly", "ASM31383v2"},
		// {"D_albo", "DroAlb_1.0"},
		// {"Sandfly", "Llon_1.0"},
		// {"WesternPredatoryMite", "Mocc_1.0"},
		// {"TobaccoHornworm", "Msex_1.0"},
		// {"PhleboSandfly", "Ppap_1.0"},
		// {"AlfalfaLeafcutterBee", "MROT_1.0"},
		// {"DwarfHoneyBee", "Aflo_1.0"},
		// {"S_maritima", "Smar_1.0"},
		// {"RedSpiderMite", "ASM23943v1"},
		// {"MonarchButterfly", "DanPle_1.0"},
		// {"CommonEasternBumbleBee", "BIMP_2.0"},
		// {"BufftailedBumbleBee", "Bter_1.0"},
		// {"ArgentineAnt", "Lhum_UMD_V04"},
		// {"PanamanianLeafcutterAnt", "Aech_3.9"},
		// {"PeaAphid", "Acyr_2.0"},
		// {"WaterFlea", "V1.0"},
		// {"RedFireAnt", "Si_gnG"},
		// {"RedHarvesterAnt", "Pbar_UMD_V03"},
		// {"AssassinBug", "Rhodnius_prolixus_3.0.1"},
		// {"LeafcutterAnt", "Attacep1.0"},
		// {"HessianFly", "Mdes_1.0"},
		// {"Silkworm", "ASM15162v1"},
		// {"JumpingAnt", "HarSal_1.0"},
		// {"BlackCarpenterAnt", "CamFlo_1.0"},
		// {"A_melas", "Anop_mela_CM1001059_A_V2"},
		// {"A_farauti", "Anop_fara_FAR1_V2"},
		// {"A_merus", "Anop_meru_MAF_V1"},
		// {"AtlanticHorseshoeCrab", "Limulus_polyphemus_2.1.2"},
		// {"GreenDrake", "Edan_1.0"},
		// {"PollinatingWasp", "CerSol_1.0"},
		// {"ColoradoPotatoBeetle", "Ldec_1.5"},
		// {"AsianCitrusPsyllid", "Diaci_psyllid_genome_assembly_version_1.1"},
		// {"A_atro", "Anop_atro_EBRO_V1"},
		// {"A_cul", "Anop_culi_species_A_37_1_V1"},
		// {"A_maculatus", "Anop_macu_maculatus3_V1"},
		// {"D_suzukii", "Dsuzukii.v01"},
		// {"CommonHouseSpider", "Ptep_1.0"},
		// {"GiantHoneyBee", "Apis_dorsata_1.3"},
		// {"HoneyBee", "apiMel3"},
		// {"A_sinensis", "AS2"},
		// {"A_steph", "ASM30077v2"},
		// {"D_willistoni", "dwil_caf1"},
		// {"N_gir", "Ngir_1.0"},
		// {"N_long", "Nlon_1.0"},
		// {"CatusWorm", "Priapulus_caudatus_4.0.1"},
		// {"D_ananassae", "droAna3"},
		// {"D_erecta", "droEre2"},
		// {"D_melano", "dm6"},
		// {"D_yakuba", "droYak2"},
		// {"D_grimshawi", "droGri2"},
		// {"D_mojavensis", "droMoj3"},
		// {"D_persimilis", "droPer1"},
		// {"D_sechellia", "droSec1"},
		// {"D_simulans", "droSim1"},
		// {"D_virilis", "droVir3"},
		// {"KiwiFruit", "Kiwifruit_v1"}, //GreenPlants
		// {"TauschsGoatgrass", "ASM34733v1"},
		// {"AethionemaArabicum", "VEGI_AA_v_1.0"},
		// {"PrincesFeather", "AHP_1.0"},
		// {"TallWaterhemp", "ASM18065v1"},
		// {"Amborellaceae", "AMTR1.0"},
		// {"AquilariaAgallochum", "Aquilaria_agallocha_v1"},
		// {"ArabidopsisHalleri_subsp_gemmifera", "Ahal_1.0"},
		// {"ArabidopsisLyrata_subsp_lyrata", "v.1.0"},
		// {"ArabidopsisThaliana", "TAIR10"},
		// {"AlpineRockcress", "A_alpina_V4"},
		// {"GreenMicroalga", "ASM73321v1"},
		// {"IndianLilac", "AzaInd2.0"},
		// {"BetaVulgaris_subsp_vulgaris", "RefBeet_1.2.1"},
		// {"DwarfBirch", "ASM32700v1"},
		// {"PurpleFalseBrome", "v1.0"},
		// {"Rapeseed", "Brassica_napus_assembly_v1.0"},
		// {"BrassicaOleracea_var_oleracea", "BOL"},
		// {"FieldMustard", "Brapa_1.0"},
		// {"PigeonPea", "Cajanus_cajan_Asha_ver1.0"},
		// {"Camelina", "Cs"},
		// {"CannabisSativa", "canSat3"},
		// {"Capsella", "Caprub1_0"},
		// {"HotPepper", "PGAv.1.5"},
		// {"Papaya", "Papaya1.0"},
		// {"ChineseChestnut", "ASM76360v1"},
		// {"Chlamydomonas", "v3.0"},
		// {"ChlorellaVariabilis", "v1.0"},
		// {"Chickpea", "ASM33114v1"},
		// {"Watermelon", "CiLa_1.0"},
		// {"Clementine", "Citrus_clementina_v1.0"},
		// {"Orange", "Csi_valencia_1.0"},
		// {"CoccomyxaSubellipsoidea_C_169", "Coccomyxa_subellipsoidae_v2.0"},
		// {"Horseweed", "ASM77593v1"},
		// {"Muskmelon", "ASM31304v1"},
		// {"Cucumber", "CucSat_1.0"},
		// {"ClovePink", "DCA_r1.0"},
		// {"CaucasianPersimmon", "ASM77412v1"},
		// {"AfricanOilPalm", "EG5"},
		// {"AmericanOilPalm", "EO8"},
		// {"EthiopianBanana", "v1.1"},
		// {"DancingDollOrchid", "EpS81.0"},
		// {"RiverRedGum", "EUC_r1.0"},
		// {"RoseGum", "Egrandis1_0"},
		// {"EutremaParvulum", "Eutrema_parvulum_v01"},
		// {"SaltwaterCress", "Eutsalg1_0"},
		// {"JapaneseStrawberry", "FII_r1.1"},
		// {"HonshuStrawberry", "FNI_r1.1"},
		// {"HimalayanStrawberry", "FNU_r1.1"},
		// {"WildAsianStrawberry", "FOR_r1.1"},
		// {"WoodlandStrawberry", "FraVesHawaii_1.0"},
		// {"GardenStrawberry", "FANhybrid_r1.2"},
		// {"EuropeanAsh", "BATG_0.4"},
		// {"CorkscrewPlant", "GenAur_1.0"},
		// {"GlycineMax", "V1.1"},
		// {"GlycineSoja", "W05v1.0"},
		// {"TreeCotton", "Gossypium_arboreum_v1.0"},
		// {"NewWorldCotton", "Graimondii2_0"},
		// {"Helicosporidium_sp_ATCC_50920", "Helico_v1.0"},
		// {"RubberTree", "Hevbra_1.0"},
		// {"HordeumPubiflorum", "Hordeum_pubiflorum_assembly1"},
		// {"PhysicNut", "JatCur_1.0"},
		// {"KlebsormidiumFlaccidum", "ASM70883v1"},
		// {"Lettuce", "Legassy_V2"},
		// {"Calabash", "Bottle_gourd"},
		// {"AlabamaGladecress", "VEGI_LA_v_1.0"},
		// {"LeersiaPerrieri", "Lperr_V1.4"},
		// {"Flax", "LinUsi_v1.1"},
		// {"BirdsfootTrefoil", "ASM18111v1"},
		// {"BlueLupin", "Lupin_genome_scaffold"},
		// {"Apple", "MalDomGD1.0"},
		// {"ManihotEsculenta_subsp_flabellifolia", "MW_v2d"},
		// {"BarrelClover", "MedtrA17_3.5"},
		// {"MicromonasPusilla_CCMP1545", "Micromonas_pusilla_CCMP1545_v2.0"},
		// {"Micromonas_sp_RCC299", "ASM9098v2"},
		// {"Monkeyflower", "Mimgu1_0"},
		// {"MulberryTree", "ASM41409v2"},
		// {"MusaAcuminata_subsp_malaccensis", "ASM31385v1"},
		// {"SacredLotus", "Chinese_Lotus_1.1"},
		// {"NicotianaOtophora", "Noto"},
		// {"NicotianaSylvestris", "Nsyl"},
		// {"NicotianaTabacum", "Ntab_K326"},
		// {"NicotianaTomentosiformis", "Ntom_v01"},
		// {"OryzaBarthii", "O.barthii_v1.3"},
		// {"OryzaBrachyantha", "Oryza_brachyantha.v1.4b"},
		// {"OryzaGlumipatula", "Oryza_glumaepatula_v1.5"},
		// {"OryzaLongistaminata", "O_longistaminata_v1.0"},
		// {"OryzaMeridionalis", "Oryza_meridionalis_v1.3"},
		// {"OryzaNivara", "Oryza_nivara_v1.0"},
		// {"OryzaPunctata", "Oryza_punctata_v1.2"},
		// {"OryzaSativa_JaponicaGroup", "Build_4.0"},
		// {"OstreococcusLucimarinus_CCE9901", "ASM9206v1"},
		// {"OstreococcusTauri", "version_050606"},
		// {"ScarletBugler", "ASM73743v1"},
		// {"GrinnellsBeardtongue", "ASM73742v1"},
		// {"CommonBean", "PhaVulg1_0"},
		// {"DatePalm", "DPV01"},
		// {"Moss", "V1.1"},
		// {"LoblollyPine", "PtaedaFosmidLib.0.8"},
		// {"DesertPoplar", "PopEup_1.0"},
		// {"BlackCottonwood", "Poptr2_0"},
		// {"Cowslip", "ASM78844v1"},
		// {"JapaneseApricot", "P.mume_V1.0"},
		// {"Peach", "Prupe1_0"},
		// {"ChineseWhitePear", "Pbr_v1.0"},
		// {"RaphanusRaphanistrum_subsp_raphanistrum", "ASM76984v1"},
		// {"Radish", "ASM80110v1"},
		// {"CastorOilPlant", "JCVI_RCG_1.1"},
		// {"Spikemoss", "v1.0"},
		// {"Sesame", "S_indicum_v1.0"},
		// {"FoxtailMillet", "Setaria_V1"},
		// {"LondonRocket", "VEGI_SI_v_1.0"},
		// {"SolanumArcanum", "Soarc10"},
		// {"SolanumHabrochaites", "Sohab10"},
		// {"Tomato", "SL2.40"},
		// {"Eggplant", "SME_r2.5.1"},
		// {"SolanumPennellii", "Sopen10"},
		// {"CurrantTomato", "Sol_pimpi_v1.0"},
		// {"Potato", "SolTub_3.0"},
		// {"Sorghum", "Sorbi1"},
		// {"Spinach", "Viroflay_1.0.1"},
		// {"GreaterDuckweed", "Spirodela_polyrhiza_v01"},
		// {"PinkQueen", "ASM46358v1"},
		// {"Cacao", "Theobroma_cacao_20110822"},
		// {"RedClover", "Tp1.0"},
		// {"WildEinkorn", "ASM34745v1"},
		// {"LargeCranberry", "ASM77533v1"},
		// {"AdzukiBean", "Vigna_angularis"},
		// {"VignaRadiata_var_radiate", "Vradiata_ver6"},
		// {"CommonGrapeVine", "12X"},
		// {"VolvoxCarteri_f_nagariensis", "v1.0"},
		// {"Maize", "B73_RefGen_v3"},
		// {"ManchurianWildRice", "Zizania_latifolia_v01"},
		// {"VaseTunicate", "KH"}, //Other
		// {"Larvaceans", "ASM20953v1"},
		// {"CionaSavignyi", "ASM14926v1"},
		// {"SchistosomaJaponicum", "ASM15177v1"},
		// {"SchistosomaHaematobium", "SchHae_1.0"},
		// {"SchistosomaCurassoni", "S_curassoni_Dakar"},
		// {"SchistosomaMargrebowiei", "S_margrebowiei_Zambia"},
		// {"SchistosomaMattheei", "S_mattheei_Denwood"},
		// {"SchistosomaRodhaini", "S_rodhaini_Burundi"},
		// {"StarletSeaAnemone", "ASM20922v1"},
		// {"Trichoplax", "v1.0"},
		// {"StarAscidian", "356a_chromosome_assembly"},
		// {"BatStar", "Pmin_1.0"},
		// {"GreenSeaUrchin", "Lvar_0.4"},
		// {"PurpleSeaUrchin", "Spur_3.1"},
		// {"AcornWorm", "Skow_1.1"},
		// {"Schisto", "ASM23792v2"},
		// {"ChineseLiverFluke", "C_sinensis_2.0"},
		// {"RodentTapeworm", "HMIC001"},
		// {"HyperTapeworm", "EGRAN001"},
		// {"FoxTapeworm", "EMULTI001"},
		// {"CaliforniaSeaHare", "AplCal3.0"},
		// {"PolychaeteWorm", "Capca1"},
		// {"Leech", "Helobdella_robusta_v1.0"},
		// {"OwlLimpet", "Helro1"},
		// {"PacificOyster", "oyster_v9"},
		// {"Rotifer", "ASM_PRJEB1171_v1"},
		// {"FreshwaterSnail", "ASM45736v1"},
		// {"FreshwaterPolyp", "Hydra_RP_1.0"},
		// {"Lancelet", "Version_2"},
		// {"WartyCombJelly", "MneLei_Aug2011"},
		// {"SeaSponge", "v1.0"},
		// {"SaltwaterCroc", "Cpor_2.0"}, //New
		// {"GharialCroc", "ggan_v0.2"},
		// {"Viper", "Vber.be_1.0"},
		// {"KingCobra", "OphHan1.0"},
		// {"PitViper", "CrotMitch1.0"},
		// {"EuropeanEel", "Anguilla_anguilla_v1_09_nov_10"},
		// {"JapaneseEel", "japanese_eel_genome_v1_25_oct_2011_japonica_c401b40"},
		// {"MormonButterfly", "Ppol_1.0"},
		// {"SwallowtailButterfly", "Pxut_1.0"},
		// {"TigerButterfly", "pgl_assembly_v1"},
		// {"ProboscisMonkey", "Charlie1.0"},
		// {"PolarBear", "UrsMar_1.0"},
		// {"FlyingLemur", "G_variegatus_3.0.2"},
		// {"PrzewalskiHorse", "Burgud"},
		// {"ChinesePangolin", "M_pentadactyla_1.1.1"},
		// {"DamaralandMolerat", "DMR_v1.0"},
		// {"Bison", "Bison_UMD1.0"},
		// {"Dromedary", "PRJNA234474_Ca_dromedarius_V1.0"},
		// {"SnubnosedMonkey", "Rrox_v1"},
		// {"BrazilianGuineaPig", "CavAp1.0"},
		// {"TibetanPig", "Tibetan_Pig_v1.0"},
		// {"MongolianHorse", "Ajinai1.0"},
		// {"WildSheep", "Oori1"},
		// {"MiniPig", "SscrofaMinipig"},
		// {"hg19", "hg19"},
		// {"Shark", "shark_ass_20150106_scafSeq"},
		// {"PogonaVitticeps", "Pogona_vitticeps.male"},
	} {
		unmaskChr(gen.genomeName, "", "fa", gen.identifier)
	}
}
