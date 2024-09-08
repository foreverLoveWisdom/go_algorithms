package main

import (
	"reflect"
	"sort"
	"testing"
)

// Find All Pairs with Sum in an Unsorted Array

// **Problem:**

// Given an unsorted array `nums = [1, 3, 2, 2, 4, 3]` and `target = 5`,
// find all pairs of indices such that the numbers at those indices add up to the target.

// **Example:**

// - **Input:** `nums = [1, 3, 2, 2, 4, 3]`, `target = 5`
// - **Output:** `[[0, 4], [1, 2], [3, 5]]`

// **Explanation:**
// - `nums[0] + nums[4] = 1 + 4 = 5`
// - `nums[1] + nums[2] = 3 + 2 = 5`
// - `nums[3] + nums[5] = 2 + 3 = 5`
// Constraints: one element can be used only once.

func equalIgnoreOrder(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	// Sort the inner slices
	for i := range a {
		sort.Ints(a[i])
	}
	for i := range b {
		sort.Ints(b[i])
	}

	// Sort the outer slices
	sort.Slice(a, func(i, j int) bool {
		return a[i][0] < a[j][0] || (a[i][0] == a[j][0] && a[i][1] < a[j][1])
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i][0] < b[j][0] || (b[i][0] == b[j][0] && b[i][1] < b[j][1])
	})

	return reflect.DeepEqual(a, b)
}

func TestPairsUnsortedArr_DuplicatedElements(t *testing.T) {
	nums := []int{1, 3, 2, 2, 4, 3, 3, 2}
	got := pairsUnsortedArr(nums, 5)
	want := [][]int{{0, 4}, {1, 2}, {3, 5}, {6, 7}}
	if !equalIgnoreOrder(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestPairsUnsortedArr_EmptyArray(t *testing.T) {
	nums := []int{}
	got := pairsUnsortedArr(nums, 5)
	want := [][]int{}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestPairsUnsortedArr_NoPairsFound(t *testing.T) {
	nums := []int{1, 2, 3}
	got := pairsUnsortedArr(nums, 10)
	want := [][]int{}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestPairsUnsortedArr_SingleElement(t *testing.T) {
	nums := []int{5}
	got := pairsUnsortedArr(nums, 5)
	want := [][]int{}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestPairsUnsortedArr_MultiplePairs(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	got := pairsUnsortedArr(nums, 6)
	want := [][]int{{0, 4}, {1, 3}}
	if !equalIgnoreOrder(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestPairsUnsortedArr_NegativeNumbers(t *testing.T) {
	nums := []int{-1, -2, -3, -4, -5}
	got := pairsUnsortedArr(nums, -8)
	want := [][]int{{2, 4}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestPairsUnsortedArr_ZeroTarget(t *testing.T) {
	nums := []int{0, 0, 0, 0}
	got := pairsUnsortedArr(nums, 0)
	want := [][]int{{0, 1}, {2, 3}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestPairsUnsortedArr_DuplicateNumbers(t *testing.T) {
	nums := []int{1, 1, 1, 1}
	got := pairsUnsortedArr(nums, 2)
	want := [][]int{{0, 1}, {2, 3}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
