package main

import "log"

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
// func lengthOfLongestSubstring(s string) int {
// 	n := len(s)
// 	if n == 0 {
// 		return 0
// 	}

// 	maxLength := 0

// 	for i := range s { // Using range to iterate over indices
// 		seen := make(map[byte]bool)
// 		count := 0

// 		for j := i; j < n; j++ {
// 			if _, exists := seen[s[j]]; exists {
// 				break
// 			}
// 			seen[s[j]] = true
// 			count++
// 		}

// 		if count > maxLength {
// 			maxLength = count
// 		}
// 	}

// 	return maxLength
// }

// func lengthOfLongestSubstring(s string) int {
// 	n := len(s)
// 	if n == 0 {
// 		return 0
// 	}

// 	maxLength := 0

// 	// Iterate over all starting points of substrings
// 	for i := range s {
// 		// Create a slice to track characters in the current substring
// 		seen := make([]bool, 256) // Assuming ASCII characters
// 		count := 0

// 		// Iterate over all ending points of substrings
// 		for j := i; j < n; j++ {
// 			char := s[j]

// 			// If the character has already been seen, break out of the loop
// 			if seen[char] {
// 				break
// 			}

// 			// Mark the character as seen
// 			seen[char] = true
// 			count++ // Increase the count of unique characters
// 		}

// 		// Update maxLength if the current substring is valid
// 		if count > maxLength {
// 			maxLength = count
// 		}
// 	}

// 	return maxLength
// }

// Let's break this down into a **higher-level abstraction** that can help you understand the steps clearly, without focusing too much on the implementation details:

// ### **Abstract Layer for Brute Force Solution (Without Hash Map)**

// #### **Step 1: State Initialization**
//    - **Key Concept**: Start by **tracking the best result** (`maxLength`), which will store the longest substring without repeating characters. We also initialize variables to **track current progress** like `count` for the length of the current substring.
//    - **Abstraction**: This is your **state-tracking** mechanism—track the best and current state to make decisions later.

// #### **Step 2: Sequential Exploration (Outer Loop)**
//    - **Key Concept**: Start from each position `i` in the string and explore substrings starting from that index. This loop allows us to explore each potential starting point of a substring.
//    - **Abstraction**: This represents the concept of **exploring possibilities** starting from different points, like exploring different strategies or approaches from various starting conditions.

// #### **Step 3: Expanding Substrings (Inner Loop)**
//    - **Key Concept**: For each starting point `i`, explore the possible substrings that can be formed by expanding the substring to include more characters. For each expanded substring, check if it satisfies the condition (no duplicate characters).
//    - **Abstraction**: This represents **gradual exploration** or **expansion of possibilities**. In real-life terms, this is like starting small and continuously testing bigger scenarios or decisions.

// #### **Step 4: Manual Check for Duplicates**
//    - **Key Concept**: For every expanded substring, check if the new character is a **duplicate** of any character already in the substring. If it is, stop expanding this substring.
//    - **Abstraction**: This is your **constraint management**. You need to stop or change direction when a certain condition is met (duplicate character found). It's similar to stopping a decision-making process when the current path violates a rule.

// #### **Step 5: Update State**
//    - **Key Concept**: If the substring is valid (i.e., no duplicates were found), increment the count of valid characters in this substring. Track the maximum valid substring length by updating the **state** of the best result (`maxLength`) if the current one is longer.
//    - **Abstraction**: This step reflects **continuous updates to the best solution** as you explore more possibilities. Each valid substring found may improve the current best solution, and you want to capture that.

// #### **Step 6: Return the Best Result**
//    - **Key Concept**: Once all possibilities have been explored and expanded from every starting point, return the maximum length of the valid substring that was found.
//    - **Abstraction**: This is your **final decision point**. After exploring all options, you arrive at the best possible outcome and return the result.

// ---

// ### **High-Level Recipe (Abstraction for Any Similar Problem)**:

// 1. **Initialize State**: Set up tracking mechanisms for the best result and current progress.
//    - **Life Example**: If you're trying to balance work and personal tasks in a week, you initialize with a tracker for "best productivity balance."

// 2. **Explore Starting Points**: Try out different possible starting points.
//    - **Life Example**: In planning a trip, start considering routes from different cities as the origin.

// 3. **Expand Gradually**: For each starting point, progressively expand the solution space, testing how far you can go before hitting a constraint.
//    - **Life Example**: For a new diet, gradually add new types of food and see how your body reacts until you find something that works or breaks the balance.

// 4. **Check Constraints**: Continuously check whether the expanded solution still satisfies the problem constraints (e.g., no duplicates, within budget, not too far).
//    - **Life Example**: As you plan tasks for the week, check if adding another task exceeds your available time or energy.

// 5. **Update Best Solution**: Whenever you find a valid solution, compare it with the best one so far and update if it's better.
//    - **Life Example**: Each time you experiment with a new work/life balance schedule, keep track of the most productive one and update if a new plan works better.

