package main

import (
	"reflect"
	"testing"
)

// Given an array of integers nums and an integer target,
//  return indices of the two numbers such that they add up to target.

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

func TestTwoSumLowerBound(t *testing.T) {
	nums := []int{1, 2}
	target := 3
	expected := []int{0, 1}
	result := twoSum(nums, target)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestTwoSumUpperBound(t *testing.T) {
	nums := make([]int, 10000)
	for i := range nums {
		nums[i] = i + 1
	}
	target := 19999
	expected := []int{9998, 9999}
	result := twoSum(nums, target)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestTwoSumNegativeNumbers(t *testing.T) {
	nums := []int{-3, 4, 3, 90}
	target := 0
	expected := []int{0, 2}
	result := twoSum(nums, target)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestTwoSumZeroes(t *testing.T) {
	nums := []int{0, 4, 3, 0}
	target := 0
	expected := []int{0, 3}
	result := twoSum(nums, target)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestTwoSumNoSolution(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	target := 8
	expected := []int{}
	result := twoSum(nums, target)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestTwoSumDuplicates(t *testing.T) {
	nums := []int{3, 3}
	target := 6
	expected := []int{0, 1}
	result := twoSum(nums, target)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
