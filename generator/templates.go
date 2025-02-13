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
}
