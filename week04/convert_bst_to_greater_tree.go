package week04

import "finger2011/algggorithm/internel"

// https://leetcode-cn.com/problems/convert-bst-to-greater-tree/
// leetcode https://leetcode-cn.com/problems/convert-bst-to-greater-tree/

var plus int

func convertBST(root *internel.TreeNode) *internel.TreeNode {
	plus = 0
	return convertBSTByPlus(root)
}

func convertBSTByPlus(root *internel.TreeNode) *internel.TreeNode {
	if root == nil {
		return root
	}
	root.Right = convertBSTByPlus(root.Right)
	root.Val += plus
	plus = root.Val
	root.Left = convertBSTByPlus(root.Left)
	return root
}
