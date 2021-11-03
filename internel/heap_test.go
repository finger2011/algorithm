package internel

import (
	"errors"
	"testing"
)

type minHeap struct {
	heap
}

type heapVal struct {
	val int
}

// func (v *heapValTest) value() interface{} {
// 	return v.val
// }

func minInArr(nums []int) (int, []int) {
	var min = []int{nums[0], 0}
	for i := 1; i < len(nums); i++ {
		if nums[i] < min[0] {
			min = []int{nums[i], i}
		}
	}
	var tmp = nums[0:min[1]]
	if min[1] < len(nums)-1 {
		nums = append(append([]int{}, tmp...), nums[min[1]+1:]...)
	} else {
		nums = tmp
	}
	return min[0], nums
}

func maxInArr(nums []int) (int, []int) {
	var min = []int{nums[0], 0}
	for i := 1; i < len(nums); i++ {
		if nums[i] > min[0] {
			min = []int{nums[i], i}
		}
	}
	var tmp = nums[0:min[1]]
	if min[1] < len(nums)-1 {
		nums = append(append([]int{}, tmp...), nums[min[1]+1:]...)
	} else {
		nums = tmp
	}
	return min[0], nums
}

func comparableTest(val1, val2 interface{}) (bool, error) {
	v1, ok := val1.(*heapVal)
	if !ok {
		return false, errConvertFailed
	}
	v2, ok := val2.(*heapVal)
	if !ok {
		return false, errConvertFailed
	}
	if v1.val > v2.val {
		return true, nil
	}
	return false, nil
}

var errConvertFailed = errors.New("convert value type failedd")

func comparableNormal(val1, val2 interface{}) (bool, error) {
	v1, ok := val1.(*heapVal)
	if !ok {
		return false, errConvertFailed
	}
	v2, ok := val2.(*heapVal)
	if !ok {
		return false, errConvertFailed
	}
	if v1.val < v2.val {
		return true, nil
	}
	return false, nil
}

func Test_heap(t *testing.T) {
	var m = newHeap(&heapVal{}, comparableNormal)
	// var err error
	var input = []int{1, 2, 4, 5, 6, 9}
	var min int
	for i := 0; i < len(input); i++ {
		err := m.Push(&heapVal{val: input[i]})
		if err != nil {
			t.Errorf("heap Push() = %v, want nil", err)
		}
	}

	val, err := m.Pop()
	if err != nil {
		t.Errorf("heap Pop() = %v, want nil", err)
	}
	min, input = minInArr(input)
	if val.(*heapVal).val != min {
		t.Errorf("heap Pop() = %v, want %d", val, min)
	}

	val, err = m.Pop()
	if err != nil {
		t.Errorf("heap Pop() = %v, want nil", err)
	}
	min, input = minInArr(input)
	if val.(*heapVal).val != min {
		t.Errorf("heap Pop() = %v, want %d", val, min)
	}

	val, err = m.Top()
	min, input = minInArr(input)
	if err != nil {
		t.Errorf("heap Top() = %v, want nil", err)
	}
	if val.(*heapVal).val != min {
		t.Errorf("heap Pop() = %v, want %d", val, min)
	}

}

func Test_heapWithFunc(t *testing.T) {
	var m = newHeap(&heapVal{}, comparableTest)
	var err error

	var input = []int{1, 2, 4, 5, 6, 9}
	var min int
	for i := 0; i < len(input); i++ {
		err = m.Push(&heapVal{val: input[i]})
		if err != nil {
			t.Errorf("heap Push() = %v, want nil", err)
		}
	}

	val, err := m.Pop()
	if err != nil {
		t.Errorf("heap Pop() = %v, want nil", err)
	}
	min, input = maxInArr(input)
	if val.(*heapVal).val != min {
		t.Errorf("heap Pop() = %v, want %d", val, min)
	}

	val, err = m.Pop()
	if err != nil {
		t.Errorf("heap Pop() = %v, want nil", err)
	}
	min, input = maxInArr(input)
	if val.(*heapVal).val != min {
		t.Errorf("heap Pop() = %v, want %d", val, min)
	}

	val, err = m.Top()
	min, input = maxInArr(input)
	if err != nil {
		t.Errorf("heap Top() = %v, want nil", err)
	}
	if val.(*heapVal).val != min {
		t.Errorf("heap Pop() = %v, want %d", val, min)
	}

}
