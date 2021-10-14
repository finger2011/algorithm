package week03onclass

// https://leetcode-cn.com/problems/combinations/
// leetcode 77 组合

func combine(n int, k int) [][]int {
	if n < k {
		return [][]int{}
	}
	var nums = make([]int, n)
	for i := 1; i <= n; i++ {
		nums[i-1] = i
	}
	return subCombine(nums, k)
}

func subCombine(nums []int, k int) [][]int {
	var ant [][]int
	if k == 1 {
		for i := 0; i < len(nums); i++ {
			ant = append(ant, []int{nums[i]})
		}
		return ant
	}
	for i := 0; i <= len(nums)-k; i++ {
		var tmp = subCombine(nums[i+1:], k-1)
		for j := 0; j < len(tmp); j++ {
			ant = append(ant, append([]int{nums[i]}, tmp[j]...))
		}
	}
	return ant
}
