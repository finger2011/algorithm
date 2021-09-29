package week01onclass

//https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
//leetcode 26

//返回参数做了修改，方便测试
func removeDuplicates(nums []int) ([]int, int) {
    if len(nums) <= 1 {
        return nums, len(nums)
    }
    var cur = 0
    for i := 1; i < len(nums); i++ {
        if nums[cur] != nums[i] {
            cur++
            nums[cur] = nums[i]        
        }
    }
    return nums[0:cur + 1], cur + 1
}

//题解中看到比较有意思的想法
//如果不是使每个元素只出现一次，而是最多保留k次

func removeDuplicatesByK(nums []int, k int) ([]int, int) {
	if len(nums) <= 1 {
        return nums, len(nums)
    }
    var cur = 0
    for i := 1; i < len(nums); i++ {
        if i < k || nums[cur - k + 1] != nums[i] {
			//前k位可以直接保留
			//超过k时，判断cur往前走k位的数是否和当前数相同，相同则删除
            cur++
            nums[cur] = nums[i]        
        }
    }
    return nums[0:cur + 1], cur + 1
}