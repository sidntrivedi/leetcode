package main

/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 *
 * https://leetcode.com/problems/spiral-matrix/description/
 *
 * algorithms
 * Medium (56.72%)
 * Likes:    17529
 * Dislikes: 1574
 * Total Accepted:    2.4M
 * Total Submissions: 4.2M
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * Given an m x n matrix, return all elements of the matrix in spiral order.
 *
 *
 * Example 1:
 *
 *
 * Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
 * Output: [1,2,3,6,9,8,7,4,5]
 *
 *
 * Example 2:
 *
 *
 * Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
 * Output: [1,2,3,4,8,12,11,10,9,5,6,7]
 *
 *
 *
 * Constraints:
 *
 *
 * m == matrix.length
 * n == matrix[i].length
 * 1 <= m, n <= 10
 * -100 <= matrix[i][j] <= 100
 *
 *
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	rows := len(matrix)
	columns := len(matrix[0])
	result := make([]int, 0)
	// Boundary markers as mentioned in the diagram below.
	top := 0
	bottom := rows - 1
	left := 0
	right := columns - 1

	/*
	   top, left            right
	   *------------------*
	   |                  |
	   |                  |
	   |                  |
	   |                  |
	   *------------------*
	   bottom

	   Populating the 2d array one layer at a time.
	   Workflow:
	   1. Populate values from left to right boundary index.
	   2. Populate from up+1 (since the last cell already is populated) in previous step to down.
	   3. Populate from right-1 (since the bottom right cell already populated in step 2) back to left boundary index.
	   4. Populate from down-1 (since the bottom left cell already populated in step 3) to the top left boundary index (up).
	   5. In step 3 and 4, there's a check to only populate if the row or column is different from what we already populated in
	       step 1 and 2.
	*/

	for {
		// Exit if the left and right boundaries have
		// exceeded each other. Same for up and down boundaries.
		// Means we have populated all the required cells in the array.
		if left > right || top > bottom {
			break
		}

		// Populate from left to right boundary.
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}

		// Populate from top right to bottom right.
		for i := top + 1; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}

		// Populate from bottom right to bottom left.
		if top != bottom {
			for i := right - 1; i >= left; i-- {
				result = append(result, matrix[bottom][i])
			}
		}

		// Populate from bottom left to top left.
		if left != right {
			for i := bottom - 1; i > top; i-- {
				result = append(result, matrix[i][left])
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return result
}

// @lc code=end
