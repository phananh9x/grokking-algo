package warmup

import (
	"regexp"
	"strings"
)

/*
Problem Statement
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.

Example 1:

Input: sentence = "A man, a plan, a canal, Panama!"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
Example 2:

Input: sentence = "Was it a car or a cat I saw?"
Output: true
Explanation: Explanation: "wasitacaroracatisaw" is a palindrome.


Solution: two pointer
We can use two pointers to solve this problem.

One pointer will start from the beginning of the string and the other from the end of the string.
The pointers move towards each other, skipping non-alphanumeric characters,
until they meet in the middle or one of the pointers points to an non-alphanumeric character that is different from its counterpart at the other end.
If all characters are the same and both pointers meet in the middle, the string is a palindrome.

Time Complexity: O(n), where n is the number of characters in the sentence, because it iterates over each character once.
Space Complexity: O(1), because our algorithm uses constant space to store a few variables, and the size of these variables does not grow with the input size.
*/
func isValidPalindrome(s string) bool {
	first := 0
	last := len(s) - 1
	for first < last {
		for first < last {
			matched, _ := regexp.MatchString("^[a-zA-Z0-9]+$", string(s[first]))
			if matched {
				break
			}
			first++
		}

		for first < last {
			matched, _ := regexp.MatchString("^[a-zA-Z0-9]+$", string(s[last]))
			if matched {
				break
			}
			last--
		}
		// compare the characters at first and last after converting them to lowercase
		if strings.ToLower(string(s[first])) != strings.ToLower(string(s[last])) {
			// if they are not equal, the string is not a palindrome
			return false
		}
		// move forward
		first++
		// move backward
		last--
	}

	return true
}
