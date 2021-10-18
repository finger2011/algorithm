package internel

//NodeTree 多叉树
type NodeTree struct {
	Val      int
	Children []*NodeTree
}

//PrintNodeTree 打印输出数组
func (node *NodeTree) PrintNodeTree() []int {
	if node == nil {
		return []int{}
	}
	var ans = []int{node.Val, MinInt}
	var nodes = []*NodeTree{node}
	var tmpNodes = []*NodeTree{}
	for len(nodes) > 0 {
		for i := 0; i < len(nodes); i++ {
			if len(nodes[i].Children) > 0 {
				for j := 0; j < len(nodes[i].Children); j++ {
					tmpNodes = append(tmpNodes, nodes[i].Children[j])
					ans = append(ans, nodes[i].Children[j].Val)
				}
			}
			ans = append(ans, MinInt)
		}
		nodes = tmpNodes
		tmpNodes = []*NodeTree{}
	}

	return ans
}

//CreateNodeTree 数组转多叉树
func CreateNodeTree(nums []int) *NodeTree {
	return createNodeTree(nums)
}

//1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14
func createNodeTree(nums []int) *NodeTree {
	if len(nums) == 0 {
		return nil
	}
	var root = &NodeTree{Val: nums[0]}
	var cur = 2
	var nodes = []*NodeTree{root}
	var tmpNodes = []*NodeTree{}
	for cur < len(nums) {
		for i := 0; i < len(nodes); i++ {
			for ;cur < len(nums); cur++ {
				if nums[cur] == MinInt {
					cur++
					break
				}
				var newNode = &NodeTree{Val:nums[cur]}
				nodes[i].Children = append(nodes[i].Children, newNode)
				tmpNodes = append(tmpNodes, newNode)
			}
		}
		nodes = tmpNodes
		tmpNodes = []*NodeTree{}
	}
	return root
}

//TreeNode 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//TreeToInt 打印输出
func (t *TreeNode) TreeToInt() []int {
	if t == nil {
		return []int{}
	}
	var ant = []int{t.Val}
	if t.Left != nil {
		ant = append(ant, t.Left.TreeToInt()...)
	}
	if t.Right != nil {
		ant = append(ant, t.Right.TreeToInt()...)
	}
	return ant
}
//ValidBST 判断是否是平衡二叉树
func (t *TreeNode) ValidBST() bool {
	if t == nil {
		return true
	}
	if t.Left == nil && t.Right == nil {
		return true
	}
	if t.Left == nil && t.Right != nil {
		return t.Right.ValidBST()
	}
	if t.Left != nil && t.Right == nil {
		return t.Left.ValidBST()
	}
	return t.Left.ValidBST() && t.Right.ValidBST()
}
//ContainVal check if contain val
func (t *TreeNode) ContainVal(val int) bool {
	if t == nil {
		return false
	}
	if t.Val == val {
		return true
	}
	if t.Val < val {
		return t.Right.ContainVal(val)
	} else {
		return t.Left.ContainVal(val)
	}
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
