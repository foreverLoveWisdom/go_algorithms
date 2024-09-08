package main

func pairsUnsortedArr(nums []int, target int) [][]int {
	result := [][]int{}
	used := make(map[int]bool) // Map to keep track of used indices

	for i, num := range nums {
		if used[i] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if used[j] {
				continue
			}
			if num+nums[j] == target {
				result = append(result, []int{i, j})
				used[i] = true
				used[j] = true
				break
			}
		}
	}
	return result
}
