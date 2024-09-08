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

// Test case 1: Basic scenario with maximum profit.
func maxProfit(prices []int) int {
	// they give array of integers, we need to find the maximum profit
	// the goal is to find the maximum: seems like i need to perform the minus
	// let's go with brute force first
	// can use nested loops
	// what are the actions I can perform here:
	// 1. loop through array?
	// 2. what are the states I need?
	// 3. what do I need to keep track?
	// 4. what are the conditions I need to check?

	prices
	hardCodedValue := 5
	return hardCodedValue
}
