package week05onclass

// https://leetcode-cn.com/problems/merge-intervals/
// leetcode 56. 合并区间

// 差分
func mergeIntervals2(intervals [][]int) [][]int {
	var diff = make([][]int, len(intervals) * 2)
	for i := 0; i < len(intervals); i++ {
		diff[i * 2] = []int{intervals[i][0], 1}
		diff[i * 2 + 1] = []int{intervals[i][1] + 1, -1}
	}
	quickSort2(diff, 0, len(diff) - 1)
	var cur, start int
	var ans [][]int
	for i := 0; i < len(diff); i++ {
		if cur == 0 {
			start = diff[i][0]
		}
		cur += diff[i][1]
		if cur == 0 {
			ans = append(ans, []int{start, diff[i][0] - 1})
		}
	}
	return ans
}


// 排序后遍历
func mergeIntervals(intervals [][]int) [][]int {
	quickSort2(intervals, 0, len(intervals)-1)
	var start, end = -1, -1
	var ans [][]int
	for i := 0; i < len(intervals); i++ {
		if intervals[i][0] <= end {
			if intervals[i][1] > end {
				end = intervals[i][1]
			}
		} else {
			if -1 != end {
				ans = append(ans, []int{start, end})
			}
			start = intervals[i][0]
			end = intervals[i][1]
		}
	}
	ans = append(ans, []int{start, end})
	return ans
}

func quickSort2(nums [][]int, left, right int) {
	if left >= right {
		return
	}
	var mid = partition2(nums, left, right)
	quickSort2(nums, left, mid)
	quickSort2(nums, mid+1, right)
}

func partition2(nums [][]int, left, right int) int {
	var pVal = nums[left]
	for left <= right {
		for left < len(nums) && (nums[left][0] < pVal[0] || (nums[left][0] == pVal[0] && nums[left][1] < pVal[1])) {
			left++
		}
		for right < len(nums) && (nums[right][0] > pVal[0] || (nums[right][0] == pVal[0] && nums[right][1] > pVal[1])) {
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
