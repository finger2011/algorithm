package week01

//https://leetcode-cn.com/problems/design-circular-deque/
//leetcode 641 设计循环双端队列

//MyCircularDeque deque
type MyCircularDeque struct {
	myQueue  []int
	capacity int
}

//Constructor new deque wwith capacity
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		capacity: k,
	}
}

//InsertFront insert front,
//false when deque is full
func (deque *MyCircularDeque) InsertFront(value int) bool {
	if len(deque.myQueue) == deque.capacity {
		return false
	}
	deque.myQueue = append([]int{value}, deque.myQueue...)
	return true
}

//InsertLast insert last
//false when deque is full
func (deque *MyCircularDeque) InsertLast(value int) bool {
	if len(deque.myQueue) == deque.capacity {
		return false
	}
	deque.myQueue = append(deque.myQueue, value)
	return true
}

//DeleteFront delete front
// false when deque is empty
func (deque *MyCircularDeque) DeleteFront() bool {
	if len(deque.myQueue) == 0 {
		return false
	}
	deque.myQueue = deque.myQueue[1:]
	return true
}

//DeleteLast delete last
// false when deque is empty
func (deque *MyCircularDeque) DeleteLast() bool {
	if len(deque.myQueue) == 0 {
		return false
	}
	deque.myQueue = deque.myQueue[0:(len(deque.myQueue) - 1)]
	return true
}

//GetFront get front
// -1 when deque is empty
func (deque *MyCircularDeque) GetFront() int {
	if len(deque.myQueue) == 0 {
		return -1
	}
	return deque.myQueue[0]
}

//GetRear get rear
//-1 when deque is empty
func (deque *MyCircularDeque) GetRear() int {
	if len(deque.myQueue) == 0 {
		return -1
	}
	return deque.myQueue[len(deque.myQueue)-1]
}

//IsEmpty deque is empty
func (deque *MyCircularDeque) IsEmpty() bool {
	return len(deque.myQueue) == 0
}

//IsFull deque is full
func (deque *MyCircularDeque) IsFull() bool {
	return len(deque.myQueue) == deque.capacity
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
