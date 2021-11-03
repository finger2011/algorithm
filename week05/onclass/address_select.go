package week05onclass

// https://www.acwing.com/problem/content/description/106/
// 货仓选址

func getMinDistance(nums []int, k int) int {
	quickSort(nums, 0, k - 1)
	var ans int
	for i := 0; i < k; i++ {
		ans += abs(nums[k/2] - nums[i])
	}
	return ans
}

func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}