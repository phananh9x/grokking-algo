package warmup

import "strings"

/*
Reverse Vowels (easy)
Problem Statement
Given a string s, reverse only all the vowels in the string and return it.

The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.

Example 1:

Input: s= "hello"
Output: "holle"
Example 2:

Input: s= "AEIOU"
Output: "UOIEA"
Example 3:

Input: s= "DesignGUrus"
Output: "DusUgnGires"
*/

//Solution: Two-pointer
//We can follow a two-pointer approach to reverse vowels in a string.
//The algorithm will start by initializing two pointers, "first" and "last" to the start and end of the string, respectively.
//The algorithm then uses these two pointers to traverse the string and skip over any non-vowel characters.
//When two vowels are found, the algorithm swaps them and updates the pointers.
//The process continues until the two pointers cross over each other.

func reverseVowels(s string) string {
	// list of vowels in English
	vowels := "aeiuoAEIUO"
	// initialize pointers for start and end of the string
	first := 0
	last := len(s) - 1
	arr := strings.Split(s, "")
	for first < last {
		// increment the start pointer until a vowel is found
		for first < last && !strings.Contains(vowels, string(s[first])) {
			first++
		}
		// decrement the end pointer until a vowel is found
		for last > first && !strings.Contains(vowels, string(s[last])) {
			last--
		}
		// swap the values of first and last if both are vowels
		arr[first], arr[last] = arr[last], arr[first]
		// move the pointers towards the center
		first++
		last--
	}
	return strings.Join(arr, "")
}
