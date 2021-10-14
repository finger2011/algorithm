package week03onclass

// https://leetcode-cn.com/problems/permutations/
// leetcode 46 全排列

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}
	var ant [][]int
	for i := 0; i < len(nums); i++ {
		var tmp = append([]int{}, nums[0:i]...)
		if i < len(nums) - 1 {
			tmp = append(tmp, nums[i + 1:]...)
		}
		var tmpAnt = permute(tmp)
		for j := 0; j < len(tmpAnt); j++ {
			ant = append(ant, append([]int{nums[i]}, tmpAnt[j]...))
		}
	}
	return ant
}
