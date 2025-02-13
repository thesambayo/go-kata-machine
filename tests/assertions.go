package tests

import (
	"reflect"
	"testing"
)

func assertError[T comparable](t testing.TB, got T, expected T) {
	t.Helper()
	if got != expected {
		t.Errorf("got %v, want %v", got, expected)
	}
}

func assertSlicesError(t testing.TB, got []int, expected []int) {
	t.Helper()
	isSame := reflect.DeepEqual(got, expected)
	if !isSame {
		t.Errorf("slice not sorted. got %d, want %d", got, expected)
	}
}
