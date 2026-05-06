package main

/*
 * @lc app=leetcode id=287 lang=golang
 *
 * [287] Find the Duplicate Number
 */

// @lc code=start

// Solution using the Floyd's Tortoise and Hare algorithm.
// Uses slow and fast pointer.
// Basically need to find start point of the cycle.
func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[0]

	for {
		slow = nums[slow]
		fast = nums[nums[fast]]

		if slow == fast {
			break
		}
	}

	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow

}

// Brute force solution.
// func findDuplicate(nums []int) int {
// 	dup := 0
// 	for i := 0; i < len(nums); i++ {
// 		pointer = nums[i]
// 		if nums[i] < 0 {
// 			pointer = nums[i] * -1
// 		}

// 		if nums[pointer] < 0 {
// 			dup = pointer
// 			break
// 		}
// 		nums[pointer] = -1 * nums[pointer]
// 	}
// 	return dup
// }

// @lc code=end
