package week04onclass

// https://leetcode-cn.com/problems/sliding-window-maximum/
// leetcode 239. 滑动窗口最大值 最大堆

type binaryMaxHeap struct {
	heap [][]int
}

func newBinaryMaxHeap() *binaryMaxHeap {
	return &binaryMaxHeap{
		heap:[][]int{{0,0}},
	}
}

func (b *binaryMaxHeap) empty() bool {
	return len(b.heap) == 1 
}

func (b *binaryMaxHeap) insert(node []int) {
	b.heap = append(b.heap, node)
	b.heapifyUp(len(b.heap) - 1)
}

func (b *binaryMaxHeap) getMax() []int {
	return b.heap[1]
}

func (b *binaryMaxHeap) deleteMax() {
	b.heap[1] = b.heap[len(b.heap) - 1]
	b.heap = b.heap[0:len(b.heap) - 1]
	b.heapifyDown(1)
}

func (b *binaryMaxHeap) heapifyDown(i int) {
	var ch = i * 2
	for ch < len(b.heap) {
		if ch + 1 < len(b.heap) && b.heap[ch + 1][0] > b.heap[ch][0] {
			ch++
		}
		if b.heap[ch][0] > b.heap[i][0] {
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

func (b *binaryMaxHeap) heapifyUp(i int) {
	for i > 1 {
		if b.heap[i][0] > b.heap[i/2][0] {
			var tmp = b.heap[i]
			b.heap[i] = b.heap[i/2]
			b.heap[i/2] = tmp
			i /= 2
		} else {
			break
		}
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	var ans []int
	var b = newBinaryMaxHeap()
	for i := 0; i < len(nums); i++ {
		b.insert([]int{nums[i], i})
		if i >= k - 1 {
			for b.getMax()[1] <= i - k {
				b.deleteMax()
			}
			ans = append(ans, b.getMax()[0])
		}
	}
	return ans
}