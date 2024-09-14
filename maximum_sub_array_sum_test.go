package main

import (
	"testing"
)

// You are given an integer array nums and an integer k.
// Find the maximum subarray sum of all the subarrays of nums that meet the following conditions:

//     The length of the subarray is k, and
//     All the elements of the subarray are distinct.

// Return the maximum subarray sum of all the subarrays that meet the conditions.
//  If no subarray meets the conditions, return 0.

// A subarray is a contiguous non-empty sequence of elements within an array.

// Example 1:

// Input: nums = [1,5,4,2,9,9,9], k = 3
// Output: 15
// Explanation: The subarrays of nums with length 3 are:
// - [1,5,4] which meets the requirements and has a sum of 10.
// - [5,4,2] which meets the requirements and has a sum of 11.
// - [4,2,9] which meets the requirements and has a sum of 15.
// - [2,9,9] which does not meet the requirements because the element 9 is repeated.
// - [9,9,9] which does not meet the requirements because the element 9 is repeated.
// We return 15 because it is the maximum subarray sum of all the subarrays that meet the conditions

// Example 2:

// Input: nums = [4,4,4], k = 3
// Output: 0
// Explanation: The subarrays of nums with length 3 are:
// - [4,4,4] which does not meet the requirements because the element 4 is repeated.
// We return 0 because no subarrays meet the conditions.

// Constraints:

//     1 <= k <= nums.length <= 105
//     1 <= nums[i] <= 105

// Function to be tested

func TestExample1(t *testing.T) {
	nums := []int{1, 5, 4, 2, 9, 9, 9}
	k := 3
	expected := int64(15)
	result := maximumSubarraySum(nums, k)
	if result != expected {
		t.Errorf("TestExample1 failed: expected %d, got %d", expected, result)
	}
}

func TestExample2(t *testing.T) {
	nums := []int{4, 4, 4}
	k := 3
	expected := int64(0)
	result := maximumSubarraySum(nums, k)
	if result != expected {
		t.Errorf("TestExample2 failed: expected %d, got %d", expected, result)
	}
}

func TestSingleElementValid(t *testing.T) {
	nums := []int{10}
	k := 1
	expected := int64(10)
	result := maximumSubarraySum(nums, k)
	if result != expected {
		t.Errorf("TestSingleElementValid failed: expected %d, got %d", expected, result)
	}
}

func TestSingleElementInvalid(t *testing.T) {
	nums := []int{10, 10}
	k := 2
	expected := int64(0)
	result := maximumSubarraySum(nums, k)
	if result != expected {
		t.Errorf("TestSingleElementInvalid failed: expected %d, got %d", expected, result)
	}
}

func TestAllDistinct(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	k := 2
	expected := int64(9) // Subarray [4, 5]
	result := maximumSubarraySum(nums, k)
	if result != expected {
		t.Errorf("TestAllDistinct failed: expected %d, got %d", expected, result)
	}
}

func TestAllSameElements(t *testing.T) {
	nums := []int{7, 7, 7, 7, 7}
	k := 1
	expected := int64(7)
	result := maximumSubarraySum(nums, k)
	if result != expected {
		t.Errorf("TestAllSameElements failed: expected %d, got %d", expected, result)
	}
}

func TestNoValidSubarray(t *testing.T) {
	nums := []int{1, 1, 1, 1}
	k := 2
	expected := int64(0)
	result := maximumSubarraySum(nums, k)
	if result != expected {
		t.Errorf("TestNoValidSubarray failed: expected %d, got %d", expected, result)
	}
}

func TestKEqualsArrayLengthValid(t *testing.T) {
	nums := []int{3, 1, 2, 4}
	k := 4
	expected := int64(10)
	result := maximumSubarraySum(nums, k)
	if result != expected {
		t.Errorf("TestKEqualsArrayLengthValid failed: expected %d, got %d", expected, result)
	}
}

func TestKEqualsArrayLengthInvalid(t *testing.T) {
	nums := []int{3, 1, 2, 3}
	k := 4
	expected := int64(0)
	result := maximumSubarraySum(nums, k)
	if result != expected {
		t.Errorf("TestKEqualsArrayLengthInvalid failed: expected %d, got %d", expected, result)
	}
}

func TestLargeKValue(t *testing.T) {
	nums := []int{1, 2, 3}
	k := 5
	expected := int64(0) // k > len(nums), so no valid subarray
	result := maximumSubarraySum(nums, k)
	if result != expected {
		t.Errorf("TestLargeKValue failed: expected %d, got %d", expected, result)
	}
}

func TestMultipleValidSubarrays(t *testing.T) {
	nums := []int{1, 2, 3, 2, 1, 4, 5}
	k := 3
	expected := int64(10) // Subarray [1, 4, 5]
	result := maximumSubarraySum(nums, k)
	if result != expected {
		t.Errorf("TestMultipleValidSubarrays failed: expected %d, got %d", expected, result)
	}
}

func TestZerosInArray(t *testing.T) {
	nums := []int{0, 0, 0, 0}
	k := 1
	expected := int64(0)
	result := maximumSubarraySum(nums, k)
	if result != expected {
		t.Errorf("TestZerosInArray failed: expected %d, got %d", expected, result)
	}
}

func TestMixedPositiveAndNegative(t *testing.T) {
	nums := []int{1, -2, 3, 4, -5, 6}
	k := 2
	expected := int64(7) // Subarray [3,4]
	result := maximumSubarraySum(nums, k)
	if result != expected {
		t.Errorf("TestMixedPositiveAndNegative failed: expected %d, got %d", expected, result)
	}
}

func TestMaximumConstraintsWithDistinct(t *testing.T) {
	nums := generateLargeDistinctArray(100000)
	k := 100000
	expected := int64(5000050000) // Sum of first 100000 natural numbers
	result := maximumSubarraySum(nums, k)
	if result != expected {
		t.Errorf("TestMaximumConstraintsWithDistinct failed: expected %d, got %d", expected, result)
	}
}

// Helper function to generate a large array with distinct elements.
func generateLargeDistinctArray(size int) []int {
	arr := make([]int, size)
	for i := range make([]struct{}, size) {
		arr[i] = i + 1
	}
	return arr
}
