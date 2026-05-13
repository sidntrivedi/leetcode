package main

import "container/heap"

/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 *
 * https://leetcode.com/problems/top-k-frequent-elements/description/
 *
 * algorithms
 * Medium (66.38%)
 * Likes:    19552
 * Dislikes: 857
 * Total Accepted:    3.6M
 * Total Submissions: 5.4M
 * Testcase Example:  '[1,1,1,2,2,3]\n2'
 *
 * Given an integer array nums and an integer k, return the k most frequent
 * elements. You may return the answer in any order.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [1,1,1,2,2,3], k = 2
 *
 * Output: [1,2]
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [1], k = 1
 *
 * Output: [1]
 *
 *
 * Example 3:
 *
 *
 * Input: nums = [1,2,1,2,1,2,3,1,3,2], k = 2
 *
 * Output: [1,2]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 10^5
 * -10^4 <= nums[i] <= 10^4
 * k is in the range [1, the number of unique elements in the array].
 * It is guaranteed that the answer is unique.
 *
 *
 *
 * Follow up: Your algorithm's time complexity must be better than O(n log n),
 * where n is the array's size.
 *
 */

// Item to be stored in the heap
type Pair struct {
	num  int
	freq int
}

type MinHeap []Pair

func (m MinHeap) Len() int {
	return len(m)
}

func (m MinHeap) Less(i, j int) bool {
	return m[i].freq < m[j].freq
}

func (m MinHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MinHeap) Push(x any) {
	*m = append(*m, x.(Pair))
}

func (m *MinHeap) Pop() any {
	old := *m
	n := len(old)
	item := old[n-1]
	*m = old[:n-1]
	return item
}

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	var minHeap MinHeap
	heap.Init(&minHeap)
	frequencyMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		frequencyMap[nums[i]] = frequencyMap[nums[i]] + 1
	}

	for key, v := range frequencyMap {
		heap.Push(&minHeap, Pair{
			num:  key,
			freq: v})
		if minHeap.Len() > k {
			for minHeap.Len() > k {
				heap.Pop(&minHeap)
			}
		}
	}

	// Return the remaining items from the minheap.
	ans := make([]int, 0, k)

	for minHeap.Len() > 0 {
		val := heap.Pop(&minHeap).(Pair)
		ans = append(ans, val.num)
	}
	return ans
}

// @lc code=end
