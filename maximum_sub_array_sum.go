package main

import "log"

// No need for pointers for the map because maps are reference types in Go
func shrinkWindow(window map[int]bool, nums []int, left *int, windowSum *int) {
	delete(window, nums[*left])
	*windowSum -= nums[*left]
	(*left)++
}

func maximumSubarraySum(nums []int, k int) int64 {
	window := make(map[int]bool)
	windowSum := 0
	maxSum := 0
	left := 0

	for right := range nums {
		log.Println("nums[right] is: ", nums[right])
		for window[nums[right]] {
			shrinkWindow(window, nums, &left, &windowSum)
		}

		window[nums[right]] = true
		windowSum += nums[right]

		if right-left+1 == k {
			if windowSum > maxSum {
				maxSum = windowSum
			}
			shrinkWindow(window, nums, &left, &windowSum)
		}
	}

	return int64(maxSum)
}

// Let's analyze the caching mechanisms implemented in this solution from a high-level abstraction.

// ### **Key Concept of Caching Mechanism**

// In a broader sense, caching refers to temporarily storing information for fast access later. In the context of this algorithm, there are multiple levels where caching, or temporary storage of information for reuse, occurs:

// ### **1. Window Caching (Element Frequency Cache)**

// The **first caching mechanism** happens with the `window` map, which stores the presence of each element within the current sliding window:

// - **Purpose**: The `window` map acts as a cache to keep track of which elements are in the current window, enabling the quick detection of duplicates. This avoids recalculating which elements exist in the window at each step.

// - **Benefits**:
//   - By caching the elements in the `window`, we avoid recalculating whether an element already exists in the current window.
//   - Accessing and updating the map (insertion, deletion) happens in O(1) time on average, making it an efficient way to manage the uniqueness constraint.

// - **Caching Level**: This is a local, in-memory cache for window state management.

// ### **2. Window Sum Caching**

// The **second caching mechanism** is related to the **sum of the current window's elements**, managed by `windowSum`.

// - **Purpose**: `windowSum` caches the sum of the current sliding window so that the sum can be updated incrementally as elements are added or removed. This eliminates the need to recalculate the sum of the entire window every time.

// - **Benefits**:
//   - Instead of recalculating the sum of elements from scratch each time the window changes, you only update the sum by adding or subtracting individual elements when moving the window's boundaries (`left` or `right` pointers).
//   - This results in O(1) updates to the sum as the window slides.

// - **Caching Level**: This is a local, incremental cache that stores the running total of the current window's sum.

// ### **3. Maximum Sum Caching**

// The **third caching mechanism** is the **`maxSum`** variable:

// - **Purpose**: This caches the maximum sum found so far as the window slides across the array. It keeps track of the global maximum, eliminating the need to revisit previous windows.

// - **Benefits**:
//   - Once you calculate the sum of a valid window (size `k` with unique elements), the `maxSum` stores this value and compares it with subsequent windows. You don’t need to reevaluate previous windows.

// - **Caching Level**: This is a high-level cache that stores the result of the maximum sum comparison, making sure we only keep the largest valid sum.

// ---

// ### **Summary of Caching Mechanisms**

// 1. **Window Caching (Map)**:
//    - Caches the elements present in the current window to manage duplicates.
//    - Avoids recalculating the presence of elements in the window.

// 2. **Window Sum Caching**:
//    - Caches the sum of the current window’s elements and updates it incrementally.
//    - Avoids recalculating the sum of the entire window each time.

// 3. **Maximum Sum Caching**:
//    - Caches the maximum valid window sum found so far.
//    - Ensures we don’t recompute results for earlier windows.

// ### **Caching Levels (High-Level View)**:

// - **Level 1**: **In-memory data structure caching (`window` map)**. This cache maintains the state of the current window elements for quick lookups and updates.
// - **Level 2**: **State caching (`windowSum`)**. This cache stores the sum of the current window's elements and is updated incrementally as the window slides.
// - **Level 3**: **Global result caching (`maxSum`)**. This cache stores the maximum sum found across all valid subarrays.

// These caching mechanisms contribute to the overall efficiency of the solution by preventing unnecessary recalculations and ensuring that the sliding window operations (expansion, contraction, and sum computation) happen in **O(1)** time per step.

// ---

// ### **Improvements and Considerations:**

// The algorithm already uses efficient caching techniques to manage the sliding window, but further improvements can include:

// - **Memory Optimization**: If the input size is extremely large, the memory usage of the `window` map can be optimized by using a fixed-size sliding window of capacity `k`. However, in most practical cases, the current approach is sufficient.

// - **Reducing Logging**: The logging statements (e.g., `log.Println`) could be removed or limited in production code to reduce unnecessary I/O, improving performance.

// This solution effectively uses caching techniques to achieve an optimal time complexity of **O(n)** with respect to both time and space usage.

// Let me know if you have further questions on this or would like more details!
