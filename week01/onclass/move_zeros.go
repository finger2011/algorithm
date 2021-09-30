package week01onclass

// https://leetcode-cn.com/problems/move-zeroes/
// leetcode 283

func moveZeroes(nums []int) []int {
	var cur int
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			//此处可以做个优化，在i > cur时才进行交换，减少了在最开始i==cur时的值比较
			nums[cur], nums[i] = nums[i], nums[cur]
			cur++
		}
	}
	return nums
}
