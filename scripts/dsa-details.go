package scripts

//type DsaMethods struct {
//	name        string
//	argsTypes   string
//	returnTypes string
//}

type DSADetails struct {
	DsaType     string
	ArgsTypes   string
	ReturnsType string
	//method      DsaMethods
}

func GetDsaDetails() map[string]DSADetails {
	return map[string]DSADetails{
		//"LinearSearchList": {
		//	DsaType:     "fn",
		//	ArgsTypes:   "haystack []int, needle int",
		//	ReturnsType: "bool",
		//},
		//"BinarySearchList": {
		//	DsaType:     "fn",
		//	ArgsTypes:   "haystack []int, needle int",
		//	ReturnsType: "bool",
		//},
		"TwoCrystalBalls": {
			DsaType:     "fn",
			ArgsTypes:   "haystack []bool",
			ReturnsType: "int",
		},
		"BubbleSort": {
			DsaType:     "fn",
			ArgsTypes:   "arr []int",
			ReturnsType: "[]int",
		},
	}
}
