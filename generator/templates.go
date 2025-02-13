package generator

// Template definitions for new files
var templates = map[string]string{
	"binarysearch.go": `package day%d

func BinarySearchList(haystack []int, needle int) bool {
    // TODO: implement
    return false
}
`,
	"bubblesort.go": `package day%d

func BubbleSort(arr []int) []int {
    // TODO: implement
    return arr
}
`,
	"linearsearch.go": `package day%d

func LinearSearch(haystack []int, needle int) bool {
	return false
}
`,
	"twocrystallballs.go": `package day%d

func TwoCrystalBalls(haystack []bool) int {
	return -1
}
`,
}
