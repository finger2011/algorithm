package internel

import (
	"testing"
)

func Test_createTree(t *testing.T) {
	var nums = []int{4, 2, 1, 3, 7, 6, 9}
	var root = createTree(nums)
	if !IntEqual(root.TreeToInt(), nums) {
		t.Errorf("createTree() = %v, want %v", root.TreeToInt(), nums)
	}
}

func Test_createTreeByLevel(t *testing.T) {
	var nums = []int{4, 2, 7, 1, 3, 6, 9}
	var want = []int{4, 2, 1, 3, 7, 6, 9}
	var root = createTreeByLevel(nums)
	if !IntEqual(root.TreeToInt(), want) {
		t.Errorf("createTree() = %v, want %v", root.TreeToInt(), want)
	}
}
