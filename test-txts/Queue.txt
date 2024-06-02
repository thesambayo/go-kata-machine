package queue

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
