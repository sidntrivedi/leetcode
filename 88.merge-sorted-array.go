package main

// num1 and num2 in increasing order.
// m and n : no. of elements in num1 and num2
// Merge num1 and num2 into single array of increasing order.
func merge(nums1 []int, m int, nums2 []int, n int) {

	// Three pointer approach.
	// Take three pointers: p1, p2 and p.
	// p1 : pointing in the first array from last
	// p2: pointing in the second array from last
	// p: point at m+n-1 in nums1 array
	//
	// if p2 > p1:
	// - set value of p2 element at nums1[p]
	// - decrement p2 and p, keep p1 same
	// if p2 <= p1:
	// - set value of p1 at p
	// - decrement p and p2

	p1, p2 := m-1, n-1

	for p := m + n - 1; p >= 0; p-- {
		if p2 < 0 {
			break
		}
		if p1 >= 0 && nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
	}
}
