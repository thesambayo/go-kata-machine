package tests

import "fmt"

// GetBinarySearchFunc retrieves the binary search function for a specific day.
func GetBinarySearchFunc(day string) (BinarySearchFunc, error) {
	if day, exists := Registry[day]; exists {
		return day.binarySearch, nil
	}
	return nil, fmt.Errorf("no %s implementation found for day: %s", "binarySearch", day)
}

// GetBinarySearchFunc retrieves the binary search function for a specific day.
func GetBubbleSortFunc(day string) (BubbleSortFunc, error) {
	if day, exists := Registry[day]; exists {
		return day.bubbleSort, nil
	}
	return nil, fmt.Errorf("no %s implementation found for day: %s", "bubbleSort", day)
}

func GetLinearSearchFunc(day string) (LinearSearchFunc, error) {
	if day, exists := Registry[day]; exists {
		return day.linearSearch, nil
	}
	return nil, fmt.Errorf("no %s implementation found for day: %s", "linearSearch", day)
}

func GetTwoCrystalBallFunc(day string) (TwoCrystalBallsFunc, error) {
	if day, exists := Registry[day]; exists {
		return day.twoCrystalBalls, nil
	}
	return nil, fmt.Errorf("no %s implementation found for day: %s", "twoCrystalBalls", day)
}
