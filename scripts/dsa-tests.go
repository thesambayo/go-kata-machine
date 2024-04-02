package scripts

func GetDsaTests() map[string]string {
	return map[string]string{
		"TwoCrystalBalls": `package twocrystalballs

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

`,
		"BubbleSort": `package bubblesort

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

`,
		"Stack": `package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	list := Stack[int]{}
	list.Push(5)
	list.Push(7)
	list.Push(9)

	assertError(t, list.Pop(), 9)
	assertError(t, list.Length, 2)

	list.Push(11)
	assertError(t, list.Pop(), 11)
	assertError(t, list.Pop(), 7)
	assertError(t, list.Peek(), 5)
	assertError(t, list.Pop(), 5)
	assertError(t, list.Pop(), 0)

	assertError(t, list.Length, 0)
	list.Push(69)
	assertError(t, list.Peek(), 69)
	assertError(t, list.Length, 1)
}

func assertError(t testing.TB, got int, expected int) {
	t.Helper()
	if got != expected {
		t.Errorf("stack not correct. got %d, want %d", got, expected)
	}
}
`,
		"Queue": `package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	list := Queue[int]{}
	list.Enqueue(5)
	list.Enqueue(7)
	list.Enqueue(9)

	assertError(t, list.Deque(), 5)
	assertError(t, list.Length, 2)

	list.Enqueue(11)
	assertError(t, list.Deque(), 7)
	assertError(t, list.Deque(), 9)
	assertError(t, list.Peek(), 11)
	assertError(t, list.Deque(), 11)
	assertError(t, list.Deque(), 0)

	assertError(t, list.Length, 0)

	list.Enqueue(69)
	assertError(t, list.Peek(), 69)
	assertError(t, list.Length, 1)
}

func assertError(t testing.TB, got int, expected int) {
	t.Helper()
	if got != expected {
		t.Errorf("slice not sorted. got %d, want %d", got, expected)
	}
}
`,
	}
}
