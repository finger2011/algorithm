package week03onclass

import (
	"finger2011/algggorithm/internel"
)

// https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/
// leetcode 429. N 叉树的层序遍历

func levelOrder(root *internel.NodeTree) [][]int {
	if root == nil {
		return [][]int{}
	}
	var ans = [][]int{}
	var nodes = []*internel.NodeTree{root}
	var tmpNodes = []*internel.NodeTree{}
	for len(nodes) > 0 {
		var tmp []int
		for i := 0; i < len(nodes); i++ {
			tmp = append(tmp, nodes[i].Val)
			if len(nodes[i].Children) > 0 {
				tmpNodes = append(tmpNodes, nodes[i].Children...)
			}
		}
		nodes = tmpNodes
		tmpNodes = []*internel.NodeTree{}
		ans = append(ans, tmp)
	}
	return ans
}
