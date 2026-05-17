package main

/*
 * @lc app=leetcode id=128 lang=golang
 *
 * [128] Longest Consecutive Sequence
 *
 * https://leetcode.com/problems/longest-consecutive-sequence/description/
 *
 * algorithms
 * Medium (47.09%)
 * Likes:    22981
 * Dislikes: 1243
 * Total Accepted:    3.3M
 * Total Submissions: 7.1M
 * Testcase Example:  '[100,4,200,1,3,2]'
 *
 * Given an unsorted array of integers nums, return the length of the longest
 * consecutive elements sequence.
 *
 * You must write an algorithm that runs in O(n) time.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [100,4,200,1,3,2]
 * Output: 4
 * Explanation: The longest consecutive elements sequence is [1, 2, 3, 4].
 * Therefore its length is 4.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [0,3,7,2,5,8,4,6,0,1]
 * Output: 9
 *
 *
 * Example 3:
 *
 *
 * Input: nums = [1,0,1,2]
 * Output: 3
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= nums.length <= 10^5
 * -10^9 <= nums[i] <= 10^9
 *
 *
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	numMap := make(map[int]int, len(nums))
	for _, n := range nums {
		numMap[n] = 1
	}
	longestStreak := 0

	for n, _ := range numMap {
		if _, exists := numMap[n-1]; !exists {
			currentNum := n
			currentStreak := 1

			for {
				if _, exists := numMap[currentNum+1]; !exists {
					break
				} else {
					currentNum++
					currentStreak++
				}
			}
			if currentStreak > longestStreak {
				longestStreak = currentStreak
			}
		}
	}
	return longestStreak
}

// @lc code=end
