package main

// Given an integer array nums,
//  return true if any value appears at least twice in the array, and return false if every element is distinct.

// Example 1:

// Input: nums = [1,2,3,1]

// Output: true

// Explanation:

// The element 1 occurs at the indices 0 and 3.

// Example 2:

// Input: nums = [1,2,3,4]

// Output: false

// Explanation:

// All elements are distinct.

// Example 3:

// Input: nums = [1,1,1,3,3,4,3,2,4,2]

// Output: true

// Constraints:

//     1 <= nums.length <= 105
//     -109 <= nums[i] <= 109

func containsDuplicate(nums []int) bool {
	for i, num := range nums {
		for j := i + 1; j < len(nums); j++ {
			if num == nums[j] {
				return true
			}
		}
	}
	return false
}

// Great! Let's go step by step to abstract the problem of finding duplicates in an array. This will help you create a higher-level understanding that you can apply to other problems as well. I'll guide you with questions and explanations along the way.

// ### 1. **Describe the problem in detail.**

// You are trying to solve the problem of checking whether an array contains any duplicates. The current solution involves comparing each element with every other element, which works, but it’s not efficient as it checks all possible pairs.

// **Questions:**
// - How do you currently approach finding duplicates?
// - What are the specific challenges or inefficiencies in your current solution?

// **Explanation:** In this step, we describe the problem in specific terms. Your current solution works, but it has a time complexity of **O(n²)** due to the nested loop. This is acceptable for small arrays but becomes inefficient for larger ones. The challenge here is to improve efficiency.

// ---

// ### 2. **Identify the core elements driving the situation.**

// Now, let's extract the key elements that define this problem.

// - **Core Elements**:
//   - **Input**: A list of numbers (`nums[]`).
//   - **Goal**: Determine whether there are duplicates in the list.
//   - **Constraints**: Efficiency matters as the list size increases.
//   - **Current approach**: A brute force method using two nested loops.

// **Questions:**
// - Can you list the inputs, outputs, and performance limitations of your current solution?
// - Why is reducing time complexity important in this context?

// **Explanation:** By breaking down the core elements, you recognize that the efficiency issue (due to nested loops) is the primary driver. Your goal is to identify a better way to find duplicates, focusing on performance.

// ---

// ### 3. **Find analogies with other situations you might be familiar with.**

// Let’s find analogies that might help you understand this problem better.

// - **Analogy**: Imagine you are checking a group of people entering a room and trying to remember whether you’ve already seen someone. If you try to remember each person by looking at everyone else again, that takes time. Instead, you could use a checklist to mark people you’ve seen. Each new person can be quickly checked against this list.

// **Questions:**
// - Does this analogy make sense in terms of reducing the effort for each comparison?
// - Can you think of another situation where keeping track of things already seen helps you avoid repetition?

// **Explanation:** By drawing analogies with real-life scenarios (e.g., checking people entering a room), you can see that keeping track of past information (like a checklist) allows you to avoid unnecessary rechecking, similar to avoiding repeated comparisons in your array.

// ---

// ### 4. **Create a more abstract representation of the problem.**

// Now, let’s abstract the problem away from specific code details and focus on its fundamental structure.

// - **Abstract Problem**: You need to detect repeated elements in a collection.
// - **Abstract Solution**: Efficiently keep track of previously seen elements while iterating through the collection, so you don’t need to compare each element with every other element.

// **Questions:**
// - Can you describe this problem without mentioning arrays, loops, or numbers?
// - What would the general structure of a solution look like?

// **Explanation:** The core of the problem is identifying duplicates in a collection, and the key to solving it efficiently is keeping track of what you’ve already seen. This abstraction helps you realize that the problem isn’t about arrays or numbers but about managing previously seen information in an optimal way.

// ---

// ### 5. **Explore different levels of abstraction for this problem.**

// Let’s consider various perspectives.

// - **Low-level abstraction**: Your current approach compares each pair of elements directly.
// - **Mid-level abstraction**: You could use a **hash set** to store elements as you encounter them. If an element is already in the set, you know it's a duplicate. This reduces the time complexity to **O(n)**.
// - **High-level abstraction**: This is a problem of efficiently managing state (what you’ve seen) to avoid redundant checks.

// **Questions:**
// - What are other ways to track whether something has been seen before, aside from using loops?
// - Can you think of higher-level problems (in software or real life) where efficiently managing past data is key to solving the problem?

// **Explanation:** At the mid-level abstraction, using data structures like a hash set becomes apparent. At the higher-level abstraction, you see this as a general problem of state management, where the goal is to avoid reprocessing or redundant work.

// ---

// ### 6. **Apply the abstracted understanding back to the original problem.**

// Now, using the insights from the abstractions, we can optimize your solution.

// - **Optimized Solution**:
//   Instead of comparing each element to every other one, use a hash set to keep track of seen elements. As you iterate through the array, check whether the current element is already in the set. If it is, return `true` (found a duplicate). If not, add it to the set and continue.

// Here’s the updated solution:

// ```go
// func containsDuplicate(nums []int) bool {
// 	seen := make(map[int]bool) // Using a hash set (map in Go)
// 	for _, num := range nums {
// 		if seen[num] {
// 			return true
// 		}
// 		seen[num] = true
// 	}
// 	return false
// }
// ```

// **Questions:**
// - How does this abstracted understanding (tracking state with a hash set) apply to your original problem?
// - What insights do you gain from this solution compared to the brute-force approach?

// **Explanation:** By applying the higher-level understanding, you’ve optimized your solution to **O(n)**. This removes the inefficiency of comparing every element with every other element. You’ve also gained insight into how tracking previously seen information can simplify problems in other contexts.

// ---

// ### Recap of the Abstraction Process:
// 1. **Describe the problem**: You want to check for duplicates in an array.
// 2. **Identify core elements**: Inputs (array), goal (find duplicates), constraint (efficiency).
// 3. **Find analogies**: Using checklists to avoid rechecking.
// 4. **Abstract representation**: Managing previously seen elements to avoid redundant checks.
// 5. **Explore levels of abstraction**: Nested loops, hash sets, state management.
// 6. **Apply abstraction**: Using a hash set reduces the problem's complexity and generates a more efficient solution.

// This process can be applied to other problems to improve problem-solving skills by focusing on core structures and abstracting out unnecessary details.
