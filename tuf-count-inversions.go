package main

/*
Source: https://takeuforward.org/plus/dsa/problems/count-inversions

Given an integer array nums. Return the number of inversions in the array.
Two elements a[i] and a[j] form an inversion if a[i] > a[j] and i < j.
It indicates how close an array is to being sorted.
A sorted array has an inversion count of 0.
An array sorted in descending order has maximum inversion.

Example 1
Input: nums = [2, 3, 7, 1, 3, 5]
Output: 5

Explanation:

The responsible indexes are:
nums[0], nums[3], values: 2 > 1 & indexes: 0 < 3
nums[1], nums[3], values: 3 > 1 & indexes: 1 < 3
nums[2], nums[3], values: 7 > 1 & indexes: 2 < 3
nums[2], nums[4], values: 7 > 3 & indexes: 2 < 4
nums[2], nums[5], values: 7 > 5 & indexes: 2 < 5

Example 2
Input: nums = [-10, -5, 6, 11, 15, 17]
Output: 0

Explanation:
nums is sorted, hence no inversions present.
*/

func numberOfInversions(nums []int) int64 {
	_, count := sortAndCount(nums)
	return count
}

func sortAndCount(nums []int) ([]int, int64) {
	if len(nums) <= 1 {
		return nums, 0
	}

	mid := len(nums) / 2
	left, leftCount := sortAndCount(nums[:mid])
	right, rightCount := sortAndCount(nums[mid:])

	merged, mergeCount := mergeAndCount(left, right)
	return merged, leftCount + rightCount + mergeCount
}

func mergeAndCount(left, right []int) ([]int, int64) {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	var count int64

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			count += int64(len(left) - i)
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result, count
}
