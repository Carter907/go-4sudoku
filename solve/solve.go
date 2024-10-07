package solve

func Solve(board [4][4]int, row, col int) bool {
	// check quadrants

	if row == 3 && col == 4 {
		return true
	}
	if col == 4 {
		row++
		col = 0
	}

	if board[row][col] != 0 {
		Solve(board, row, col+1)
	}

	for i := range 4 {
		if IsPossible(board, row, col, 1) {
			board[row][col] = i + 1
			if Solve(board, row, col+1) {
				return true
			}
		}
		board[row][col] = 0
	}

	return false
}

func IsPossible(board [4][4]int, row, col, e int) bool {
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
	startRow, startCol := row-row%2, col-col%2

	// traverse assuming 2 by 2 subgrid
	for i := startRow; i < 2; i++ {
		for j := startCol; j < 2; j++ {
			if board[i][j] == e {
				return false
			}
		}
	}

	return true
}
