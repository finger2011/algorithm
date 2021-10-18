package week04onclass

import "finger2011/algggorithm/internel"

// https://leetcode-cn.com/problems/successor-lcci/
// leecode 面试题 04.06. 后继者

func inorderSuccessor(root *internel.TreeNode, p *internel.TreeNode) *internel.TreeNode {
	return getNext(root, p.Val)
}

func getNext(root *internel.TreeNode, val int) *internel.TreeNode {
	var ans *internel.TreeNode
	var cur = root
	for cur != nil {
		if cur.Val == val {
			if cur.Right != nil {
				ans = cur.Right
				for ans.Left != nil {
					ans = ans.Left
				}
			}
			break
		}
		if cur.Val < val {
			cur = cur.Right
		} else {
			if ans == nil || ans.Val > cur.Val {
				ans = cur
			}
			cur = cur.Left
		}
	}
	return ans
}
