package week01onclass

// https://leetcode-cn.com/problems/min-stack/
// leetcode 155 最小栈

//另一种解题思路：只用一个栈，栈中每个元素同时保存栈顶元素和最小值，时间和空间复杂度是一样的

//MinStack 最小栈
type MinStack struct {
    stack []int
    minStack []int
}

//NewStack 创建新栈
func NewStack() MinStack {
    return MinStack{}
}

//Push 将元素 x 推入栈中
func (stack *MinStack) Push(val int)  {
    stack.stack = append(stack.stack,  val)
    if len(stack.minStack) == 0 {
        stack.minStack = append(stack.minStack, val)
    } else {
        if val < stack.minStack[len(stack.minStack) - 1] {
            stack.minStack = append(stack.minStack, val)
        } else {
            stack.minStack = append(stack.minStack, stack.minStack[len(stack.minStack) - 1])
        }
    }
}

//Pop 删除栈顶的元素
func (stack *MinStack) Pop()  {
    if len(stack.stack) == 0 {
        panic("stack is empty")
    }
    stack.stack = stack.stack[0:(len(stack.stack) - 1)]
    stack.minStack = stack.minStack[0:(len(stack.minStack) - 1)]
}

//Top 获取栈顶元素
func (stack *MinStack) Top() int {
    if len(stack.stack) == 0 {
        panic("stack is empty")
    }
    return stack.stack[len(stack.stack) - 1]
}

//GetMin 检索栈中的最小元素
func (stack *MinStack) GetMin() int {
    if len(stack.stack) == 0 {
        panic("stack is empty")
    }
    return stack.minStack[len(stack.minStack) - 1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */