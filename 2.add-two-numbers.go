package main

/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 *
 * https://leetcode.com/problems/add-two-numbers/description/
 *
 * algorithms
 * Medium (48.40%)
 * Likes:    36818
 * Dislikes: 7263
 * Total Accepted:    7M
 * Total Submissions: 14.5M
 * Testcase Example:  '[2,4,3]\n[5,6,4]'
 *
 * You are given two non-empty linked lists representing two non-negative
 * integers. The digits are stored in reverse order, and each of their nodes
 * contains a single digit. Add the two numbers and return the sum as a linked
 * list.
 *
 * You may assume the two numbers do not contain any leading zero, except the
 * number 0 itself.
 *
 *
 * Example 1:
 *
 *
 * Input: l1 = [2,4,3], l2 = [5,6,4]
 * Output: [7,0,8]
 * Explanation: 342 + 465 = 807.
 *
 *
 * Example 2:
 *
 *
 * Input: l1 = [0], l2 = [0]
 * Output: [0]
 *
 *
 * Example 3:
 *
 *
 * Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
 * Output: [8,9,9,9,0,0,0,1]
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in each linked list is in the range [1, 100].
 * 0 <= Node.val <= 9
 * It is guaranteed that the list represents a number that does not have
 * leading zeros.
 *
 *
 */

// @lc code=start

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// Keep on adding individual node values.
	// Keep persisting the carry value in a variable to be used
	// in further loops.
	dummy := &ListNode{}
	tail := dummy
	carry := 0
	actualSum := 0
	l1Val := 0
	l2Val := 0

	for l1 != nil || l2 != nil {
		if l1 == nil {
			l1Val = 0
		} else {
			l1Val = l1.Val
		}
		if l2 == nil {
			l2Val = 0
		} else {
			l2Val = l2.Val
		}

		// Find the sum first, store that value in a node,
		// mark the next of the node to the new node.
		actualSum = (l1Val + l2Val + carry) % 10
		carry = (l1Val + l2Val + carry) / 10
		// fmt.Println("Sum, carry", actualSum, carry)
		newNode := &ListNode{
			Val: actualSum,
		}
		tail.Next = newNode
		tail = tail.Next

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry != 0 {
		newNode := &ListNode{
			Val: carry,
		}
		tail.Next = newNode
		tail = tail.Next
	}

	tail.Next = nil

	return dummy.Next

}

// @lc code=end
