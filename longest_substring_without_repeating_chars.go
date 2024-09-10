package main

// Abstraction
// State Tracking:

//     Track the best result (e.g., maxLength) or other relevant metrics.
//     Use this to keep track of the current state (e.g., sums, lengths, counts).

// Looping through Input:

//     Outer loop for sliding starting points (for i := 0; i < len(s); i++).
//     Inner loop for expanding substrings (for j := i; j < len(s); j++).

// Tracking Seen Elements (Hash Map):

//     Use a hash map or set to track elements you’ve encountered, allowing quick lookup to avoid duplicates.

// Condition-Based Checking:

//     Use conditions to break or continue loops when specific criteria are met
// (e.g., duplicate detection, subarray length constraints).

// State Updates:

//     Continuously update the state (e.g., hash map, counters) as you explore new elements in your window or subset.

// Tracking the Optimal Solution:

//     Continuously update the best result as you process each new possibility (maxLength = count).

// Returning the Final Result:

// Return the best solution after processing the entire input
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	maxLength := 0

	for i := range s { // Using range to iterate over indices
		seen := make(map[byte]struct{})
		count := 0

		for j := i; j < n; j++ {
			if _, exists := seen[s[j]]; exists {
				break
			}
			seen[s[j]] = struct{}{}
			count++
		}

		if count > maxLength {
			maxLength = count
		}
	}

	return maxLength
}

// Key Problems to Focus On (20% Effort, 80% Results):

//     Longest Substring with At Most K Distinct Characters
//         Why: This problem teaches you sliding window techniques and hash map usage to manage constraints on distinct elements, which is a pattern in many problems (subarrays, substrings, etc.).

//     Smallest Subarray with Sum Greater Than K
//         Why: This problem helps you master sliding windows while adjusting window size dynamically, which can be applied to many subarray and window problems.

//     Longest Palindromic Substring
//         Why: Expanding around the center and handling different length palindromes gives you insights into two-pointer approaches, which apply to many other problems.

//     Subarray with Equal Number of 0s and 1s
//         Why: This problem introduces the concept of using cumulative sums with hash maps, a powerful technique to solve array-based problems efficiently.

// Why These Problems?

//     Common Patterns: These problems use key techniques like sliding windows, two pointers, and hash maps, which are widely applicable across a range of algorithmic challenges.
//     Skill Transfer: Mastering these problems will enable you to tackle variations with similar structures, covering the majority of medium difficulty problems involving subarrays, substrings, or constraints on elements.

// Less Priority Problems (80% Effort, 20% Results):

//     Maximum Product of Subarray: Although useful, it's a specific problem that doesn’t generalize as much as the others.
//     Longest Subarray of Ones After Deleting One Element: This is more niche and does not offer as broad applicability as the problems listed above.

// By focusing on the top 4 problems, you can maximize your learning outcomes and transfer the patterns to many other related problems with less effort!
