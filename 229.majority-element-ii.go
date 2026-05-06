package main

/*
 * @lc app=leetcode id=229 lang=golang
 *
 * [229] Majority Element II
 *
 * https://leetcode.com/problems/majority-element-ii/description/
 *
 * algorithms
 * Medium (56.10%)
 * Likes:    11019
 * Dislikes: 503
 * Total Accepted:    1.3M
 * Total Submissions: 2.2M
 * Testcase Example:  '[3,2,3]'
 *
 * Given an integer array of size n, find all elements that appear more than ⌊
 * n/3 ⌋ times.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [3,2,3]
 * Output: [3]
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [1]
 * Output: [1]
 *
 *
 * Example 3:
 *
 *
 * Input: nums = [1,2]
 * Output: [1,2]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 5 * 10^4
 * -10^9 <= nums[i] <= 10^9
 *
 *
 *
 * Follow up: Could you solve the problem in linear time and in O(1) space?
 *
 */

// @lc code=start
func majorityElement(nums []int) []int {
	candidate1 := 0
	candidate2 := 0
	count1 := 0
	count2 := 0

	for _, n := range nums {
		if n == candidate1 {
			count1++
		} else if n == candidate2 {
			count2++
		} else if count1 == 0 {
			candidate1 = n
			count1 = 1
		} else if count2 == 0 {
			candidate2 = n
			count2 = 1
		} else {
			count1--
			count2--
		}
	}

	count1 = 0
	count2 = 0

	for _, n := range nums {
		if n == candidate1 {
			count1++
		} else if n == candidate2 {
			count2++
		}
	}

	result := []int{}

	if count1 > len(nums)/3 {
		result = append(result, candidate1)
	}

	if count2 > len(nums)/3 {
		result = append(result, candidate2)
	}

	return result

}

// @lc code=end
