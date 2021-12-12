package week10

// https://leetcode-cn.com/problems/sliding-window-maximum/
// leetcode 239. 滑动窗口最大值

func maxSlidingWindow(nums []int, k int) []int {
	var ans []int
	var b = newAvlTree()
	for i := 0; i < len(nums); i++ {
		b = b.insert(nums[i])
		if i >= k-1 {
			ans = append(ans, b.getMax())
			b = b.delete(nums[i-k+1])
		}
	}
	return ans
}
