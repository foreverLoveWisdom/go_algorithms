package main

func maximumSubarraySum(nums []int, k int) int64 {
	maxSum := 0
	windowSum := 0
	window := make(map[int]bool)

	left := 0

	for right := range nums {
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
