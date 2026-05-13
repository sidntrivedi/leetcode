package main

/*
 * @lc app=leetcode id=572 lang=golang
 *
 * [572] Subtree of Another Tree
 *
 * https://leetcode.com/problems/subtree-of-another-tree/description/
 *
 * algorithms
 * Easy (51.54%)
 * Likes:    8935
 * Dislikes: 605
 * Total Accepted:    1.3M
 * Total Submissions: 2.5M
 * Testcase Example:  '[3,4,5,1,2]\n[4,1,2]'
 *
 * Given the roots of two binary trees root and subRoot, return true if there
 * is a subtree of root with the same structure and node values of subRoot and
 * false otherwise.
 *
 * A subtree of a binary tree tree is a tree that consists of a node in tree
 * and all of this node's descendants. The tree tree could also be considered
 * as a subtree of itself.
 *
 *
 * Example 1:
 *
 *
 * Input: root = [3,4,5,1,2], subRoot = [4,1,2]
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]
 * Output: false
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the root tree is in the range [1, 2000].
 * The number of nodes in the subRoot tree is in the range [1, 1000].
 * -10^4 <= root.val <= 10^4
 * -10^4 <= subRoot.val <= 10^4
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
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	// Iterate to the nodes and if the node is same as root of
	// subtree root compare the left and right of original and
	// subtree.
	if root == nil {
		return false
	}

	if isSameTreeCheck(root, subRoot) {
		return true
	}

	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

// isSameTree returns if the tree starting from the two nodes
// passed as params are same or not.
func isSameTreeCheck(root *TreeNode, root1 *TreeNode) bool {
	if root == nil && root1 != nil {
		return false
	}
	if root != nil && root1 == nil {
		return false
	}
	if root == nil && root1 == nil {
		return true
	}
	if root.Val != root1.Val {
		return false
	}
	leftTree := isSameTreeCheck(root.Left, root1.Left)
	rightTree := isSameTreeCheck(root.Right, root1.Right)
	if leftTree && rightTree {
		return true
	}
	return false
}

// @lc code=end
