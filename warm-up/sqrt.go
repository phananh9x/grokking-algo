package warmup

import "math"

/**
Solotion: Sqrt (easy)

Given a non-negative integer x, return the square root of x rounded down to the nearest integer. The returned integer should be non-negative as well.

You must not use any built-in exponent function or operator.

For example, do not use pow(x, 0.5) in c++ or x ** 0.5 in python.

Example 1:

Input: x = 8
Output: 2
Explanation: The square root of 8 is 2.8284, and since we need to return the floor of the square root (integer), hence we returned 2.
Example 2:

Input: x = 4
Output: 2
Explanation: The square root of 4 is 2.
Example 3:

Input: x = 2
Output: 1
Explanation: The square root of 2 is 1.414, and since we need to return the floor of the square root (integer), hence we returned 1.
*/

// Solotion 1: Bruth Force
// Time Complexity: O(N)
// Space Complexity: O(N)
func sqrt(num int) int {
	if num < 2 {
		return num
	}
	myArray := []int{}
	for i := 1; i <= num/2; i++ {
		if i*i > num {
			break
		}
		if i*i == num {
			return i
		}
		myArray = append(myArray, i)
	}

	return myArray[len(myArray)-1]
}

// Solotion 2: Binary Search
// We can follow the Binary Search approach to solve this problem.
// The algorithm will start with two pointers "left" and "right",
//  pointing to the lower and upper bounds of the range in which the square root of the input number is expected to be.
//
// The algorithm then finds the pivot point in this range, calculates the square of pivot, and checks if the result is equal to,
// greater than, or less than the input number.
//
// If the result is greater than the input number,
// the algorithm narrows down the range by moving the right pointer to the pivot point minus 1.
//
// If the result is less than the input number, the algorithm expands the range by moving the left pointer to the pivot point plus 1.
//
// The algorithm repeats these steps until the result is equal to the input number,
// or the range has become so small that left pointer is greater than the right pointer. In this case, the algorithm returns the right pointer, which will be the floor of the square root of the input number.
//
// Time Complexity: O(logN) where N is the input number  because it uses binary search to find the square root.
// Space Complexity: O(1) because it only uses a few variables to store the pointers (left, right, pivot, and num), and the size of these variables does not grow with the input size.
func sqrtBinarySearch(num int) int {
	if num < 2 {
		return num
	}
	left := 2
	right := int(math.Floor(float64(num) / 2))
	var pivot, temp int
	for left <= int(right) {
		pivot = left + (right-left)/2
		temp = pivot * pivot
		if temp > num {
			right = pivot - 1
		} else if temp < num {
			left = pivot + 1
		} else {
			return pivot
		}
	}

	return right
}
