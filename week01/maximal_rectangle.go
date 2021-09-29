package week01

//https://leetcode-cn.com/problems/maximal-rectangle/
//leetcode 85 最大矩形

func maximalRectangle(matrix [][]byte) int {
	var ant int
	var heights = make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		heights[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			if "0" == string(matrix[i][j]) {
				heights[i][j] = 0
			} else {
				if i == 0 {
					heights[i][j] = 1
				} else {
					heights[i][j] = 1 + heights[i-1][j]
				}
			}
		}
	}
	// fmt.Printf("height:%v\n", heights)
	for i := 0; i < len(heights); i++ {
		var large = LargestRectangleArea(heights[i])
		if large > ant {
			ant = large
		}
	}
	return ant
}
