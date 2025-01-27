package tests

import "fmt"

// GetBinarySearchFunc retrieves the binary search function for a specific day.
func GetBinarySearchFunc(day string) (BinarySearchFunc, error) {
	if day, exists := Registry[day]; exists {
		return day.binarySearch, nil
	}
	return nil, fmt.Errorf("no implementation found for day: %s", day)
}

// GetBinarySearchFunc retrieves the binary search function for a specific day.
func GetBubbleSortFunc(day string) (BubbleSortFunc, error) {
	if day, exists := Registry[day]; exists {
		return day.bubbleSort, nil
	}
	return nil, fmt.Errorf("no implementation found for day: %s", day)
}
