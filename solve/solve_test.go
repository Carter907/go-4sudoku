package solve_test

import (
	"testing"

	"github.com/Carter907/go-4sodoku/solve"
)

func TestSolve(t *testing.T) {
	board := [4][4]int{
		{1, 2, 0, 3},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	solve.Solve(board, 0, 0)
	t.Log(board)
}

func TestRange(t *testing.T) {
	for i := range 4 {
		t.Log(i)
	}
}

func TestIsPossible(t *testing.T) {
	board := [4][4]int{
		{1, 2, 0, 3},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	if !solve.IsPossible(board, 1, 1, 3) {
		t.Fail()
	}
}
