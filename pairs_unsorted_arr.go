package main

func pairsUnsortedArr(nums []int, target int) [][]int {
	seen := make(map[int]int)  // Map to store the number and its index
	used := make(map[int]bool) // Map to track used indices
	result := [][]int{}
	for i, num := range nums {
		complement := target - num
		if j, found := seen[complement]; found && !used[j] {
			// We found a complement and the complement's index hasn't been used
			result = append(result, []int{j, i})
			used[j] = true
		} else {
			seen[num] = i
		}
	}
	return result
}
