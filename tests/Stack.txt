package stack

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
