package week01

//https://leetcode-cn.com/problems/maximal-rectangle/
//leetcode 85 最大矩形


//思路：
//row1 1 0 1 0 0            height1 1 0 1 0 0   ==>  利用单调栈取到最大矩形 max1
//row2 1 0 1 1 1  -----\    height2 2 0 2 1 1   ==>  利用单调栈取到最大矩形 max2
//row3 1 1 1 1 1  -----/    height3 3 1 3 2 2   ==>  利用单调栈取到最大矩形 max3
//row4 1 0 0 1 0            height4 4 0 0 3 0   ==>  利用单调栈取到最大矩形 max4

// 由于0的存在导致不连续，所以把问题拆解为先从每一行获取最大矩形，最后汇总得到所有行的最大矩形
// 例如针对row1获取以row1为基础的最大矩形，然后是row2,row3,row4
// 由于需要取矩形的最大值，考虑用单调栈
// 如果该行某列上的值为0， 则在这一行计算最大矩形的逻辑中，该列的高度height = 0
// 根据以上逻辑生成行高的二维数组
// 针对行高的每一行，调用单调栈(leetcode 84)代码获取在该行下的最大矩形面积
// 取所有行的最大值，即可得出最终结果

func maximalRectangle(matrix [][]byte) int {
	var ant int
	var heights = make([][]int, len(matrix))
	//生成行高二维数组
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
		//这里直接调用单调栈最大矩形方法
		var large = LargestRectangleArea(heights[i])
		if large > ant {
			ant = large
		}
	}
	return ant
}
