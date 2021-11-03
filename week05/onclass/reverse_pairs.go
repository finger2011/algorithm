package week05onclass

// https://leetcode-cn.com/problems/reverse-pairs/
// leetcode 493. 翻转对
var ans int
var pairNum []int
func reversePairs(nums []int) int {
	ans = 0
	pairNum = nums
	mergeSort1(0, len(nums) - 1)
	return ans
}

func calPairs(left, right, mid int) {
	var r = mid
	for l := left; l <= mid; l++ {
		for r < right && pairNum[l] > 2 * pairNum[r + 1] {
			r++
		}
		ans += r - mid
	}
}

func mergeSort1(left, right int) {
	if left >= right {
		return
	}
	var mid = (left + right) / 2
	mergeSort1(left, mid)
	mergeSort1(mid+1, right)
	calPairs(left, right, mid)
	merge1(left, mid, right)
}

func merge1(left, mid, right int) {
	var tmp = make([]int, right-left+1)
	var i, j = left, mid + 1
	for k := 0; k < len(tmp); k++ {
		if j > right || (i <= mid && pairNum[i] <= pairNum[j]) {
			tmp[k] = pairNum[i]
			i++
		} else {
			tmp[k] = pairNum[j]
			j++
		}
	}
	for k := 0; k < len(tmp); k++ {
		pairNum[left+k] = tmp[k]
	}
}