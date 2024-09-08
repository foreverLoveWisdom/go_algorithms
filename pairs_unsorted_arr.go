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
	for i, num := range nums {
		flatNums := flatten(result)
		if contains(flatNums, i) {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if num+nums[j] == target {
				result = append(result, []int{i, j})
				break
			}
		}
	}
	return result
}
