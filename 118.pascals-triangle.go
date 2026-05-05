/*
 * @lc app=leetcode id=118 lang=golang
 *
 * [118] Pascal's Triangle
 */
package main

// @lc code=start
func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}

	start := [][]int{
		{1},
	}

	for i := 1; i < numRows; i++ {
		upperRow := start[len(start)-1]
		newRow := make([]int, len(upperRow)+1)

		newRow[0] = 1
		newRow[len(newRow)-1] = 1

		for j := 1; j < len(newRow)-1; j++ {
			newRow[j] = upperRow[j-1] + upperRow[j]
		}

		start = append(start, newRow)
	}

	return start
}

// @lc code=end
