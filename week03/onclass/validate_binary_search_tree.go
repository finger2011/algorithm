package week03onclass

import (
	"finger2011/algggorithm/internel"
)

// https://leetcode-cn.com/problems/validate-binary-search-tree/
// leetcode 98. 验证二叉搜索树

func isValidBST(root *internel.TreeNode) bool {
	return validBST(root, internel.MinInt, internel.MaxInt)
}

func validBST(root *internel.TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left != nil {
		if root.Left.Val >= root.Val || root.Left.Val <= min {
			return false
		}
	}
	if root.Right != nil {
		if root.Right.Val <= root.Val || root.Right.Val >= max {
			return false
		}
	}
	return validBST(root.Left, min, root.Val) && validBST(root.Right, root.Val, max)
}
