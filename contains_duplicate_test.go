package main

import (
	"math"
	"testing"
)

// Given an integer array nums,
//  return true if any value appears at least twice in the array, and return false if every element is distinct.

// Example 1:

// Input: nums = [1,2,3,1]

// Output: true

// Explanation:

// The element 1 occurs at the indices 0 and 3.

// Example 2:

// Input: nums = [1,2,3,4]

// Output: false

// Explanation:

// All elements are distinct.

// Example 3:

// Input: nums = [1,1,1,3,3,4,3,2,4,2]

// Output: true

// Constraints:

//     1 <= nums.length <= 105
//     -109 <= nums[i] <= 109

// containsDuplicate checks if there are any duplicates in the array.

func TestContainsDuplicateExample1(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	want := true
	got := containsDuplicate(nums)
	if got != want {
		t.Errorf("Test failed for input %v: expected %v, got %v", nums, want, got)
	}
}

func TestContainsDuplicateExample2(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	want := false
	got := containsDuplicate(nums)
	if got != want {
		t.Errorf("Test failed for input %v: expected %v, got %v", nums, want, got)
	}
}

func TestContainsDuplicateExample3(t *testing.T) {
	nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	want := true
	got := containsDuplicate(nums)
	if got != want {
		t.Errorf("Test failed for input %v: expected %v, got %v", nums, want, got)
	}
}

func TestContainsDuplicateEmptyArray(t *testing.T) {
	nums := []int{}
	want := false
	got := containsDuplicate(nums)
	if got != want {
		t.Errorf("Test failed for input %v: expected %v, got %v", nums, want, got)
	}
}

func TestContainsDuplicateSingleElement(t *testing.T) {
	nums := []int{5}
	want := false
	got := containsDuplicate(nums)
	if got != want {
		t.Errorf("Test failed for input %v: expected %v, got %v", nums, want, got)
	}
}

func TestContainsDuplicateLargeArrayWithDuplicates(t *testing.T) {
	nums := make([]int, 100000)
	for i := 0; i < 99999; i++ {
		nums[i] = i
	}
	nums[99999] = 50000 // Introduce a duplicate
	want := true
	got := containsDuplicate(nums)
	if got != want {
		t.Errorf("Test failed for large input: expected %v, got %v", want, got)
	}
}

func TestContainsDuplicateLargeArrayWithoutDuplicates(t *testing.T) {
	nums := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		nums[i] = i
	}
	want := false
	got := containsDuplicate(nums)
	if got != want {
		t.Errorf("Test failed for large input: expected %v, got %v", want, got)
	}
}

func TestContainsDuplicateTwoElementsNoDuplicate(t *testing.T) {
	nums := []int{1, 2}
	want := false
	got := containsDuplicate(nums)
	if got != want {
		t.Errorf("Test failed for input %v: expected %v, got %v", nums, want, got)
	}
}

func TestContainsDuplicateTwoElementsWithDuplicate(t *testing.T) {
	nums := []int{1, 1}
	want := true
	got := containsDuplicate(nums)
	if got != want {
		t.Errorf("Test failed for input %v: expected %v, got %v", nums, want, got)
	}
}

func TestContainsDuplicateNegativeNumbersWithDuplicate(t *testing.T) {
	nums := []int{-1, -2, -3, -1}
	want := true
	got := containsDuplicate(nums)
	if got != want {
		t.Errorf("Test failed for input %v: expected %v, got %v", nums, want, got)
	}
}

func TestContainsDuplicateNegativeNumbersNoDuplicate(t *testing.T) {
	nums := []int{-1, -2, -3, -4}
	want := false
	got := containsDuplicate(nums)
	if got != want {
		t.Errorf("Test failed for input %v: expected %v, got %v", nums, want, got)
	}
}

func TestContainsDuplicateMaxIntWithDuplicate(t *testing.T) {
	nums := []int{math.MaxInt32, math.MaxInt32}
	want := true
	got := containsDuplicate(nums)
	if got != want {
		t.Errorf("Test failed for input %v: expected %v, got %v", nums, want, got)
	}
}

func TestContainsDuplicateMaxIntNoDuplicate(t *testing.T) {
	nums := []int{math.MaxInt32, math.MaxInt32 - 1}
	want := false
	got := containsDuplicate(nums)
	if got != want {
		t.Errorf("Test failed for input %v: expected %v, got %v", nums, want, got)
	}
}

func TestContainsDuplicateLargeRange(t *testing.T) {
	nums := []int{-100000, 0, 100000}
	want := false
	got := containsDuplicate(nums)
	if got != want {
		t.Errorf("Test failed for input %v: expected %v, got %v", nums, want, got)
	}
}
