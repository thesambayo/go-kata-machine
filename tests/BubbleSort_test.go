package tests

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	bubbleSort, err := GetBubbleSortFunc(*day)
	if err != nil {
		t.Skipf("Cannot find  bubble sort funcion for %v", *day)
		// return
	}
	t.Run("bubble sort", func(t *testing.T) {
		unsorted := []int{71, 81, 90, 69420, 420, 1, 3, 99, 4, 69, 1337}
		got := bubbleSort(unsorted)
		expected := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}
		assertSlicesError(t, got, expected)
	})
}
