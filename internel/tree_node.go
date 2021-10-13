package internel

//TreeNode 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//TreeToInt 打印输出
func (t *TreeNode) TreeToInt() []int {
	var ant = []int{t.Val}
	if t.Left != nil {
		ant = append(ant, t.Left.TreeToInt()...)
	}
	if t.Right != nil {
		ant = append(ant, t.Right.TreeToInt()...)
	}
	return ant
}

//CreateTreeByInt 创建一棵树
func CreateTreeByInt(nums []int) *TreeNode {
	// var nums = []int{4, 2, 1, 3, 7, 6, 9}
	return createTree(nums)
}

//CreateTreeByIntLevel 创建一棵树
func CreateTreeByIntLevel(nums []int) *TreeNode {
	// var nums = []int{4, 2, 1, 3, 7, 6, 9}
	return createTreeByLevel(nums)
}

func createTreeByLevel(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	var root = &TreeNode{Val: nums[0]}
	var cur = 1
	var nodes = []*TreeNode{root}
	for {
		var tmpNodes = []*TreeNode{}
		for i := 0; i < len(nodes); i++ {
			if cur >= len(nums) {
				return root
			}
			if nums[cur] > MinInt {
				nodes[i].Left = &TreeNode{Val: nums[cur]}
				tmpNodes = append(tmpNodes, nodes[i].Left)
			}
			cur++
			if cur >= len(nums) {
				return root
			}
			if nums[cur] > MinInt {
				nodes[i].Right = &TreeNode{Val: nums[cur]}
				tmpNodes = append(tmpNodes, nodes[i].Right)
			}
			cur++
		}
		nodes = tmpNodes
	}
}

func createTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	var root = &TreeNode{Val: nums[0]}
	if len(nums) == 1 {
		return root
	}
	var done bool
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[0] {
			done = true
			root.Left = createTree(nums[1:i])
			if i < len(nums)-1 {
				root.Right = createTree(nums[i+1:])
			}
		}
	}
	if !done {
		root.Left = createTree(nums[1:])
	}
	return root
}
