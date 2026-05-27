package main

/*
 * @lc app=leetcode id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 *
 * https://leetcode.com/problems/search-a-2d-matrix/description/
 *
 * algorithms
 * Medium (53.95%)
 * Likes:    17992
 * Dislikes: 486
 * Total Accepted:    2.9M
 * Total Submissions: 5.4M
 * Testcase Example:  '[[1,3,5,7],[10,11,16,20],[23,30,34,60]]\n3'
 *
 * You are given an m x n integer matrix matrix with the following two
 * properties:
 *
 *
 * Each row is sorted in non-decreasing order.
 * The first integer of each row is greater than the last integer of the
 * previous row.
 *
 *
 * Given an integer target, return true if target is in matrix or false
 * otherwise.
 *
 * You must write a solution in O(log(m * n)) time complexity.
 *
 *
 * Example 1:
 *
 *
 * Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
 * Output: false
 *
 *
 *
 * Constraints:
 *
 *
 * m == matrix.length
 * n == matrix[i].length
 * 1 <= m, n <= 100
 * -10^4 <= matrix[i][j], target <= 10^4
 *
 *
 */

// @lc code=start

/*
Better approach: direct binary search on 2d matrix.

Lets say if the matrix is :
[
	[1,3,5,7]
	[10,11,16,20]
	[23,30,34,60]
]

in that case instead of converting whole 2d matrix into array,
we can map each element of the resulting 1d array as to what its row
and column would be in the 2d array.
*/

// Seems like binary search of a matrix.
func searchMatrix(matrix [][]int, target int) bool {
	// Put all the elements into a single 1d array and then do binary search.
	arr := make([]int, 0)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			arr = append(arr, matrix[i][j])
		}
	}

	// Binary search.
	begin := 0
	end := len(arr) - 1

	for begin <= end {
		mid := (begin + end) / 2
		if arr[mid] == target {
			return true
		}

		if arr[mid] < target {
			begin = mid + 1
		} else {
			end = mid - 1
		}
	}
	return false
}

// // @lc code=end
