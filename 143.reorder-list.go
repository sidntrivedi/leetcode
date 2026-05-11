package main

/*
 * @lc app=leetcode id=143 lang=golang
 *
 * [143] Reorder List
 *
 * https://leetcode.com/problems/reorder-list/description/
 *
 * algorithms
 * Medium (65.18%)
 * Likes:    12590
 * Dislikes: 522
 * Total Accepted:    1.5M
 * Total Submissions: 2.3M
 * Testcase Example:  '[1,2,3,4]'
 *
 * You are given the head of a singly linked-list. The list can be represented
 * as:
 *
 *
 * L0 → L1 → … → Ln - 1 → Ln
 *
 *
 * Reorder the list to be on the following form:
 *
 *
 * L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
 *
 *
 * You may not modify the values in the list's nodes. Only nodes themselves may
 * be changed.
 *
 *
 * Example 1:
 *
 *
 * Input: head = [1,2,3,4]
 * Output: [1,4,2,3]
 *
 *
 * Example 2:
 *
 *
 * Input: head = [1,2,3,4,5]
 * Output: [1,5,2,4,3]
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the list is in the range [1, 5 * 10^4].
 * 1 <= Node.val <= 1000
 *
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	curr := head
	length := 0
	nodeStack := make([]*ListNode, 0)

	for curr != nil {
		// Pushing elements to the stack.
		nodeStack = append(nodeStack, curr)
		curr = curr.Next
		length++
	}

	curr = head
	for i := 0; i < length/2; i++ {
		originalNext := curr.Next
		popElement := pop(&nodeStack)
		curr.Next = popElement
		popElement.Next = originalNext
		curr = curr.Next.Next
	}
	curr.Next = nil
}

// Pop functionality for the stack
// to retrieve the linked list elements.
func pop(nodeStack *[]*ListNode) *ListNode {
	stack := *nodeStack
	element := stack[len(stack)-1]
	*nodeStack = stack[:len(stack)-1]
	return element
}

// @lc code=end
