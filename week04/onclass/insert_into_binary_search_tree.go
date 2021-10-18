package week04onclass

import "finger2011/algggorithm/internel"

// https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/
// 701. 二叉搜索树中的插入操作

func insertIntoBST(root *internel.TreeNode, val int) *internel.TreeNode {
	if root == nil {
		root = &internel.TreeNode{Val:val}
	} else {
		if val < root.Val {
			root.Left = insertIntoBST(root.Left, val)
		} else {
			root.Right = insertIntoBST(root.Right, val)
		}
	}
	return root
}
