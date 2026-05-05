package main

import "fmt"

func sortColors(nums []int) {
	/*
	   Dutch National Flag problem solution.
	*/
	low, mid := 0, 0
	high := len(nums) - 1
	for mid <= high {

		if nums[mid] == 0 {
			nums[low], nums[mid] = nums[mid], nums[low]
			mid++
			low++
		} else if nums[mid] == 2 {
			nums[high], nums[mid] = nums[mid], nums[high]
			high--
		} else {
			mid++
		}
	}
	fmt.Println(nums)
}
