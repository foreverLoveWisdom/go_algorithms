package main

import "log"

func pairsUnsortedArr(nums []int, target int) [][]int {
	log.Println("Nums:", nums)
	result := [][]int{}
	indexMap := map[int]int{}
	for i, num := range nums {
		log.Println("Current indexMap:", indexMap)
		if _, ok := indexMap[target-num]; ok {
			log.Println("Current num:", num)
			log.Println("Matched pair found:", num, target-num)
			result = append(result, []int{indexMap[target-num], i})
			log.Println("Current result:", result)
		}
		indexMap[num] = i
	}

	return result
}
