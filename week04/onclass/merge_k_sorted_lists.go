package week04onclass

import "finger2011/algggorithm/internel"

// https://leetcode-cn.com/problems/merge-k-sorted-lists/
// leetcode 23. 合并K个升序链表 堆解法

type binaryHeap struct {
	heap []*internel.ListNode
}

func newBinaryHead() *binaryHeap {
	return &binaryHeap{
		heap:[]*internel.ListNode{&internel.ListNode{}},
	}
}

func (b *binaryHeap) empty() bool {
	return len(b.heap) == 1 
}

func (b *binaryHeap) insert(node *internel.ListNode) {
	b.heap = append(b.heap, node)
	b.heapifyUp(len(b.heap) - 1)
}

func (b *binaryHeap) getMin() *internel.ListNode {
	return b.heap[1]
}

func (b *binaryHeap) deleteMin() {
	b.heap[1] = b.heap[len(b.heap) - 1]
	b.heap = b.heap[0:len(b.heap) - 1]
	b.heapifyDown(1)
}

func (b *binaryHeap) heapifyDown(i int) {
	var ch = i * 2
	for ch < len(b.heap) {
		if ch + 1 < len(b.heap) && b.heap[ch + 1].Val < b.heap[ch].Val {
			ch++
		}
		if b.heap[ch].Val < b.heap[i].Val {
			var tmp = b.heap[i]
			b.heap[i] = b.heap[ch]
			b.heap[ch] = tmp
			i = ch
			ch = i * 2
		} else {
			break
		}
	}
}

func (b *binaryHeap) heapifyUp(i int) {
	for i > 1 {
		if b.heap[i].Val < b.heap[i/2].Val {
			var tmp = b.heap[i]
			b.heap[i] = b.heap[i/2]
			b.heap[i/2] = tmp
			i /= 2
		} else {
			break
		}
	}
}

func mergeKLists(lists []*internel.ListNode) *internel.ListNode {
	var head = &internel.ListNode{}
	var tail = head
	var b = newBinaryHead()
	for _, node := range lists {
		if node != nil {
			b.insert(node)
		}
	}
	for !b.empty() {
		var node = b.getMin()
		b.deleteMin()
		tail.Next = node
		if node.Next != nil {
			b.insert(node.Next)
		}
		tail = tail.Next
	}
	return head.Next
}
