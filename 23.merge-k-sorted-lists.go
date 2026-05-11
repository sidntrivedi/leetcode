package main

import "container/heap"

/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
 *
 * https://leetcode.com/problems/merge-k-sorted-lists/description/
 *
 * algorithms
 * Hard (59.47%)
 * Likes:    21329
 * Dislikes: 794
 * Total Accepted:    3M
 * Total Submissions: 5M
 * Testcase Example:  '[[1,4,5],[1,3,4],[2,6]]'
 *
 * You are given an array of k linked-lists lists, each linked-list is sorted
 * in ascending order.
 *
 * Merge all the linked-lists into one sorted linked-list and return it.
 *
 *
 * Example 1:
 *
 *
 * Input: lists = [[1,4,5],[1,3,4],[2,6]]
 * Output: [1,1,2,3,4,4,5,6]
 * Explanation: The linked-lists are:
 * [
 * ⁠ 1->4->5,
 * ⁠ 1->3->4,
 * ⁠ 2->6
 * ]
 * merging them into one sorted linked list:
 * 1->1->2->3->4->4->5->6
 *
 *
 * Example 2:
 *
 *
 * Input: lists = []
 * Output: []
 *
 *
 * Example 3:
 *
 *
 * Input: lists = [[]]
 * Output: []
 *
 *
 *
 * Constraints:
 *
 *
 * k == lists.length
 * 0 <= k <= 10^4
 * 0 <= lists[i].length <= 500
 * -10^4 <= lists[i][j] <= 10^4
 * lists[i] is sorted in ascending order.
 * The sum of lists[i].length will not exceed 10^4.
 *
 *
 */

// Definition for singly-linked list.
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest node.
	return pq[i].Val < pq[j].Val
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*ListNode)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // don't stop the GC from reclaiming the item eventually
	*pq = old[0 : n-1]
	return item
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Divide and Conquer approach.
func mergeKLists(lists []*ListNode) *ListNode {
	// Iterate over the lists.
	// Pick first pair and use the merge
	// two sorted list logic on that.
	// Similary for next pair.
	// Merge both pairs post this.
	pq := make(PriorityQueue, 0)

	// Iterate over the linked lists and keep
	// adding the elements to the priority queue.
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Push(&pq, lists[i])
		}
	}

	// Once the elements have been added to the heap, lets start to
	// pop them one by one and replace them with their next element in
	// the list.

	dummy := &ListNode{}
	tail := dummy
	for pq.Len() != 0 {
		// Pop an element.
		poppedElement := heap.Pop(&pq)
		poppedNode := poppedElement.(*ListNode)

		next := poppedNode.Next
		tail.Next = poppedNode
		tail = poppedNode

		// Push the next element.
		if next != nil {
			heap.Push(&pq, next)
		}
	}
	tail.Next = nil
	return dummy.Next
}

// @lc code=end
