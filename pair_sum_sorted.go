package main

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

func pairSumSorted(nums []int, target int) []int {
	firstPointer := 0
	secondPointer := len(nums) - 1

	for firstPointer < secondPointer {
		sum := nums[firstPointer] + nums[secondPointer]
		switch {
		case sum == target:
			return []int{firstPointer, secondPointer}
		case sum < target:
			firstPointer++
		default:
			secondPointer--
		}
	}

	return []int{}
}
