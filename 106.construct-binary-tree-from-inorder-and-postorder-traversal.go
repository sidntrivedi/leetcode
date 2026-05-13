package main

/*
 * @lc app=leetcode id=106 lang=golang
 *
 * [106] Construct Binary Tree from Inorder and Postorder Traversal
 *
 * https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
 *
 * algorithms
 * Medium (68.56%)
 * Likes:    8715
 * Dislikes: 159
 * Total Accepted:    957.5K
 * Total Submissions: 1.4M
 * Testcase Example:  '[9,3,15,20,7]\n[9,15,7,20,3]'
 *
 * Given two integer arrays inorder and postorder where inorder is the inorder
 * traversal of a binary tree and postorder is the postorder traversal of the
 * same tree, construct and return the binary tree.
 *
 *
 * Example 1:
 *
 *
 * Input: inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
 * Output: [3,9,20,null,null,15,7]
 *
 *
 * Example 2:
 *
 *
 * Input: inorder = [-1], postorder = [-1]
 * Output: [-1]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= inorder.length <= 3000
 * postorder.length == inorder.length
 * -3000 <= inorder[i], postorder[i] <= 3000
 * inorder and postorder consist of unique values.
 * Each value of postorder also appears in inorder.
 * inorder is guaranteed to be the inorder traversal of the tree.
 * postorder is guaranteed to be the postorder traversal of the tree.
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 || len(postorder) != len(inorder) {
		return nil
	}

	// Figure out the root from post order.
	root := postorder[len(postorder)-1]
	rootIdx := 0
	for i, e := range inorder {
		if e == root {
			rootIdx = i
		}
	}

	leftInOrder := inorder[:rootIdx]
	rightInOrder := inorder[(rootIdx + 1):]
	leftSize := rootIdx
	// rightSize := len(inorder) - rootIdx - 1

	leftPostOrder := postorder[:leftSize]
	rightPostOrder := postorder[leftSize : len(postorder)-1]
	rootNode := &TreeNode{
		Val: root,
	}

	rootNode.Left = buildTree(leftInOrder, leftPostOrder)
	rootNode.Right = buildTree(rightInOrder, rightPostOrder)

	return rootNode
}

// @lc code=end
