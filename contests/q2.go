package contests

/* Problem

nums , integer k

in one op, increase or decrease any element of nums by 1

Modulo alternating if two distinct ints x and y ( 0 <= x , y <k )

For every even index i , nums[i] %k == x
For every odd index, nums[i] %k == y

Min number of ops to make nums modulo alternating.
*/

func minOperations(nums []int, k int) int {
	evenMod := make([]int, 0)
	oddMod := make([]int, 0)
	allEqual := true

	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			evenMod = append(evenMod, nums[i]%k)
		} else {
			oddMod = append(oddMod, nums[i]%k)
		}
		if i > 0 && nums[i] != nums[i-1] {
			allEqual = false
		}
	}
	if len(nums) == 1 {
		return 0
	}
	// If all the elements are same in array, return 1.
	if allEqual && len(nums) > 1 {
		return 1
	}

	leastEvenOps := 100000
	for i := 0; i < len(evenMod); i++ {
		steps := 0
		for j := 0; j < len(evenMod); j++ {
			if j == i {
				continue
			}

			if evenMod[j] >= evenMod[i] {
				steps = steps + evenMod[j] - evenMod[i]
			} else {
				steps = steps + evenMod[i] - evenMod[j]
			}
		}
		if steps < leastEvenOps {
			leastEvenOps = steps
		}
	}

	leastOddOps := 100000000
	for i := 0; i < len(oddMod); i++ {
		steps := 0
		for j := 0; j < len(oddMod); j++ {
			if j == i {
				continue
			}

			if oddMod[j] >= oddMod[i] {
				steps = steps + oddMod[j] - oddMod[i]
			} else {
				steps = steps + oddMod[i] - oddMod[j]
			}
		}
		if steps < leastOddOps {
			leastOddOps = steps
		}
	}

	return leastEvenOps + leastOddOps
}
