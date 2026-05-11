package main

/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 *
 * https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/
 *
 * algorithms
 * Medium (51.49%)
 * Likes:    21290
 * Dislikes: 910
 * Total Accepted:    4.2M
 * Total Submissions: 8.2M
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * Given the head of a linked list, remove the n^th node from the end of the
 * list and return its head.
 *
 *
 * Example 1:
 *
 *
 * Input: head = [1,2,3,4,5], n = 2
 * Output: [1,2,3,5]
 *
 *
 * Example 2:
 *
 *
 * Input: head = [1], n = 1
 * Output: []
 *
 *
 * Example 3:
 *
 *
 * Input: head = [1,2], n = 1
 * Output: [1]
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the list is sz.
 * 1 <= sz <= 30
 * 0 <= Node.val <= 100
 * 1 <= n <= sz
 *
 *
 *
 * Follow up: Could you do this in one pass?
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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var previousIndex, removeIndex int
	previousIndex = 0
	curr := head
	addressMap := make(map[int]*ListNode, 0)
	counter := 0

	// Store all the nodes in a map
	// for making changes based on index later.
	for curr != nil {
		addressMap[counter] = curr
		counter++
		curr = curr.Next
	}

	removeIndex = counter - n
	if removeIndex == 0 {
		return head.Next
	}

	if removeIndex <= 0 {
		previousIndex = -1
	} else {
		previousIndex = removeIndex - 1
	}

	nodeToRemove := addressMap[removeIndex]
	if previousIndex != -1 {
		previousNode := addressMap[previousIndex]
		if nodeToRemove.Next == nil {
			previousNode.Next = nil
		} else {
			previousNode.Next = nodeToRemove.Next
		}
	}
	return head
}

// @lc code=end
