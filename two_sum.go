package main

// Given an array of integers nums and an integer target,
// return indices of the two numbers such that they add up to target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// You can return the answer in any order.

// Example 1:

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

// Example 2:

// Input: nums = [3,2,4], target = 6
// Output: [1,2]

// Example 3:

// Input: nums = [3,3], target = 6
// Output: [0,1]

// Constraints:

//     2 <= nums.length <= 104
//     -109 <= nums[i] <= 109
//     -109 <= target <= 109
//     Only one valid answer exists.

// Follow-up: Can you come up with an algorithm that is less than O(n2) time complexity?

func twoSum(nums []int, target int) []int {
	// the end goal here is to search for the 2nd value
	// and currently I am using a loop
	// so, what if I can find the 2nd value without the need for the 2nd loop
	// since this is an array problem
	// my tool is loop, if condition here
	// and I have the target number
	// the 2nd value should be constructed somewhere already
	// and its lookup time should be better than O(n)
	// Otherwise, I will have to use a loop
	// But 2 flat loops still better than nested loops since it is O(n)
	// what is the data structure that has be the best lookup time:
	// map
	// so, I will create a map
	// so, since the map lookup time is O(1) and I need the key, so the 2nd value should be the key of this map
	// and the value should be the index
	// so, I will loop through the array, starting from the first element, I check if the 2nd number is in the map
	// if yes, then I will return the current index and the value of the 2nd number,
	// if not, I will move on to the next number
	hash := make(map[int]int)
	for i, firstNum := range nums {
		secondNum := target - firstNum
		if _, ok := hash[secondNum]; ok {
			return []int{hash[secondNum], i}
		}
		hash[firstNum] = i
	}
	return []int{}
}
