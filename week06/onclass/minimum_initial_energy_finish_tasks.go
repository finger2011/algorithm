package week06onclasss

// https://leetcode-cn.com/problems/minimum-initial-energy-to-finish-tasks/
// leetcode 1665. 完成所有任务的最少初始能量

func minimumEffort(tasks [][]int) int {
	quickSortEffort(tasks, 0, len(tasks) - 1)
	var ans int
	for i := len(tasks) - 1; i >= 0; i-- {
		if tasks[i][1] > ans + tasks[i][0] {
			ans = tasks[i][1]
		} else {
			ans = ans + tasks[i][0]
		}
	}
	return ans
}

func compare(a, b []int) bool {
	return a[0] - a[1] < b[0] - b[1]
}

//快速排序
func quickSortEffort(nums [][]int, left, right int) {
	if left >= right {
		return
	}
	var mid = partitionEffort(nums, left, right)
	quickSortEffort(nums, left, mid)
	quickSortEffort(nums, mid+1, right)
}

func partitionEffort(nums [][]int, left, right int) int {
	var pVal = nums[left]
	for left <= right {
		for left < len(nums) && compare(nums[left], pVal) {
			left++
		}
		for right < len(nums) && compare(pVal, nums[right]) {
			right--
		}
		if left == right {
			break
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
	return right
}