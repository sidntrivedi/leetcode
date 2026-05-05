package main

import "slices"

// Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals,
// and return an array of the non-overlapping intervals that cover all the intervals in the input.

// Example 1:
// Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
// Output: [[1,6],[8,10],[15,18]]
// Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].
// 1,3
// 2,6
// 8,10
// 15,18
// 1,2,3,6,8,10,15,18
//
// Example 2:
// Input: intervals = [[1,4],[4,5]]
// Output: [[1,5]]
// Explanation: Intervals [1,4] and [4,5] are considered overlapping.
//
// Example 3:
// Input: intervals = [[4,7],[1,4]]
// Output: [[1,7]]
// Explanation: Intervals [1,4] and [4,7] are considered overlapping.

func merge(intervals [][]int) [][]int {
	n := len(intervals)
	// 1. Sort the 2d array using the start index.
	// 2. Iterate from the first element and check if the end of first element is
	// less than the start of the next, if yes then combine them together.
	slices.SortFunc(intervals, func(a, b []int) int {
		if a[0] < b[0] {
			return -1
		}
		if a[0] > b[0] {
			return 1
		}
		return 0
	})

	res := make([][]int, 0)
	res = append(res, intervals[0])

	for i := 1; i < n; i++ {
		last := res[len(res)-1]
		if intervals[i][0] <= last[1] {
			end := max(last[1], intervals[i][1])
			res[len(res)-1][1] = end
		} else {
			res = append(res, intervals[i])
		}
	}

	return res
}
