package solve_test

import (
	"testing"

	"github.com/Carter907/go-4sodoku/solve"
)

func TestSolve(t *testing.T) {
	if row, col := solve.IToMatCoord(7, 4, 4); row != 1 || col != 3 {
		t.Fail()
	} else {
		t.Log("row:", row, "col:", col)
	}
	row, col := solve.IToMatCoord(7, 4, 4)
	if i := solve.MatCoordToI(row, col, 4, 4); i != 7 {
		t.Fail()
	} else {
		t.Log("i:", 7)
	}
}

func TestUnique(t *testing.T) {
	if !solve.IsUnique([]int{1, -1, -1, 4}) {
		t.Fatal("whatt")
	}
}
