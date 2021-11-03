package internel

import (
	"errors"
	"fmt"
	"reflect"
)

var _ heapInterface = &heap{}

// 堆实现
type heapInterface interface {
	top() (interface{}, error)
	empty() bool
	size() int
	push(val interface{}) error
	pop() (interface{}, error)
	// swap() error
	compare(val1, val2 interface{}) (bool, error)
}

type comparable func(val1, val2 interface{}) (bool, error)

var errHeapEmpty = errors.New("heap is empty")
var errCompareableEmpty = errors.New("comparable function is empty")
var errValueType = errors.New("value type error")

type heap struct {
	heapInterface
	heapList  []interface{}
	fn        comparable
	valueType reflect.Type
}

func newHeap(obj interface{}, fn comparable) *heap {
	var heap = &heap{
		heapList:  make([]interface{}, 1, 10),
		fn:        fn,
		valueType: reflect.TypeOf(obj),
	}
	heap.heapList[0] = nil
	return heap
}

func (m *heap) Top() (interface{}, error) {
	if len(m.heapList) <= 1 {
		return nil, errHeapEmpty
	}
	return m.heapList[1], nil
}

func (m *heap) Empty() bool {
	return len(m.heapList) == 1
}

func (m *heap) Size() int {
	return len(m.heapList) - 1
}

func (m *heap) Push(val interface{}) error {
	if reflect.TypeOf(val) != m.valueType {
		return fmt.Errorf("%w:want type %s, but get type %s", errValueType, m.valueType, reflect.TypeOf(val))
	}
	m.heapList = append(m.heapList, val)
	return m.heapifyUp(len(m.heapList) - 1)
}

func (m *heap) heapifyUp(i int) error {
	for i > 1 {
		var tmp, err = m.compare(m.heapList[i], m.heapList[i/2])
		if err != nil {
			return err
		}
		if tmp {
			var tmp = m.heapList[i]
			m.heapList[i] = m.heapList[i/2]
			m.heapList[i/2] = tmp
			i /= 2
		} else {
			break
		}
	}
	return nil
}

func (m *heap) Pop() (interface{}, error) {
	var v = m.heapList[1]
	m.heapList[1] = m.heapList[len(m.heapList)-1]
	m.heapList = m.heapList[0 : len(m.heapList)-1]
	err := m.heapifyDown(1)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *heap) heapifyDown(i int) error {
	var ch = i * 2
	for ch < len(m.heapList) {
		if ch+1 < len(m.heapList) {
			tmp, err := m.compare(m.heapList[ch+1], m.heapList[ch])
			if err != nil {
				return err
			}
			if tmp {
				ch++
			}
		}
		tmp, err := m.compare(m.heapList[ch], m.heapList[i])
		if err != nil {
			return err
		}
		if tmp {
			var tmp = m.heapList[i]
			m.heapList[i] = m.heapList[ch]
			m.heapList[ch] = tmp
		} else {
			break
		}
	}
	return nil
}

// func (m *minHeap) swap(val1, val2 heapVal) error {
// 	return nil
// }

func (m *heap) compare(val1, val2 interface{}) (bool, error) {
	if m.fn == nil {
		return false, errCompareableEmpty
	}
	return m.fn(val1, val2)
}
