package scripts

type DsaMethod struct {
	Name        string
	ArgsTypes   string
	ReturnTypes string
}

type DSADetails struct {
	DsaType     string
	ArgsTypes   string
	ReturnsType string
	Generic     string
	Methods     []DsaMethod
}

// TODO: tests for LinearSearchList and BinarySearchList, not ready

func GetDsaDetails() map[string]DSADetails {
	return map[string]DSADetails{
		"LinearSearchList": {
			DsaType:     "fn",
			ArgsTypes:   "haystack []int, needle int",
			ReturnsType: "bool",
		},
		"BinarySearchList": {
			DsaType:     "fn",
			ArgsTypes:   "haystack []int, needle int",
			ReturnsType: "bool",
		},
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
		"Stack": {
			DsaType: "struct",
			Generic: "T",
			Methods: []DsaMethod{
				{
					Name:        "Push",
					ArgsTypes:   "item T",
					ReturnTypes: "",
				},
				{
					Name:        "Pop",
					ArgsTypes:   "",
					ReturnTypes: "T",
				},
				{
					Name:        "Peek",
					ArgsTypes:   "",
					ReturnTypes: "T",
				},
			},
		},
		"Queue": {
			DsaType: "struct",
			Generic: "T",
			Methods: []DsaMethod{
				{
					Name:        "Enqueue",
					ArgsTypes:   "item T",
					ReturnTypes: "",
				},
				{
					Name:        "Deque",
					ArgsTypes:   "",
					ReturnTypes: "T",
				},
				{
					Name:        "Peek",
					ArgsTypes:   "",
					ReturnTypes: "T",
				},
			},
		},
	}
}
