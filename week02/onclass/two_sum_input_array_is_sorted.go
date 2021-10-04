package week02onclass

// https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/
// leetcode 167. 两数之和 II - 输入有序数组

func twoSumUp(numbers []int, target int) []int {
	for i, j := 0, len(numbers)-1; i < len(numbers); i++ {
		for i < j && numbers[i]+numbers[j] > target {
			j--
		}
		if i < j && numbers[i]+numbers[j] == target {
			return []int{i + 1, j + 1}
		}
	}
	return []int{}
}
