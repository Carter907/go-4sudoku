package solve

import (
	"math"
	"math/rand"

	"github.com/Carter907/go-4sodoku/consts"
)

const (
	Rows = consts.Rows
	Cols = consts.Cols
)

func Solve(board *[Rows][Cols]int, row, col int) bool {
	// check quadrants

	if row == Rows-1 && col == Cols {
		return true
	}

	if col == Cols {
		row++
		col = 0
	}

	if board[row][col] > 0 {
		Solve(board, row, col+1)
	}

	for range Rows {
		if num := rand.Intn(Rows) + 1; IsPossible(*board, row, col, num) {
			board[row][col] = num
			if Solve(board, row, col+1) {
				return true
			}
		}
		board[row][col] = 0
	}

	return false
}

func Sparcify(board *[Rows][Cols]int, amount float32) {
	for range int32(amount * float32(Rows*Rows)) {
		board[rand.Intn(Rows)][rand.Intn(Cols)] = 0
	}
}

func IsPossible(board [Rows][Cols]int, row, col, e int) bool {
	// check for element in similar row
	for _, num := range board[row] {
		if num == e {
			return false
		}
	}

	for _, rows := range board {
		if rows[col] == e {
			return false
		}
	}

	// start row and column of subgrid
	startRow, startCol := row-row%int(math.Sqrt(Rows)), col-col%int(math.Sqrt(Cols))

	// traverse assuming 2 by 2 subgrid
	for i := 0; i < int(math.Sqrt(Rows)); i++ {
		for j := 0; j < int(math.Sqrt(Cols)); j++ {
			if board[i+startRow][j+startCol] == e {
				return false
			}
		}
	}

	return true
}
