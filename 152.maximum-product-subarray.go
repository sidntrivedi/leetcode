package main

/*
 * @lc app=leetcode id=152 lang=golang
 *
 * [152] Maximum Product Subarray
 *
 * https://leetcode.com/problems/maximum-product-subarray/description/
 *
 * algorithms
 * Medium (36.34%)
 * Likes:    20502
 * Dislikes: 831
 * Total Accepted:    2M
 * Total Submissions: 5.6M
 * Testcase Example:  '[2,3,-2,4]'
 *
 * Given an integer array nums, find a subarray that has the largest product,
 * and return the product.
 *
 * The test cases are generated so that the answer will fit in a 32-bit
 * integer.
 *
 * Note that the product of an array with a single element is the value of that
 * element.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [2,3,-2,4]
 * Output: 6
 * Explanation: [2,3] has the largest product 6.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [-2,0,-1]
 * Output: 0
 * Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 2 * 10^4
 * -10 <= nums[i] <= 10
 * The product of any subarray of nums is guaranteed to fit in a 32-bit
 * integer.
 *
 *
 */

// @lc code=start
func maxProduct(nums []int) int {
	n := len(nums)
	maxSoFar := nums[0]
	minSoFar := nums[0]
	result := maxSoFar
	for i := 1; i < n; i++ {
		curr := nums[i]
		temp_max := max(curr, minSoFar*curr, maxSoFar*curr)
		minSoFar = min(curr, minSoFar*curr, maxSoFar*curr)
		maxSoFar = temp_max
		result = max(maxSoFar, result)

	}
	return result
}

// @lc code=end
