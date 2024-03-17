package binarysearch

import (
	"math"
)

func BinarySearch(haystack []int, needle int) bool {
	lowPoint := 0
	highPoint := len(haystack)

	for lowPoint < highPoint {
		intMiddlePoint := (lowPoint) + ((highPoint - lowPoint) / 2)
		middlePoint := int(math.Floor(float64(intMiddlePoint)))
		middleValue := haystack[middlePoint]

		if needle == middleValue {
			return true
		}

		if needle > middleValue {
			lowPoint = middlePoint + 1
		}

		if needle < middleValue {
			highPoint = middlePoint
		}

	}

	return false
}
