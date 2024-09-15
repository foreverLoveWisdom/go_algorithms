package main

import (
	"testing"
)

// Given an array of positive integers nums and a positive integer target, return the minimal length of a
// subarray
// whose sum is greater than or equal to target. If there is no such subarray, return 0 instead.

// Example 1:

// Input: target = 7, nums = [2,3,1,2,4,3]
// Output: 2
// Explanation: The subarray [4,3] has the minimal length under the problem constraint.

// Example 2:

// Input: target = 4, nums = [1,4,4]
// Output: 1

// Example 3:

// Input: target = 11, nums = [1,1,1,1,1,1,1,1]
// Output: 0

// Constraints:

//     1 <= target <= 109
//     1 <= nums.length <= 105
//     1 <= nums[i] <= 104

// Follow up: If you have figured out the O(n) solution, try coding another solution of which the time complexity is O(n log(n)).

func TestMinimumSubarraySum_Example0(t *testing.T) {
	target := 15
	nums := []int{1, 2, 3, 4, 5}
	expected := 5

	result := minSubArrayLen(target, nums)
	if result != expected {
		t.Errorf("TestMinimumSubarraySum_Example0 failed, expected %d, got %d", expected, result)
	}
}

// TestCase 1: Normal case with a valid subarray of minimal length
func TestMinimumSubarraySum_Example1(t *testing.T) {
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	expected := 2 // The subarray [4,3] has a sum of 7

	result := minSubArrayLen(target, nums)
	if result != expected {
		t.Errorf("TestMinimumSubarraySum_Example1 failed, expected %d, got %d", expected, result)
	}
}

// TestCase 2: Single element can be the answer
func TestMinimumSubarraySum_Example2(t *testing.T) {
	target := 4
	nums := []int{1, 4, 4}
	expected := 1 // The subarray [4] satisfies the condition

	result := minSubArrayLen(target, nums)
	if result != expected {
		t.Errorf("TestMinimumSubarraySum_Example2 failed, expected %d, got %d", expected, result)
	}
}

// TestCase 3: No valid subarray exists
func TestMinimumSubarraySum_Example3(t *testing.T) {
	target := 11
	nums := []int{1, 1, 1, 1, 1, 1, 1, 1}
	expected := 0 // No subarray sum >= 11

	result := minSubArrayLen(target, nums)
	if result != expected {
		t.Errorf("TestMinimumSubarraySum_Example3 failed, expected %d, got %d", expected, result)
	}
}

// TestCase 4: Single element equals target
func TestMinimumSubarraySum_SingleElementEqualsTarget(t *testing.T) {
	target := 10
	nums := []int{10}
	expected := 1 // The single element 10 meets the target

	result := minSubArrayLen(target, nums)
	if result != expected {
		t.Errorf("TestMinimumSubarraySum_SingleElementEqualsTarget failed, expected %d, got %d", expected, result)
	}
}

// TestCase 5: Multiple elements, no valid subarray
func TestMinimumSubarraySum_NoValidSubarray(t *testing.T) {
	target := 15
	nums := []int{1, 2, 3, 4, 5}
	expected := 5 // No subarray sum >= 15

	result := minSubArrayLen(target, nums)
	if result != expected {
		t.Errorf("TestMinimumSubarraySum_NoValidSubarray failed, expected %d, got %d", expected, result)
	}
}

// TestCase 6: Exact match at the end of the array
func TestMinimumSubarraySum_MatchAtEnd(t *testing.T) {
	target := 6
	nums := []int{1, 1, 1, 1, 1, 6}
	expected := 1 // The last element 6 meets the target

	result := minSubArrayLen(target, nums)
	if result != expected {
		t.Errorf("TestMinimumSubarraySum_MatchAtEnd failed, expected %d, got %d", expected, result)
	}
}

// TestCase 7: Large array boundary case
func TestMinimumSubarraySum_LargeArray(t *testing.T) {
	target := 100
	nums := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		nums[i] = 1
	}
	expected := 100 // The sum of the first 100 elements is 100

	result := minSubArrayLen(target, nums)
	if result != expected {
		t.Errorf("TestMinimumSubarraySum_LargeArray failed, expected %d, got %d", expected, result)
	}
}

// TestCase 8: Minimal array
func TestMinimumSubarraySum_MinimalArray(t *testing.T) {
	target := 1
	nums := []int{1}
	expected := 1 // Single element meets the target

	result := minSubArrayLen(target, nums)
	if result != expected {
		t.Errorf("TestMinimumSubarraySum_MinimalArray failed, expected %d, got %d", expected, result)
	}
}
