package week02onclass

import (
	"testing"
)

func Test_Matrix(t *testing.T) {
	var matrix = newMatrix([][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	})
	var sum int
	sum = matrix.sumRegion(2, 1, 4, 3)
	if sum != 8 {
		t.Errorf("sumRegion() = %v, want %v", sum, 8)
	}
	sum = matrix.sumRegion(1, 1, 2, 2)
	if sum != 11 {
		t.Errorf("sumRegion() = %v, want %v", sum, 11)
	}
	sum = matrix.sumRegion(1, 2, 2, 4)
	if sum != 12 {
		t.Errorf("sumRegion() = %v, want %v", sum, 12)
	}
}
