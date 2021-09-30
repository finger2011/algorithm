package week01onclass

// https://leetcode-cn.com/problems/sliding-window-maximum/
// leetcode 239 滑动窗口最大值

//
func maxSlidingWindow(nums []int, k int) []int {
	if k == 1 {
		return nums
	}
	var ant []int
	var queue []int
	for i := 0; i < len(nums); i++ {
		for len(queue) > 0 && queue[0] <= i-k {
			queue = queue[1:]
		}
		for len(queue) > 0 && nums[queue[len(queue)-1]] <= nums[i] {
			queue = queue[0:(len(queue) - 1)]
		}
		queue = append(queue, i)
		if i >= k-1 {
			ant = append(ant, nums[queue[0]])
		}
	}
	return ant
}

//queue中不存位置，直接存值，减少nums寻址次数
func maxSlidingWindow2(nums []int, k int) []int {
	if k == 1 {
		return nums
	}
	var ant []int
	var queue []int
	for i := 0; i < len(nums); i++ {
		//queue[len(queue)-1] < nums[i]，不能取等值，防止出现相同的最大值
		for len(queue) > 0 && queue[len(queue)-1] < nums[i] {
			queue = queue[0:(len(queue) - 1)]
		}
		queue = append(queue, nums[i])

		if i >= k {
			//出界时，判断出界的值是否是最大值，是的话清除掉
			if queue[0] == nums[i-k] {
				queue = queue[1:]
			}
		}
		if i >= k-1 {
			ant = append(ant, queue[0])
		}
	}
	return ant
}
