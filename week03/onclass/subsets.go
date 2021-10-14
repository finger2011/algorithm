package week03onclass

// https://leetcode-cn.com/problems/subsets/
// leetcode 78 子集

func subsets(nums []int) [][]int {
	var ant = [][]int{{}}
	for i := 0; i < len(nums); i++ {
		if i < len(nums)-1 {
			var tmp = subsets(nums[i+1:])
			for j := 0; j < len(tmp); j++ {
				ant = append(ant, append([]int{nums[i]}, tmp[j]...))
			}
		} else {
			ant = append(ant, []int{nums[i]})
		}
	}
	return ant
}
