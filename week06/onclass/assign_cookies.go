package week06onclasss

// https://leetcode-cn.com/problems/assign-cookies/
// leetcode 455. 分发饼干

func findContentChildren(g []int, s []int) int {
	quickSort(g, 0, len(g) - 1)
	quickSort(s, 0, len(s) - 1)
	var j int
	var ans int
	for i := 0; i < len(g); i++ {
		for j < len(s) && s[j] < g[i] {
			j++
		}
		if j < len(s) {
			ans++
			j++
		} else {
			break
		}
	}
	return ans
}

//快速排序
func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	var mid = partition(nums, left, right)
	quickSort(nums, left, mid)
	quickSort(nums, mid+1, right)
}

func partition(nums []int, left, right int) int {
	var pVal = nums[left]
	for left <= right {
		for left < len(nums) && nums[left] < pVal {
			left++
		}
		for right < len(nums) && nums[right] > pVal {
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