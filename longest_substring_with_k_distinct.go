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

// Return the best solution after processing the entire input.
// func lengthOfLongestSubstringKDistinct(s string, k int) int {
// 	if k == 0 {
// 		return 0
// 	}

// 	maxLength := 0

// 	// Outer loop for starting point of each substring
// 	for i := range s {
// 		seen := make(map[byte]int) // Track the frequency of each character in the current substring
// 		count := 0                 // Track the length of the current substring
// 		distinct := 0              // Track the number of distinct characters

// 		// Inner loop to expand the substring
// 		for j := i; j < len(s); j++ {
// 			if seen[s[j]] == 0 {
// 				distinct++ // Increment distinct count for a new character
// 			}
// 			seen[s[j]]++ // Increment the frequency of the character

// 			// If distinct characters exceed k, break the inner loop
// 			if distinct > k {
// 				break
// 			}

// 			count++ // Increase the length of the current valid substring
// 		}

// 		// Update the maximum length of valid substring
// 		if count > maxLength {
// 			maxLength = count
// 		}
// 	}

// 	return maxLength
// }

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	n := len(s)
	if n == 0 || k == 0 {
		return 0
	}

	maxLength := 0

	// Iterate over all starting points of substrings
	for i := range s {
		// Use a slice to count occurrences of characters
		count := make([]int, 256) // Assuming ASCII characters
		distinctCount := 0

		// Iterate over all ending points of substrings
		for j := i; j < n; j++ {
			char := s[j]
			count[char]++

			// If this character was not seen before, increase distinct count
			if count[char] == 1 {
				distinctCount++
			}

			// If we exceed k distinct characters, break out of the loop
			if distinctCount > k {
				break
			}

			// Update maxLength if the current substring is valid
			if j-i+1 > maxLength {
				maxLength = j - i + 1
			}
		}
	}

	return maxLength
}
