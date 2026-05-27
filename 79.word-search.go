package main

/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 *
 * https://leetcode.com/problems/word-search/description/
 *
 * algorithms
 * Medium (47.33%)
 * Likes:    17728
 * Dislikes: 759
 * Total Accepted:    2.5M
 * Total Submissions: 5.3M
 * Testcase Example:  '[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"ABCCED"'
 *
 * Given an m x n grid of characters board and a string word, return true if
 * word exists in the grid.
 *
 * The word can be constructed from letters of sequentially adjacent cells,
 * where adjacent cells are horizontally or vertically neighboring. The same
 * letter cell may not be used more than once.
 *
 *
 * Example 1:
 *
 *
 * Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word
 * = "ABCCED"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word
 * = "SEE"
 * Output: true
 *
 *
 * Example 3:
 *
 *
 * Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word
 * = "ABCB"
 * Output: false
 *
 *
 *
 * Constraints:
 *
 *
 * m == board.length
 * n = board[i].length
 * 1 <= m, n <= 6
 * 1 <= word.length <= 15
 * board and word consists of only lowercase and uppercase English letters.
 *
 *
 *
 * Follow up: Could you use search pruning to make your solution faster with a
 * larger board?
 *
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	// Start iterating from the first element of the matrix.
	// As soon as the first letter is found, start dfs from that cell.
	// If the letter matches, mark that cell as visited.
	// Recursively search in all directions of the letter.
	// If found, return true.
	// If not, backtrack to the last visited word that was part of the
	// original word.

	rows := len(board)
	cols := len(board[0])

	var backtrack func(row, col, index int) bool
	backtrack = func(row, col, index int) bool {
		if index == len(word) {
			return true
		}

		if row < 0 || row >= rows || col < 0 || col >= cols || board[row][col] != word[index] {
			return false
		}

		// If we are here, it means we have got a letter matching.
		// Need to find others.

		// Add a visited mark at the current array spot.
		// Why? So that while checking in all directions later
		// we dont go back to the previous cell if the letters are
		// repeated.
		temp := board[row][col]
		board[row][col] = '#'

		// Seach in all 4 directions of the current cell.
		rowOffset := []int{0, 1, 0, -1}
		colOffset := []int{1, 0, -1, 0}

		resp := false
		for i := 0; i < 4; i++ {
			resp = backtrack(row+rowOffset[i], col+colOffset[i], index+1)
			if resp {
				break
			}
		}

		board[row][col] = temp
		return resp
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {

			// Match the first letter.
			if board[i][j] == word[0] && backtrack(i, j, 0) {
				return true
			}
		}
	}

	return false
}

// @lc code=end
