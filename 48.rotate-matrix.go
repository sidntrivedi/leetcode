package main

import "fmt"

// You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).

// You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.

func rotate(matrix [][]int) {
	// Transpose the array.
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix[0]); j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = tmp
		}
	}

	// // Reverse the array.
	l := len(matrix[0])
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < l/2; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[i][l-j-1]
			matrix[i][l-j-1] = temp
		}
	}
	fmt.Println(matrix)
}
