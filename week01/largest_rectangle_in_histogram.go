package week01

//https://leetcode-cn.com/problems/largest-rectangle-in-histogram/
//leetcode 84
//柱状图中的最大矩形

//Rect rect
type Rect struct {
    width int
    height int
}

//LargestRectangleArea func
func LargestRectangleArea(heights []int) int {
    var ant int
    var stack []*Rect
    heights = append(heights, 0)
    for i := 0; i < len(heights); i++ {
        var width int
        for len(stack) > 0 && stack[len(stack) - 1].height >= heights[i] {
            width += stack[len(stack) - 1].width
            if ant < stack[len(stack) - 1].height * width {
                ant = stack[len(stack) - 1].height * width
            }
            stack = stack[0:(len(stack) - 1)]
        }
        stack = append(stack, &Rect{width:width + 1, height:heights[i]})
    }
    return ant
}