// 6. **Return Final Result**: Once all possibilities have been explored, select the best solution and apply it.
//    - **Life Example**: After trying various schedules, routes, or budgets, you settle on the one that works best and implement it.

// ---

// ### **Mapping to Your Code:**

// 1. **Initialize State**:
//    - Variables: `maxLength = 0`, `count = 0`
// 2. **Outer Loop** (Explore Starting Points):
//    - Loop: `for i := range s`
// 3. **Inner Loop** (Expand Substrings):
//    - Loop: `for j := i; j < len(s)`
// 4. **Manual Check for Duplicates**:
//    - Condition: Check `s[j]` exists in the substring `s[i:j]`.
// 5. **Update Best Solution**:
//    - Update: Compare `count` with `maxLength`.
// 6. **Return Final Result**:
//    - Return: `return maxLength`

// ### **Generalized Problem-Solving Using This Pattern**:
// You can now apply this high-level abstraction pattern to any problem that involves:
// - Exploring starting points.
// - Gradually expanding or growing a solution space.
// - Checking for constraints or limits.
// - Updating the best solution as you go.

// This pattern is not just about programming but also applies to decision-making, event planning, budgeting, or any other task that requires finding the best solution among many possibilities.
func lengthOfLongestSubstring(s string) int {
	log.Println("Input: ", s)
	start := 0
	hash := make(map[byte]int)
	maxLength := 0

	for end := range s {
		// I understand. I'll walk you through the process of forming a higher abstraction layer for this problem, providing my own analysis and answers. Then you can digest this information and ask any follow-up questions you may have.

		// 1. Problem Description:
		// The `val >= start` check in the longest substring without repeating characters problem ensures that we only consider characters within the current sliding window. It prevents us from incorrectly updating our window based on character occurrences that are no longer relevant.

		// 2. Core Elements:
		// The essential elements in this problem are:
		// - The sliding window (defined by `start` and `end` pointers)
		// - The hash map (storing the last seen position of each character)
		// - The concept of relevance (whether a character's previous occurrence is still within our current window)

		// 3. Analogies:
		// This situation is analogous to maintaining a guest list for a venue with limited capacity:
		// - As new guests arrive (new characters), you add them to the list.
		// - When the venue is full and a new guest arrives, you need to remove someone.
		// - You want to remove the person who's been there the longest, but only if they're still in the venue (still in our current window).

		// 4. Abstract Representation:
		// At a more abstract level, this check is about maintaining the integrity of a dynamic, bounded set of elements. It ensures that decisions about the set's composition are made based only on currently relevant information.

		// 5. Different Levels of Abstraction:
		// - Data Structures: This concept relates to maintaining the validity of cached or memoized data.
		// - Algorithms: It's a form of "staleness check" in algorithms that process streaming data.
		// - Systems Design: Similar checks are used in distributed systems to ensure operations are performed on up-to-date data.

		// 6. Applying Abstracted Understanding:
		// By understanding this as a "relevance check" or "staleness prevention mechanism", we can see why it's crucial in our original algorithm:
		// - It prevents "stale" data (outdated character positions) from influencing our current decision-making.
		// - It ensures the integrity of our sliding window by only considering truly relevant information.
		// - It allows us to correctly handle cases where a character repeats after a long interval, maintaining the accuracy of our substring length calculation.

		// This higher-level abstraction helps us see that the `val >= start` check is not just a quirk of this specific algorithm, but a fundamental concept in maintaining the integrity of dynamic, bounded data sets. It's a pattern that can be applied in various scenarios where we need to ensure we're working with current, relevant information.
		if val, ok := hash[s[end]]; ok && val >= start {
			log.Println("Current char: ", string(s[end]))
			log.Println("start BEFORE: ", start)
			log.Println("start is gonna be: ", val, " + 1")
			start = val + 1
			log.Println("start AFTER: ", start)
		}
		hash[s[end]] = end
		log.Println("Current hash: ", hash)

		if end-start+1 > maxLength {
			log.Println("maxLengh is gonna be re-assigned cuz of end - start + 1 > maxLengh: ", end, " - ", start, " + ", "1", " > ", maxLength)
			maxLength = end - start + 1
		}
	}
	return maxLength
}

// func lengthOfLongestSubstring(s string) int {
// 	n := len(s)
// 	if n == 0 {
// 		return 0
// 	}

// 	maxLength := 0

// 	for i := range s {
// 		count := 0

// 		for j := i; j < n; j++ {
// 			// Check if s[j] exists in the substring s[i:j]
// 			isDuplicate := false
// 			for k := i; k < j; k++ {
// 				if s[k] == s[j] {
// 					isDuplicate = true
// 					break
// 				}
// 			}

// 			// If duplicate character is found, break the inner loop
// 			if isDuplicate {
// 				break
// 			}

// 			count++ // Increment the length of current valid substring
// 		}

// 		// Update the maximum length found so far
// 		if count > maxLength {
// 			maxLength = count
// 		}
// 	}

// 	return maxLength
// }

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
