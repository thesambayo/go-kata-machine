package tests

import (
	"testing"
)

func TestLinearSearch(t *testing.T) {
	linearSearch, err := GetLinearSearchFunc(*day)
	if err != nil {
		t.Skipf("Cannot find linear search funcion for %v", *day)
		// return
	}

	t.Run("", func(t *testing.T) {
		arr := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}
		assertError(t, linearSearch(arr, 69), true)
		assertError(t, linearSearch(arr, 1336), false)
		assertError(t, linearSearch(arr, 69420), true)
		assertError(t, linearSearch(arr, 69421), false)
		assertError(t, linearSearch(arr, 1), true)
		assertError(t, linearSearch(arr, 0), false)
	})
}
