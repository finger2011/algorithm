package week02onclass

// https://leetcode-cn.com/problems/container-with-most-water/
// leetcode 11. 盛最多水的容器

func maxArea(height []int) int {
	var ant int
	left, right := 0, len(height)-1
	for left < right {
		var minH int
		var area = right - left
		if height[left] < height[right] {
			minH = height[left]
			for left < right && height[left] == height[left+1] {
				left++
			}
			left++

		} else {
			minH = height[right]
			for left < right && height[right] == height[right-1] {
				right--
			}
			right--
		}
		area *= minH
		if area > ant {
			ant = area
		}
	}
	return ant
}
