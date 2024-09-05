package main

import (
	"reflect"
	"testing"
)

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// You can return the answer in any order.

// Example 1:

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

// Example 2:

// Input: nums = [3,2,4], target = 6
// Output: [1,2]

// Example 3:

// Input: nums = [3,3], target = 6
// Output: [0,1]

// Constraints:

//     2 <= nums.length <= 104
//     -109 <= nums[i] <= 109
//     -109 <= target <= 109
//     Only one valid answer exists.

// Follow-up: Can you come up with an algorithm that is less than O(n2) time complexity?

func TestTwoSum(t *testing.T) {
	// First test case
	nums := []int{2, 7, 11, 15}
	target := 9
	expected := []int{0, 1}
	result := twoSum(nums, target)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	// Second test case
	nums = []int{3, 2, 4}
	target = 6
	expected = []int{1, 2}
	result = twoSum(nums, target)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	// Third test case
	nums = []int{3, 3}
	target = 6
	expected = []int{0, 1}
	result = twoSum(nums, target)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	// Fourth test case
	nums = []int{1, 2, 3, 4, 5}
	target = 9
	expected = []int{3, 4}
	result = twoSum(nums, target)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	// Fifth test case (Boundary: Empty array)
	nums = []int{}
	target = 9
	expected = []int{}
	result = twoSum(nums, target)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	// Sixth test case (Boundary: Single element array)
	nums = []int{9}
	target = 9
	expected = []int{}
	result = twoSum(nums, target)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	// Seventh test case (Extreme: Large array)
	nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target = 19
	expected = []int{8, 9}
	result = twoSum(nums, target)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	// Eighth test case (Extreme: Negative numbers)
	nums = []int{-2, -7, -11, -15}
	target = -9
	expected = []int{0, 1}
	result = twoSum(nums, target)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
