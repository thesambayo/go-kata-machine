package twocrystalballs

import (
	"math/rand"
	"testing"
)

func GenerateRandomNumber(minBound, maxBound int) int {
	return rand.Intn(maxBound-minBound) + minBound
}

func TestTwoCrystalBalls(t *testing.T) {
	t.Run("two crystal balls where there is a break", func(t *testing.T) {
		numOfFlights := 10000
		idx := GenerateRandomNumber(1, numOfFlights)
		testData := make([]bool, numOfFlights)
		for i := idx; i < numOfFlights; i++ {
			testData[i] = true
		}

		got := TwoCrystalBalls(testData)
		expected := idx
		assertError(t, got, expected)
	})

	t.Run("two crystal balls where there is no break", func(t *testing.T) {
		numOfFlights := 1000
		testData := make([]bool, numOfFlights)
		got := TwoCrystalBalls(testData)
		expected := -1
		assertError(t, got, expected)
	})
}

func assertError(t testing.TB, got int, expected int) {
	t.Helper()
	if got != expected {
		t.Errorf("got %d, want %d", got, expected)
	}
}
