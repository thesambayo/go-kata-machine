package tests

import (
	"flag"
	"go-kata-machine/day4"
	"go-kata-machine/day5"
	"go-kata-machine/day6"
	"go-kata-machine/day7"
)

// var day string
var day = flag.String("day", "day1", "Name of location to greet")

type BinarySearchFunc func([]int, int) bool
type BubbleSortFunc func([]int) []int

type Day struct {
	binarySearch BinarySearchFunc
	bubbleSort   BubbleSortFunc
}

// Registry maps day identifiers to their corresponding implementations.
var Registry = map[string]Day{
	"day7": {
		binarySearch: day7.BinarySearchList,
		bubbleSort:   day7.BubbleSort,
	},
	"day6": {
		binarySearch: day6.BinarySearchList,
		bubbleSort:   day6.BubbleSort,
	
},
	"day4": {
		binarySearch: day4.BinarySearchList,
		bubbleSort:   day4.BubbleSort,
	
},
	"day5": {
		binarySearch: day5.BinarySearchList,
		bubbleSort:   day5.BubbleSort,
	},
}
