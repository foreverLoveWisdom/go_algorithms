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

// - **Analogy**: Imagine you are checking a group of people entering a room and trying to remember whether you’ve already seen someone.
//  If you try to remember each person by looking at everyone else again, that takes time. Instead, you could use a checklist to mark people you’ve seen. Each new person can be quickly checked against this list.

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

// Great! Let's build abstraction layers based on your nested loop solution so you can apply this structure to similar problems in life. We'll form **three layers of abstraction**:

// 1. **Low-level (Concrete Implementation)**: What you are doing right now with nested loops.
// 2. **Mid-level (General Structure)**: A pattern that can apply to many problems.
// 3. **High-level (Abstract Concepts)**: Broad principles you can use in other real-life contexts.

// ---

// ### 1. **Low-Level (Concrete Implementation)**

// The nested loop solution works by **comparing each element with every other element**. This gives you a brute-force approach that guarantees you’ll find duplicates, but it can be inefficient as the input size grows. The trade-off here is simplicity over performance.

// #### **Use cases for the nested loop approach:**
// - **Direct Pairwise Comparison**: Anytime you need to compare all combinations of things, such as checking relationships between every two items.
//   - **Example**: Imagine you’re organizing a dinner party and want to avoid seating guests who dislike each other next to one another. You manually compare every pair of guests to see if there’s a conflict.

// #### **Steps to recognize this in real life:**
// - **Identify** a situation where you need to compare all elements with each other.
// - **Apply** the brute-force method (nested loop) as a simple, guaranteed way to get an answer, even though it might take time.

// ---

// ### 2. **Mid-Level (General Structure)**

// At a higher level, your nested loop approach solves the problem by **exhaustively testing every combination**. This pattern applies whenever you want to ensure that every possible pair or combination is checked.

// #### **Key idea**:
// - **Exhaustive Search**: Try every combination or option to guarantee no stone is left unturned.

// #### **Use cases**:
// - **Matching problems**: Anytime you need to test all pairs or combinations, like in:
//   - **Teamwork assignments**: Assigning every pair of people to see who works best together.
//   - **Event scheduling**: Checking every possible time slot for an event to see which times conflict for participants.

// #### **Steps to recognize this in real life**:
// - You have multiple items, and every item needs to be compared with every other item (or option) to get the full picture.
// - **Decision criteria**: Use this method when you **value completeness over efficiency** and need to ensure no option is left unchecked.

// ---

// ### 3. **High-Level (Abstract Concepts)**

// At the most abstract level, the nested loop solution is a **brute-force comparison**. This approach applies to **any problem where you don’t know any shortcuts or optimizations yet** but want to ensure every possibility is explored.

// #### **Key idea**:
// - **Systematic Exhaustion**: Test every possible option without making assumptions.

// #### **Use cases**:
// - **Decision making**: You’re trying to find the best option out of many but don’t know any shortcuts to narrow down choices.
//   - **Example**: Imagine trying to choose the best city to move to. Without knowing which factors matter most (e.g., weather, job market, cost of living), you systematically go through every factor for each city and compare them all.

// - **Trial and error**: When you don’t have a clear formula or pattern to follow, brute-force can be the first step before discovering patterns.
//   - **Example**: Testing every possible recipe ingredient to perfect a dish.

// #### **Steps to recognize this in real life**:
// - You face a situation where there’s **no clear shortcut** or optimization, and you need to evaluate each option methodically.
// - **Decision criteria**: Use this when you're dealing with an unknown system, and you need to explore all the options first before finding ways to improve.

// ---

// ### Putting It All Together: Solving Similar Problems in Life

// Here’s how you can use this abstraction process to apply the **nested loop** mechanism in other real-life problems:

// 1. **Start with the brute-force approach**:
//    - Just like your nested loop, when facing a new problem where you don’t know any optimizations or patterns, start by checking all combinations (i.e., compare each option to every other option).

// 2. **Look for patterns**:
//    - Once you’ve used the brute-force method and explored all possibilities, you’ll often start to see **patterns** (like repeated comparisons). This will help you identify where optimizations, like using a "checklist" (hash set analogy), could make things more efficient in future iterations.

// 3. **Move to higher-level thinking**:
//    - After solving the problem once using brute force, think about the **patterns you’ve learned**. Can you eliminate unnecessary comparisons next time? Can you create a more abstract framework to solve this type of problem faster in the future?

// ---

// ### Example for Future Use: Choosing a Supplier for a Product

// 1. **Low-level (Concrete)**: Compare every supplier on every criterion (price, quality, delivery time). You use nested loops to check each supplier against every criterion, ensuring no supplier is missed.

// 2. **Mid-level (General)**: You’re performing an exhaustive search of suppliers, guaranteeing you get the full picture even though it takes time. Once you’ve done this, you might realize certain suppliers consistently score well, meaning you can eliminate poor performers early in future decisions.

// 3. **High-level (Abstract)**: At the core, you’re solving a **decision-making** problem. You’re evaluating a large number of options (suppliers) without shortcuts, which means brute-force is the starting point. After this initial round, you can begin to **optimize** by focusing on key suppliers who consistently meet your needs, reducing future effort.

// ---

// ### Summary:

// - **Low-level (Concrete)**: Start with brute-force methods like nested loops when you need to compare everything to everything else.
// - **Mid-level (General)**: Recognize the pattern of exhaustive search when you want to guarantee no option is missed.
// - **High-level (Abstract)**: In life, use brute-force when you don’t have shortcuts, and it’s more important to ensure completeness than efficiency.

// By applying this structure, you’ll be able to use the nested loop approach as a base method for solving any problem where no shortcuts are known, then optimize from there as you find patterns.
