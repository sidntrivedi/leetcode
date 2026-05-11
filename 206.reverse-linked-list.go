package main

/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
 *
 * https://leetcode.com/problems/reverse-linked-list/description/
 *
 * algorithms
 * Easy (80.51%)
 * Likes:    24574
 * Dislikes: 585
 * Total Accepted:    6.3M
 * Total Submissions: 7.8M
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * Given the head of a singly linked list, reverse the list, and return the
 * reversed list.
 *
 *
 * Example 1:
 *
 *
 * Input: head = [1,2,3,4,5]
 * Output: [5,4,3,2,1]
 *
 *
 * Example 2:
 *
 *
 * Input: head = [1,2]
 * Output: [2,1]
 *
 *
 * Example 3:
 *
 *
 * Input: head = []
 * Output: []
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the list is the range [0, 5000].
 * -5000 <= Node.val <= 5000
 *
 *
 *
 * Follow up: A linked list can be reversed either iteratively or recursively.
 * Could you implement both?
 *
 */

// @lc code=start

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// Recursive solution.
func reverseList(head *ListNode) *ListNode {
	// Base case.
	if head == nil || head.Next == nil {
		return head
	}

	newNode := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newNode
}

// Iterative solution.
// func reverseList(head *ListNode) *ListNode {
// 	var curr, prev *ListNode
// 	curr = head
// 	prev = nil
// 	for curr != nil {
// 		next := curr.Next
// 		curr.Next = prev
// 		prev = curr
// 		curr = next
// 	}

// 	return prev
// }

// @lc code=end
