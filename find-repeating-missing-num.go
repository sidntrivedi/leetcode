/*
Source: https://takeuforward.org/data-structure/find-the-repeating-and-missing-numbers

Given an integer array nums of size n containing values from [1, n] and each value appears exactly once in the array,
except for A, which appears twice and B which is missing.
Return the values A and B, as an array of size 2, where A appears in the 0-th index and B in the 1st index.
Note: You are not allowed to modify the original array.
Example 1

Input: nums = [3, 5, 4, 1, 1]

Output: [1, 2]

Explanation:

1 appears two times in the array and 2 is missing from nums

# Example 2

Input: nums = [1, 2, 3, 6, 7, 5, 7]

Output: [7, 4]

Explanation:

7 appears two times in the array and 4 is missing from nums.
*/
package main

func findMissingRepeatingNumbers(nums []int) []int {
	xorVal := 0
	n := len(nums)

	// 1. XOR all nums and 1..n
	for i := 0; i < n; i++ {
		xorVal ^= nums[i]
		xorVal ^= i + 1
	}

	// 2. Find any set bit from the XOR result.
	// It can't be possible there's no set bit because
	// that would mean that the number is present as well as missing
	// from the array.
	setBit := xorVal & -xorVal

	// 3. Split into two groups.
	x, y := 0, 0
	for i := 0; i < n; i++ {
		if nums[i]&setBit != 0 {
			x ^= nums[i]
		} else {
			y ^= nums[i]
		}

		val := i + 1
		if val&setBit != 0 {
			x ^= val
		} else {
			y ^= val
		}
	}

	// 4. Check which one is repeated.
	for _, v := range nums {
		if v == x {
			return []int{x, y}
		}
	}
	return []int{y, x}
}
