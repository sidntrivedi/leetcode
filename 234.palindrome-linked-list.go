package main

/*
 * @lc app=leetcode id=234 lang=golang
 *
 * [234] Palindrome Linked List
 *
 * https://leetcode.com/problems/palindrome-linked-list/description/
 *
 * algorithms
 * Easy (57.88%)
 * Likes:    18381
 * Dislikes: 983
 * Total Accepted:    2.9M
 * Total Submissions: 5M
 * Testcase Example:  '[1,2,2,1]'
 *
 * Given the head of a singly linked list, return true if it is a palindrome or
 * false otherwise.
 *
 *
 * Example 1:
 *
 *
 * Input: head = [1,2,2,1]
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: head = [1,2]
 * Output: false
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the list is in the range [1, 10^5].
 * 0 <= Node.val <= 9
 *
 *
 *
 * Follow up: Could you do it in O(n) time and O(1) space?
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func isPalindrome(head *ListNode) bool {
	// Iterate till the middle of the linked list
	// and then reverse the second half of the LL.
	// Compare if the reversed second half is same
	// as the first half.

	// Find the mid using the slow and fast pointers.
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Now that we have found the mid node, we
	// need to start a loop from start till mid and
	// from mid to end and compare both halves.

	// Reverse the second half.
	var prev *ListNode
	for slow != nil {
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}

	dummy1 := head
	for prev != nil {
		if dummy1.Val != prev.Val {
			return false
		}
		dummy1 = dummy1.Next
		prev = prev.Next
	}
	return true

}

// Solution using Stack to store the elements
// in reverse and then popping them one by one
// and comparing with the original linked list.
// func isPalindrome(head *ListNode) bool {

// 	// Iterate all the nodes of the linked list
// 	// and push the elements into a stack.
// 	dummy := head
// 	st := make([]int, 0)

// 	for dummy != nil {
// 		st = append(st, dummy.Val)
// 		dummy = dummy.Next
// 	}

// 	isPalindrome := true
// 	dummy1 := head
// 	for dummy1 != nil {
// 		llVal := dummy1.Val
// 		popVal := st[len(st)-1]
// 		if popVal != llVal {
// 			isPalindrome = false
// 			break
// 		}
// 		st = st[:len(st)-1]
// 		dummy1 = dummy1.Next
// 	}

// 	return isPalindrome
// }

// @lc code=end
