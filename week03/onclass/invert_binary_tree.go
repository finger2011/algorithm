package week03onclass

import (
	"finger2011/algggorithm/internel"
)

// https://leetcode-cn.com/problems/invert-binary-tree/description/
// leetcode 226. 翻转二叉树
func invertTree(root *internel.TreeNode) *internel.TreeNode {
	if root != nil {
		if root.Left != nil {
			if root.Right != nil {
				var tmp = root.Right
				root.Right = root.Left
				root.Left = tmp
			} else {
				root.Right = root.Left
				root.Left = nil
			}
		} else if root.Right != nil {
			root.Left = root.Right
			root.Right = nil
		}
		if root.Left != nil {
			_ = invertTree(root.Left)
		}
		if root.Right != nil {
			_ = invertTree(root.Right)
		}
	}
	return root
}
