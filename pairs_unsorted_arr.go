package main

func flatten(arr [][]int) []int {
	var flat []int
	for _, subArr := range arr {
		flat = append(flat, subArr...)
	}
	return flat
}

func contains(arr []int, target int) bool {
	for _, elem := range arr {
		if elem == target {
			return true
		}
	}
	return false
}

func pairsUnsortedArr(nums []int, target int) [][]int {
	result := [][]int{}
	used := make(map[int]bool) // Map to keep track of used indices

	for i, num := range nums {
		if used[i] {
			continue // Skip if the index is already used
		}
		for j := i + 1; j < len(nums); j++ {
			if used[j] {
				continue // Skip if the index is already used
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
