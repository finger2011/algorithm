package week05

// https://leetcode-cn.com/problems/search-a-2d-matrix/
// leetcode 74. 搜索二维矩阵

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	var left, right = 0, len(matrix) * len(matrix[0]) - 1
	for left <= right {
		var mid = (left + right + 1) / 2
		var row, col = mid / len(matrix[0]), mid % len(matrix[0])
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}