package main

import (
	"testing"
)

func TestMaximizeDelta_Minimization(t *testing.T) {
	values := []int{10, 5, 12, 4, 15, 2, 20}
	want := 18
	got := maximizeDelta(values, true)

	if got != want {
		t.Errorf("Expected %d, but got %d", want, got)
	}
}

func TestMaximizeDelta_Maximization(t *testing.T) {
	values := []int{10, 15, 8, 12, 5, 18, 3}
	want := 15
	got := maximizeDelta(values, false)

	if want != got {
		t.Errorf("Expected %d, but got %d", want, got)
	}
}

func TestMaximizeDelta_EmptySequence(t *testing.T) {
	values := []int{}
	want := 0
	got := maximizeDelta(values, true)

	if want != got {
		t.Errorf("Expected %d, but got %d", want, got)
	}
}

func TestMaximizeDelta_ConstantValues(t *testing.T) {
	values := []int{5, 5, 5, 5}
	want := 0
	got := maximizeDelta(values, true)

	if want != got {
		t.Errorf("Expected %d, but got %d", want, got)
	}
}

func TestMaximizeDelta_TwoElementsDecreasing(t *testing.T) {
	values := []int{10, 3}
	want := 7
	got := maximizeDelta(values, false)

	if want != got {
		t.Errorf("Expected %d, but got %d", want, got)
	}
}

func TestMaximizeDelta_TwoElementsIncreasing(t *testing.T) {
	values := []int{3, 10}
	want := 7 // Delta is 10 - 3 = 7
	got := maximizeDelta(values, true)

	if want != got {
		t.Errorf("Expected %d, but got %d", want, got)
	}
}

func TestMaximizeDelta_SingleElement(t *testing.T) {
	values := []int{5}
	want := 0 // No delta can be calculated from a single element
	got := maximizeDelta(values, true)

	if want != got {
		t.Errorf("Expected %d, but got %d", want, got)
	}
}
