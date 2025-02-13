package tests

import (
	"flag"
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
var Registry = map[string]Day{}
