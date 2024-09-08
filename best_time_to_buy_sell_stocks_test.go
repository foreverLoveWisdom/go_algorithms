package main

import (
	"testing"
)

// You are given an array prices where prices[i] is the price of a given stock on the ith day.

// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

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
func TestMaxProfit_SimpleProfit(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	expected := 5

	result := maxProfit(prices)

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

// Test case 2: No profit possible (prices decreasing).
// func TestMaxProfit_NoProfit(t *testing.T) {
// 	prices := []int{7, 6, 4, 3, 1}
// 	expected := 0

// 	result := maxProfit(prices)

// 	if result != expected {
// 		t.Errorf("Expected %d, but got %d", expected, result)
// 	}
// }

// // Test case 3: Single price point (no transaction possible).
// func TestMaxProfit_SinglePrice(t *testing.T) {
// 	prices := []int{5}
// 	expected := 0

// 	result := maxProfit(prices)

// 	if result != expected {
// 		t.Errorf("Expected %d, but got %d", expected, result)
// 	}
// }

// // Test case 4: Constant prices (no profit possible).
// func TestMaxProfit_ConstantPrices(t *testing.T) {
// 	prices := []int{3, 3, 3, 3, 3}
// 	expected := 0

// 	result := maxProfit(prices)

// 	if result != expected {
// 		t.Errorf("Expected %d, but got %d", expected, result)
// 	}
// }

// // Test case 5: Large array with random prices.
// func TestMaxProfit_LargeArray(t *testing.T) {
// 	prices := make([]int, 100000)
// 	for i := 0; i < 50000; i++ {
// 		prices[i] = 100000 - i // Decreasing
// 	}
// 	for i := 50000; i < 100000; i++ {
// 		prices[i] = i // Increasing
// 	}
// 	expected := 99999 - 50000 // The maximum profit can be achieved by buying at index 49999 and selling at index 99999

// 	result := maxProfit(prices)

// 	if result != expected {
// 		t.Errorf("Expected %d, but got %d", expected, result)
// 	}
// }

// // Test case 6: Edge case with very low prices and small array.
// func TestMaxProfit_LowPrices(t *testing.T) {
// 	prices := []int{0, 1, 0, 1}
// 	expected := 1

// 	result := maxProfit(prices)

// 	if result != expected {
// 		t.Errorf("Expected %d, but got %d", expected, result)
// 	}
// }

// // Test case 7: Edge case with the maximum allowed prices.
// func TestMaxProfit_HighPrices(t *testing.T) {
// 	prices := []int{104, 103, 102, 101, 100}
// 	expected := 0

// 	result := maxProfit(prices)

// 	if result != expected {
// 		t.Errorf("Expected %d, but got %d", expected, result)
// 	}
// }
