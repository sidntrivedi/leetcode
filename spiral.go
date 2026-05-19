package main

import "fmt"

type SpiralPrinter struct {
	arr     [][]int
	rows    int
	columns int
}

func (s *SpiralPrinter) generateSpiral() {
	// Allocate a m x n 2d array.
	s.arr = make([][]int, s.rows)
	for i := 0; i < s.rows; i++ {
		s.arr[i] = make([]int, s.columns)
	}

	// Boundary markers as mentioned in the diagram below.
	top := 0
	bottom := s.rows - 1
	left := 0
	right := s.columns - 1
	value := 0 // first written value will be 1.

	/*
	   top, left            right
	   *------------------*
	   |                  |
	   |                  |
	   |                  |
	   |                  |
	   *------------------*
	   bottom

	   Populating the 2d array one layer at a time.
	   Workflow:
	   1. Populate values from left to right boundary index.
	   2. Populate from up+1 (since the last cell already is populated) in previous step to down.
	   3. Populate from right-1 (since the bottom right cell already populated in step 2) back to left boundary index.
	   4. Populate from down-1 (since the bottom left cell already populated in step 3) to the top left boundary index (up).
	   5. In step 3 and 4, there's a check to only populate if the row or column is different from what we already populated in
	       step 1 and 2.
	*/

	for {
		// Exit if the left and right boundaries have
		// exceeded each other. Same for up and down boundaries.
		// Means we have populated all the required cells in the array.
		if left > right || top > bottom {
			break
		}

		// Populate from left to right boundary.
		for i := left; i <= right; i++ {
			value++
			s.arr[top][i] = value
		}

		// Populate from top right to bottom right.
		for i := top + 1; i <= bottom; i++ {
			value++
			s.arr[i][right] = value
		}

		// Populate from bottom right to bottom left.
		if top != bottom {
			for i := right - 1; i >= left; i-- {
				value++
				s.arr[bottom][i] = value
			}
		}

		// Populate from bottom left to top left.
		if left != right {
			for i := bottom - 1; i > top; i-- {
				value++
				s.arr[i][left] = value
			}
		}
		left++
		right--
		top++
		bottom--
	}
}

func main() {
	m, n := 4, 5
	s := SpiralPrinter{
		rows:    m,
		columns: n,
	}
	s.generateSpiral()

	// Print the generated spiral matrix.
	for i := 0; i < len(s.arr); i++ {
		for j := 0; j < len(s.arr[0]); j++ {
			fmt.Printf("%d\t", s.arr[i][j])
		}
		fmt.Println()
	}
}
