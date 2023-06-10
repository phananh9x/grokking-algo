package warmup

/*
Solution: Valid Anagram
Problem Statement
Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

Example 1:

Input: s = "listen", t = "silent"
Output: true
Example 2:

Input: s = "rat", t = "car"
Output: false
Example 3:

Input: s = "hello", t = "world"
Output: false

Solution
We can solve this problem by calculating how many times each character appears in both the strings. If the frequency of each character is the same in both the given strings, we can conclude that the strings are anagrams of each other.

We can use a hash map to store the frequency of each character in both strings.

For each character in the string, the frequency of that character in string s is incremented and the frequency of that character in string t is decremented. After iterating over all characters in the strings, we will check if the frequency of all characters is 0. If it is, the strings are anagrams of each other and the function returns true. Otherwise, it returns false.

Time Complexity: O(n)
Space Complexity: O(1)

*/
func isValidAnagram(s1, s2 string) bool {
	flatMap := map[string]int{}

	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		value, existed := flatMap[string(s1[i])]
		if existed {
			flatMap[string(s1[i])] = value + 1
		} else {
			flatMap[string(s1[i])] = 1
		}

		valueS2, existedS2 := flatMap[string(s2[i])]
		if existedS2 {
			flatMap[string(s2[i])] = valueS2 - 1
		} else {
			flatMap[string(s2[i])] = -1
		}
	}
	for _, v := range flatMap {
		if v != 0 {
			return false
		}
	}

	return true

}
