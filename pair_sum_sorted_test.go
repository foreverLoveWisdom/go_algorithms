package main

import (
	"reflect"
	"testing"
)

// Find Pair with Given Sum in a Sorted Array

// Problem:

// Given a sorted array of integers nums and an integer target,
// return indices of the two numbers such that they add up to target. Assume each input would have exactly one solution.

// Constraints:

//     The array is sorted in ascending order.
//     2 <= nums.length <= 10^4
//     -10^9 <= nums[i] <= 10^9
//     -10^9 <= target <= 10^9

// Example:

//     Input: nums = [1, 2, 3, 4, 5], target = 6
//     Output: [1, 3] (because nums[1] + nums[3] = 2 + 4 = 6)

// Follow-up: Can you solve this problem in O(n) time complexity?
func TestPairSumSortedLowerBound(t *testing.T) {
	nums := []int{1, 3}
	want := []int{0, 1}
	got := pairSumSorted(nums, 4)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Expected %v, but got %v", want, got)
	}
}

func TestPairSumSortedUpperBound(t *testing.T) {
	nums := make([]int, 10000)
	for i := range [10000]int{} {
		nums[i] = i + 1
	}
	target := nums[0] + nums[9999]
	want := []int{0, 9999}
	got := pairSumSorted(nums, target)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Expected %v, but got %v", want, got)
	}
}
