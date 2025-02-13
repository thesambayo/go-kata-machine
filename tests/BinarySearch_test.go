package tests

import (
	"testing"
)

func TestBinarySearchList(t *testing.T) {

	binarySearch, err := GetBinarySearchFunc(*day)
	if err != nil {
		// t.Errorf("Cannot find Binary search funcion for %v", *day)
		t.Skipf("Cannot find  Binary search funcion for %v", *day)
		return
	}

	t.Run("binary search", func(t *testing.T) {
		arr := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}
		assertError(t, binarySearch(arr, 69), true)
		assertError(t, binarySearch(arr, 1336), false)
		assertError(t, binarySearch(arr, 69420), true)
		assertError(t, binarySearch(arr, 69421), false)
		assertError(t, binarySearch(arr, 1), true)
		assertError(t, binarySearch(arr, 0), false)
	})
}
