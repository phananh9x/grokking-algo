package warmup

import (
	"strings"
)

/*
Solution: Pangram (easy)

A pangram is a sentence where every letter of the English alphabet appears at least once.

Given a string sentence containing English letters (lower or upper-case), return true if sentence is a pangram, or false otherwise.

Note: The given sentence might contain other characters like digits or spaces, your solution should handle these too.

Example 1:

Input: sentence = "TheQuickBrownFoxJumpsOverTheLazyDog"
Output: true
Explanation: The sentence contains at least one occurrence of every letter of the English alphabet either in lower or upper case.

Example 2:

Input: sentence = "This is not a pangram"
Output: false
Explanation: The sentence doesn't contain at least one occurrence of every letter of the English alphabet.
*/

// Solotion: Hashmap
// We can use a Hashmap to check if the given sentence is a pangram or not.
// The Hashmap will be used to store all the unique characters in the sentence. The algorithm works as follows:
// 1. Converts the sentence to lowercase.
// 2. Iterate over each character of the sentence using a loop.
// 3. Add each character to the Hashmap.
// 4. After looping through all characters,
//	compare the size of the Hashmap with 26 (total number of alphabets).
//  If the size of the Hashmap is equal to 26,
//  it means the sentence contains all the alphabets and is a pangram,
//  so the function will return true. Otherwise, it will return false.
// Time Complexity: O(N), where N is the number of characters in the sentence because it iterates over each character once.
// Space Complexity: O(1) because the Hashmap can store at most 26 characters.
func checkIfPagram(text string) bool {
	memo := map[string]bool{}

	for _, ch := range text {
		charToLower := strings.ToLower(string(ch))
		memo[charToLower] = true
	}

	return len(memo) == 26
}
