package week06onclasss

// https://leetcode-cn.com/problems/jump-game-ii/
// leetcode 45. 跳跃游戏 II

func jump(nums []int) int {
	var cur, max, ans int
	for cur < len(nums) - 1 {
		max = cur + nums[cur]
		if max >= len(nums) - 1 {
			return ans + 1
		}
		next := cur
		for i := cur + 1; i <= max; i++ {
			if i + nums[i] > next {
				next = i + nums[i]
				cur = i
			}
		}
		ans++
	}
	return ans
}