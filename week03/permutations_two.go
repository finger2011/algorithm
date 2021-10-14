package week03

// https://leetcode-cn.com/problems/permutations-ii/
// leetcode 47 全排列 II

func permuteUnique(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}
	var ant [][]int
	var repeat = make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if _, ok := repeat[nums[i]]; ok {
			continue
		} else {
			repeat[nums[i]] = 1
		}
		var tmp = append([]int{}, nums[0:i]...)
		if i < len(nums)-1 {
			tmp = append(tmp, nums[i+1:]...)
		}
		var tmpAnt = permuteUnique(tmp)
		for j := 0; j < len(tmpAnt); j++ {
			ant = append(ant, append([]int{nums[i]}, tmpAnt[j]...))
		}
	}
	return ant
}
