/*
 * @lc app=leetcode id=31 lang=golang
 *
 * [31] Next Permutation
 */

package main

// @lc code=start

// Start from the right.
// Find the first index where nums[i] < nums[i+1]. That is the pivot point.
// Once that is found, find the smallest element to the right that is
// larger than the pivot element.
// Swap the pivot with that element and then reverse the tail from the pivot.
func nextPermutation(nums []int) {
	pivot := -1

	for i := len(nums) - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			pivot = i - 1
			break
		}
	}

	if pivot == -1 {
		reverse(nums, 0)
		return
	}

	swap := -1
	for j := len(nums) - 1; j > pivot; j-- {
		if nums[j] > nums[pivot] {
			swap = j
			break
		}
	}

	// Do swapping.
	nums[pivot], nums[swap] = nums[swap], nums[pivot]

	// Reverse the tail.
	reverse(nums, pivot+1)
}

func reverse(nums []int, left int) {
	right := len(nums) - 1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}

}

// @lc code=end
