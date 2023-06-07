package warmup

// Contains Duplicate (easy)

// Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

// Example 1:

// Input: nums= [1, 2, 3, 4]
// Output: false
// Explanation: There are no duplicates in the given array.
// Example 2:

// Input: nums= [1, 2, 3, 1]
// Output: true
// Explanation: '1' is repeating.

// Solution 1: bruth force
// We can use a brute force approach and compare each element with all other elements in the array.
// If any two elements are the same, we'll return true. If we've gone through the entire array and haven't found any duplicates, we'll return false.
// Time Complexity: The algorithm takes O(N^2) where  is the number of elements in the input array.
// 					This is because we need to compare each element with all other elements, which takes O(N^2) time.
// Space Complexity: The algorithm takes O(N1), because we only need a few variables to store the indices, which takes constant space.
func containsDupplicateBruthForce(params []int) bool {
	for i := 0; i < len(params); i++ {
		for j := i + 1; j < len(params); j++ {
			if params[i] == params[j] {
				return true
			}
		}
	}
	return false
}

// Solution 2: hash map
// containsDupplicate
// Time Complexity: The algorithm takes O(N) where N is the number of elements in the input array. This is because we iterate the array only once.
// Space Complexity: The algorithm takes O(1), as it stores the numbers in a set.
func containsDupplicate(params []int) bool {
	memo := map[int]bool{}

	for i := 0; i < len(params); i++ {
		if memo[params[i]] {
			return true
		}
		memo[params[i]] = true
	}

	return false
}
