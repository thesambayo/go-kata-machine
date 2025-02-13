package tests

import (
	"math/rand"
	"testing"
)

func GenerateRandomNumber(minBound, maxBound int) int {
	return rand.Intn(maxBound-minBound) + minBound
}

func TestTwoCrystalBalls(t *testing.T) {
	twoCrystallBalls, err := GetTwoCrystalBallFunc(*day)
	if err != nil {
		t.Skipf("Cannot find  two crystal balls funcion for %v", *day)
		// return
	}

	t.Run("two crystal balls where there is a break", func(t *testing.T) {
		numOfFlights := 10000
		idx := GenerateRandomNumber(1, numOfFlights)
		testData := make([]bool, numOfFlights)
		for i := idx; i < numOfFlights; i++ {
			testData[i] = true
		}

		got := twoCrystallBalls(testData)
		expected := idx
		assertError(t, got, expected)
	})

	t.Run("two crystal balls where there is no break", func(t *testing.T) {
		numOfFlights := 1000
		testData := make([]bool, numOfFlights)
		got := twoCrystallBalls(testData)
		expected := -1
		assertError(t, got, expected)
	})
}
