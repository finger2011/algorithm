package week10onclass

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	nums := []int{1,3,5}
	obj := Constructor(nums)
	res := obj.SumRange(0, 2)
	if res != 9 {
		t.Errorf("SumRange want %d, got %d", 9, res)
	}
	obj.Update(1,2)
	res = obj.SumRange(0, 2)
	if res != 8 {
		t.Errorf("SumRange want %d, got %d", 8, res)
	}
}
