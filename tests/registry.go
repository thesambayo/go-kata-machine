package tests

import (
	"flag"
	"go-kata-machine/day4"
	"go-kata-machine/day5"
	"go-kata-machine/day6"
)

// var day string
var day = flag.String("day", "day1", "Name of location to greet")

type BinarySearchFunc func([]int, int) bool
type BubbleSortFunc func([]int) []int
type LinearSearchFunc func([]int, int) bool
type TwoCrystalBallsFunc func([]bool) int

type Day struct {
	binarySearch    BinarySearchFunc
	bubbleSort      BubbleSortFunc
	linearSearch    LinearSearchFunc
	twoCrystalBalls TwoCrystalBallsFunc
}

// Registry maps day identifiers to their corresponding implementations.
var Registry = map[string]Day{
	"day6": {
		binarySearch:    day6.BinarySearchList,
		bubbleSort:      day6.BubbleSort,
		linearSearch:    day6.LinearSearch,
		twoCrystalBalls: day6.TwoCrystalBalls,
	},
	"day5": {
		binarySearch:    day5.BinarySearchList,
		bubbleSort:      day5.BubbleSort,
		linearSearch:    day5.LinearSearch,
		twoCrystalBalls: day5.TwoCrystalBalls,
	},
	"day4": {
		binarySearch:    day4.BinarySearchList,
		bubbleSort:      day4.BubbleSort,
		linearSearch:    day4.LinearSearchList,
		twoCrystalBalls: day4.TwoCrystalBalls,
	},
}
