package main

/*
 * @lc app=leetcode id=100 lang=golang
 *
 * [100] Same Tree
 *
 * https://leetcode.com/problems/same-tree/description/
 *
 * algorithms
 * Easy (67.06%)
 * Likes:    12924
 * Dislikes: 301
 * Total Accepted:    3.4M
 * Total Submissions: 5.1M
 * Testcase Example:  '[1,2,3]\n[1,2,3]'
 *
 * Given the roots of two binary trees p and q, write a function to check if
 * they are the same or not.
 *
 * Two binary trees are considered the same if they are structurally identical,
 * and the nodes have the same value.
 *
 *
 * Example 1:
 *
 *
 * Input: p = [1,2,3], q = [1,2,3]
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: p = [1,2], q = [1,null,2]
 * Output: false
 *
 *
 * Example 3:
 *
 *
 * Input: p = [1,2,1], q = [1,1,2]
 * Output: false
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in both trees is in the range [0, 100].
 * -10^4 <= Node.val <= 10^4
 *
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
	// Traverse both the trees and keep on comparing.

	if p == nil && q == nil {
		return true
	}
	if p == nil && q != nil {
		return false
	}
	if p != nil && q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	leftSame := isSameTree(p.Left, q.Left)
	rightSame := isSameTree(p.Right, q.Right)
	return leftSame && rightSame
}

// @lc code=end
