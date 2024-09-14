package main

//import "math"

func maximumSubarraySum(nums []int, k int) int64 {
	maxSum := 0
	windowSum := 0
	n := len(nums)
	window := make(map[int]bool)

	left := 0

	for right := 0; right < n; right++ {
		for window[nums[right]] {
			delete(window, nums[left])
			windowSum -= nums[left]
			left++
		}

		window[nums[right]] = true
		windowSum += nums[right]

		if right-left+1 == k {
			if windowSum > maxSum {
				maxSum = windowSum
			}
			delete(window, nums[left])
			windowSum -= nums[left]
			left++
		}
	}
	return int64(maxSum)
}
