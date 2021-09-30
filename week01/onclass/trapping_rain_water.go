package week01onclass

// "fmt"

// https://leetcode-cn.com/problems/trapping-rain-water/
// leetcode 42 接雨水

//Rect rect
type Rect struct {
	width  int
	height int
}

func trap(heights []int) int {
	var ant int
	var stack []Rect
	for i := 0; i < len(heights); i++ {
		var width = 0
		for len(stack) > 0 && stack[len(stack)-1].height <= heights[i] {
			if len(stack) == 1 {
				stack = stack[0:(len(stack) - 1)]
				continue
			}
			width += stack[len(stack)-1].width
			minHeight := min(heights[i], stack[len(stack)-2].height) - stack[len(stack)-1].height
			ant += width * minHeight
			stack = stack[0:(len(stack) - 1)]
		}
		stack = append(stack, Rect{width: width + 1, height: heights[i]})
	}
	return ant
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
