package week10

import (
	"math/rand"
)

// https://leetcode-cn.com/problems/design-skiplist/
// leetcode 1206. 设计跳表

const maxLevel = 10

type slNode struct {
	val      int
	cnt      int
	maxLevel int
	levels   []*slNode
}

type skiplist struct {
	head   *slNode
	tail   *slNode
	level  int
	length int
}

func constructor() skiplist {
	s := skiplist{
		head: &slNode{
			levels: make([]*slNode, maxLevel),
			val:    -1e9,
		},
		tail: &slNode{
			levels: make([]*slNode, maxLevel),
			val:    1e9,
		},
		level:maxLevel,
	}
	for i := 0; i < s.level; i++{
		s.head.levels[i] = s.tail
	}
	return s
}

func (sl *skiplist) search(target int) bool {
	return sl.searchTarget(target).val == target
}

func (sl *skiplist) searchTarget(target int) *slNode {
	cur := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for cur.levels[i] != nil && cur.levels[i].val < target {
			cur = cur.levels[i]
		}
	}
	cur = cur.levels[0]
	return cur
}

func (sl *skiplist) add(num int) {
	node := sl.searchTarget(num)
	if node.val == num {
		node.cnt++
		return
	}
	newNode := newSkiplistNode(num)
	cur := sl.head
	for i := newNode.maxLevel - 1; i >= 0; i-- {
		for cur.levels[i] != nil && cur.levels[i].val < num {
			cur = cur.levels[i]
		}
		next := cur.levels[i]
		cur.levels[i] = newNode
		newNode.levels[i] = next
	}
}

func (sl *skiplist) erase(num int) bool {
	node := sl.searchTarget(num)
	if node.val != num {
		return false
	}
	if node.cnt > 1 {
		node.cnt--
		return true
	}
	
	for i := node.maxLevel - 1; i >= 0 ; i-- {
		cur := sl.head
		for cur.levels[i] != nil && cur.levels[i].val != num {
			cur = cur.levels[i]
		}
		next := node.levels[i]
		cur.levels[i] = next
	}
	return true
}

func newSkiplistNode(num int) *slNode {
	l := dial()
	return &slNode{
		val:      num,
		cnt:      1,
		maxLevel: l,
		levels:   make([]*slNode, l),
	}
}

func dial() int {
	level := 1
	for (rand.Int()&1) == 1 && level < maxLevel {
		level++
	}
	return level
}

/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */
