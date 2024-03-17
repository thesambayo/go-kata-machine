package bubbleSort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	t.Run("", func(t *testing.T) {
		unsorted := []int{71, 81, 90, 69420, 420, 1, 3, 99, 4, 69, 1337}
		got := BubbleSort(unsorted)
		expected := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}
		assertError(t, got, expected)
	})
}

func assertError(t testing.TB, got []int, expected []int) {
	t.Helper()
	isSame := reflect.DeepEqual(got, expected)
	if !isSame {
		t.Errorf("slice not sorted. got %d, want %d", got, expected)
	}
}
