package week03onclass

import (
	"finger2011/algggorithm/internel"
)

// https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/description/
// leetcode 589. N 叉树的前序遍历

func preorder(root *internel.NodeTree) []int {
    if root == nil {
		return []int{}
	}
	var ans = []int{root.Val}
	if len(root.Children) > 0 {
		for i := 0; i < len(root.Children); i++ {
			ans = append(ans, preorder(root.Children[i])...)
		}
	}
	return ans
}