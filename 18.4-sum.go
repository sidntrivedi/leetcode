package main

import "slices"

/*
 * @lc app=leetcode id=18 lang=golang
 *
 * [18] 4Sum
 *
 * https://leetcode.com/problems/4sum/description/
 *
 * algorithms
 * Medium (40.56%)
 * Likes:    12898
 * Dislikes: 1540
 * Total Accepted:    1.7M
 * Total Submissions: 4.2M
 * Testcase Example:  '[1,0,-1,0,-2,2]\n0'
 *
 * Given an array nums of n integers, return an array of all the unique
 * quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:
 *
 *
 * 0 <= a, b, c, d < n
 * a, b, c, and d are distinct.
 * nums[a] + nums[b] + nums[c] + nums[d] == target
 *
 *
 * You may return the answer in any order.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [1,0,-1,0,-2,2], target = 0
 * Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [2,2,2,2,2], target = 8
 * Output: [[2,2,2,2]]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 200
 * -10^9 <= nums[i] <= 10^9
 * -10^9 <= target <= 10^9
 *
 *
 */

// @lc code=start

// Approach:
// 1. Sort the array.
// 2. Need to find a, b,c,d.
// 3. Fix a and b. Find c and d by (target - (a+b))
// 4. If a and next element is same, skip recalculating the quadruplets.

func fourSum(nums []int, target int) [][]int {
	slices.Sort(nums)
	res := [][]int{}
	n := len(nums)

	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			if nums[j] == nums[j-1] {
				continue
			}
			a := nums[i]
			b := nums[j]

			for k := n - 1; k > j; k-- {
				for l := k - 1; l > j; l-- {
					// Best case, append to result.
					if (nums[k] + nums[l]) == target-(a+b) {
						res = append(res, []int{a, b, nums[l], nums[k]})
					}
				}
			}
		}
	}
	return res
}

// @lc code=end
