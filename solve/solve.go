package solve

import (
	"fmt"
)

func CheckValid(board [4][4]int) bool {
	boardCheck := [][]int{
		{0, 1, 5, 4},
		{3, 2, 6, 7},
		{12, 8, 9, 13},
		{15, 10, 11, 14},
		{0, 4, 8, 12},
		{1, 5, 9, 13},
		{2, 6, 10, 14},
		{3, 7, 11, 15},
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{8, 9, 10, 11},
		{12, 13, 14, 15},
	}
	// check quadrants
	for _, group := range boardCheck {
		eles := make([]int, 4)
		for _, n := range group {
			row, col := IToMatCoord(n, 4, 4)
			ele := board[row][col]
			eles = append(eles, ele)
		}
		fmt.Println("checking if", eles, "is unique")
		if !IsUnique(eles) {
			return false
		}
	}

	return true
}

func BackTrackSolve(board [4][4]int) {
	for i, row := range board {
		for j := range row {
			k := 1
			board[i][j] = k
			for !CheckValid(board) && k < 4 {
				k++
				board[i][j] = k
			}

			if !CheckValid(board) {
			}
		}
	}
}

func IsUnique(nums []int) bool {
	hashset := make(map[int]struct{})
	for _, n := range nums {
		if n == 0 {
			continue
		}
		if _, ok := hashset[n]; ok {
			return false
		}
		hashset[n] = struct{}{}
	}
	return true
}

func IToMatCoord(i, rows, cols int) (row, col int) {
	row = i / rows
	col = i % cols
	return
}

func MatCoordToI(row, col, rows, cols int) int {
	return row*cols + col
}
