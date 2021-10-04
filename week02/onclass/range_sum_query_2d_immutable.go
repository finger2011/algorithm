package week02onclass

// https://leetcode-cn.com/problems/range-sum-query-2d-immutable/
// leetcode 304 二维区域和检索 - 矩阵不可变

//NumMatrix num matrix
type NumMatrix struct {
	matrix [][]int
}

func newMatrix(matrix [][]int) NumMatrix {
	var newMatrix = make([][]int, len(matrix) + 1)
	newMatrix[0] = make([]int, len(matrix[0]) + 1)
	for i := 0; i < len(matrix); i++ {
		newMatrix[i + 1] = make([]int, len(matrix[i]) + 1)
		for j := 0; j < len(matrix[i]); j++ {
			newMatrix[i + 1][j + 1] = newMatrix[i + 1][j] + newMatrix[i][j + 1] - newMatrix[i][j] + matrix[i][j]
		}
	}
	return NumMatrix{matrix: newMatrix}
}

func (matrix *NumMatrix) sumRegion(row1 int, col1 int, row2 int, col2 int) int {
	row1++
	row2++
	col1++
	col2++
	return matrix.matrix[row2][col2] - matrix.matrix[row1 - 1][col2] - matrix.matrix[row2][col1 - 1] + matrix.matrix[row1 - 1][col1 - 1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
