package main

/*
 * @lc app=leetcode id=50 lang=golang
 *
 * [50] Pow(x, n)
 *
 * https://leetcode.com/problems/powx-n/description/
 *
 * algorithms
 * Medium (38.61%)
 * Likes:    11727
 * Dislikes: 10557
 * Total Accepted:    2.7M
 * Total Submissions: 7M
 * Testcase Example:  '2.00000\n10'
 *
 * Implement pow(x, n), which calculates x raised to the power n (i.e.,
 * x^n).
 *
 *
 * Example 1:
 *
 *
 * Input: x = 2.00000, n = 10
 * Output: 1024.00000
 *
 *
 * Example 2:
 *
 *
 * Input: x = 2.10000, n = 3
 * Output: 9.26100
 *
 *
 * Example 3:
 *
 *
 * Input: x = 2.00000, n = -2
 * Output: 0.25000
 * Explanation: 2^-2 = 1/2^2 = 1/4 = 0.25
 *
 *
 *
 * Constraints:
 *
 *
 * -100.0 < x < 100.0
 * -2^31 <= n <= 2^31-1
 * n is an integer.
 * Either x is not zero or n > 0.
 * -10^4 <= x^n <= 10^4
 *
 *
 */

// @lc code=start

// Use fast exponentitation for calculating powers faster.
// FE is to compute x^n by repeatedly squaring instead of
// multiplying x n times.

// Eg: x^10 = (x^2)^5
// x^5 = x * (x^2)^2
func myPow(x float64, n int) float64 {

	// Base case.
	if n == 0 {
		return 1
	}

	if n < 0 {
		return 1 / myPow(x, -n)
	}

	half := myPow(x, n/2)

	if n%2 == 0 {
		return half * half
	} else {
		return half * half * x
	}
}

// // Brute force solution.
// func myPow(x float64, n int) float64 {
// 	var p float64
// 	p = 1.0

// 	if n < 0 {
// 		for i := n; i < 0; i++ {
// 			p = p * (1 / x)
// 		}
// 	} else {
// 		for i := 1; i <= n; i++ {
// 			p = p * x
// 		}
// 	}
// 	return p
// }

// @lc code=end
