package week02onclass

import (
	"fmt"
)

//https://leetcode-cn.com/problems/lru-cache/
//leetcode 146 LRU缓存机制

//LRUCache lru cache struct
type LRUCache struct {
	capacity int
	length   int
	htable   map[ValueInt]*ListNodeStruct
	head     *ListNodeStruct
	tail     *ListNodeStruct
}

//ListNodeStruct list node
type ListNodeStruct struct {
	Key  ValueInt
	Val  ValueInt
	Next *ListNodeStruct
	Prev *ListNodeStruct
}

//ValueInt int value
type ValueInt int

//Constructor create lrucache
func Constructor(capacity int) LRUCache {
	var c = LRUCache{
		capacity: capacity,
		htable:   make(map[ValueInt]*ListNodeStruct, capacity),
		head:     new(ListNodeStruct),
		tail:     new(ListNodeStruct),
	}
	c.head.Next = c.tail
	c.tail.Prev = c.head
	return c
}

//Get get cache by key 这里方法中与leetcode不同，把this改为了c
func (c *LRUCache) Get(key int) int {
	if c.htable[ValueInt(key)] == nil {
		return -1
	}
	c.activeKey(c.htable[ValueInt(key)], false)
	return int(c.htable[ValueInt(key)].Val)
}

//Put put cache key
func (c *LRUCache) Put(key int, value int) {
	if c.htable[ValueInt(key)] != nil {
		c.htable[ValueInt(key)].Val = ValueInt(value)
		c.activeKey(c.htable[ValueInt(key)], false)
	} else {
		var node = new(ListNodeStruct)
		node.Val = ValueInt(value)
		node.Key = ValueInt(key)
		c.htable[ValueInt(key)] = node
		c.activeKey(node, true)
		c.length++
	}
}

func (c *LRUCache) printValues() {
	cur := c.head.Next
	fmt.Println("print cache:")
	for {
		fmt.Printf("key:%v ====> value:%v\n", cur.Key, cur.Val)
		cur = cur.Next
		if cur == nil || cur.Next == nil {
			break
		}
	}
}

func (c *LRUCache) activeKey(node *ListNodeStruct, insert bool) error {
	if !insert {
		var prev = node.Prev
		var next = node.Next
		prev.Next = next
		next.Prev = prev
	} else {
		if c.length == c.capacity && c.length > 0 {
			c.htable[c.head.Next.Key] = nil
			c.head.Next.Next.Prev = c.head
			c.head.Next = c.head.Next.Next
			c.length--
		}
	}
	var prev = c.tail.Prev
	prev.Next = node
	node.Next = c.tail
	node.Prev = prev
	c.tail.Prev = node
	return nil
}
