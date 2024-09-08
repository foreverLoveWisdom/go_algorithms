package main

// You are given an array prices where prices[i] is the price of a given stock on the ith day.

// You want to maximize your profit by choosing a single day to buy one stock and
//  choosing a different day in the future to sell that stock.

// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

// Example 1:

// Input: prices = [7,1,5,3,6,4]
// Output: 5
// Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
// Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.

// Example 2:

// Input: prices = [7,6,4,3,1]
// Output: 0
// Explanation: In this case, no transactions are done and the max profit = 0.

// Constraints:

//     1 <= prices.length <= 105
//     0 <= prices[i] <= 104

// they give array of integers, we need to find the maximum profit
// the goal is to find the maximum: seems like i need to perform the minus
// let's go with brute force first
// can use nested loops
// what are the actions I can perform here:
// 1. loop through array?
// 2. what are the states I need?
// 3. what do I need to keep track?
// 4. what are the conditions I need to check?
// 5. Start with easiest step that lead to working solution
// Current TC is O(n^2)
// Current SC is O(1).
// func maxProfit(prices []int) int {
// 	count := len(prices)
// 	invalidPriceLength := 2
// 	if count < invalidPriceLength {
// 		return 0 // Not enough prices to make a profit
// 	}

// 	maxProfit := 0

// 	log.Println("Current prices: ", prices)
// 	for i := 0; i < count; i++ {
// 		for j := i + 1; j < count; j++ {
// 			// Calculate the profit from buying at index i and selling at index j
// 			if prices[i] < prices[j] {
// 				profit := prices[j] - prices[i]
// 				if profit > maxProfit {
// 					maxProfit = profit
// 				}
// 			}
// 		}
// 	}

// 	return maxProfit
// }

func maxProfit(prices []int) int {
	firstPointer := 0
	secondPointer := len(prices) - 1
	maxValue := 0
	for firstPointer < secondPointer {
		if prices[firstPointer] > prices[secondPointer] {
			firstPointer++
			continue
		}

		profit := prices[secondPointer] - prices[firstPointer]
		if profit > maxValue {
			maxValue = profit
		}
		secondPointer--
	}
	return maxValue
}
