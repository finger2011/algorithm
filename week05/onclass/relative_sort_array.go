package week05onclass

// https://leetcode-cn.com/problems/relative-sort-array/
// leetcode 1122. 数组的相对排序

func relativeSortArray(arr1 []int, arr2 []int) []int {
	var arr = make([]int, 1001)
	for i := 0; i < len(arr1); i++ {
		arr[arr1[i]]++
	}
	var ans = make([]int, len(arr1))
	var cur int
	for i := 0; i < len(arr2); i++ {
		for arr[arr2[i]] > 0 {
			ans[cur] = arr2[i]
			arr[arr2[i]]--
			cur++
		}
	}
	for i := 0; i < len(arr); i++ {
		for arr[i] > 0 {
			ans[cur] = i
			arr[i]--
			cur++
		}
	}
	return ans
}