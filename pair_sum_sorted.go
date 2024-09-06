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
	// So the 2 pointers approach is:
	// 1. Initialize 2 pointers, one at the start of the array and the other at the end of the array.
	// 2. If the sum of the 2 numbers at the pointers is equal to the target, return the indices of the 2 numbers.
	// 3. If the sum is less than the target, move the first pointer to the right.
	// 4. If the sum is greater than the target, move the second pointer to the left.
	// 5. Repeat steps 2-4 until the pointers meet.
	// The benefit of this approach is that it has a time complexity of O(n) and a space complexity of O(1).
	// Because we are only using 2 pointers to traverse the array, the space complexity is constant.
	// Compared to the hash map approach, this approach is more efficient in terms of space complexity.
	// However, the array must be sorted for this approach to work.
	// If not, we can sort the array first, which would increase the time complexity to O(n log n) due to the sorting step.
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
