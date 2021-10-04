package week02

// https://leetcode-cn.com/problems/number-of-submatrices-that-sum-to-target/
// leetcode 1074 元素和为目标值的子矩阵数量

//时间复杂度O(m*m*n)， m需要遍历2遍
func numSubmatrixSumTarget(matrix [][]int, target int) int {
	var ant int
	for i := 0; i < len(matrix); i++ {
		sum := make([]int, len(matrix[0]))
		for j := i; j < len(matrix); j++ {
			for k := 0; k < len(matrix[j]); k++ {
				sum[k] += matrix[j][k]
			}
			ant += subSum(sum, target)
		}
	}
	return ant
}

func subSum(nums []int, target int) int {
	var sum = make(map[int]int, len(nums))
	sum[0] = 1
	var ant int
	for i, pre := 0, 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := sum[pre-target]; ok {
			ant += sum[pre-target]
		}
		sum[pre]++
	}
	return ant
}
