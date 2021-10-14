package week03onclass

import (
	"encoding/json"
	"finger2011/algggorithm/internel"
	"strconv"
	"strings"
)

// https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/
// leetcode 297. 二叉树的序列化与反序列化

//Codec codec
type Codec struct {
	cur int
	seq []string
}

type nodeVal struct {
	Val  int
	Stop bool
}

//Constructor Constructor
func Constructor() Codec {
	return Codec{}
}

func (c *Codec) serialize2(root *internel.TreeNode) string {
	c.seq = []string{}
	c.dfs(root)
	return strings.Join(c.seq, ",")
}

func (c *Codec) deserialize2(data string) *internel.TreeNode {
	c.seq = strings.Split(data, ",")
	c.cur = 0
	return c.restore()
}
func (c *Codec) restore() *internel.TreeNode {
	if c.seq[c.cur] == "null" || c.cur >= len(c.seq) {
		c.cur++
		return nil
	}
	var val, _ = strconv.Atoi(c.seq[c.cur])
	var root = &internel.TreeNode{Val: val}
	c.cur++
	root.Left = c.restore()
	root.Right = c.restore()
	return root
}

func (c *Codec) dfs(root *internel.TreeNode) {
	if root == nil {
		c.seq = append(c.seq, "null")
		return
	}
	c.seq = append(c.seq, strconv.Itoa(root.Val))
	c.dfs(root.Left)
	c.dfs(root.Right)
	return
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *internel.TreeNode) string {
	if root == nil {
		return ""
	}
	var ans = [][]*nodeVal{[]*nodeVal{&nodeVal{Val: root.Val}, &nodeVal{Stop: true}}}
	var nodes = []*internel.TreeNode{root}
	var tmpNodes = []*internel.TreeNode{}
	for len(nodes) > 0 {
		var tmpAns = []*nodeVal{}
		for i := 0; i < len(nodes); i++ {
			if nodes[i].Left != nil {
				tmpAns = append(tmpAns, &nodeVal{Val: nodes[i].Left.Val})
				tmpNodes = append(tmpNodes, nodes[i].Left)
			} else {
				tmpAns = append(tmpAns, &nodeVal{Stop: true})
			}
			if nodes[i].Right != nil {
				tmpAns = append(tmpAns, &nodeVal{Val: nodes[i].Right.Val})
				tmpNodes = append(tmpNodes, nodes[i].Right)
			} else {
				tmpAns = append(tmpAns, &nodeVal{Stop: true})
			}
		}
		ans = append(ans, tmpAns)
		nodes = tmpNodes
		tmpNodes = []*internel.TreeNode{}
	}
	bytes, _ := json.Marshal(ans)
	return string(bytes)
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *internel.TreeNode {
	if data == "" {
		return nil
	}
	var ans [][]*nodeVal
	_ = json.Unmarshal([]byte(data), &ans)
	if len(ans) == 0 {
		return nil
	}
	var root = &internel.TreeNode{Val: ans[0][0].Val}
	var nodes = []*internel.TreeNode{root}
	var tmpNodes = []*internel.TreeNode{}
	var deep = 1
	for len(nodes) > 0 {
		for i := 0; i < len(nodes); i++ {
			if i*2 < len(ans[deep]) && !ans[deep][i*2].Stop {
				nodes[i].Left = &internel.TreeNode{Val: ans[deep][i*2].Val}
				tmpNodes = append(tmpNodes, nodes[i].Left)
			}
			if i*2+1 < len(ans[deep]) && !ans[deep][i*2+1].Stop {
				nodes[i].Right = &internel.TreeNode{Val: ans[deep][i*2+1].Val}
				tmpNodes = append(tmpNodes, nodes[i].Right)
			}
		}
		nodes = tmpNodes
		tmpNodes = []*internel.TreeNode{}
		deep++
	}

	return root
}
