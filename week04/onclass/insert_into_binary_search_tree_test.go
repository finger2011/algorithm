package week04onclass

import (
	"finger2011/algggorithm/internel"
	"testing"
)

func Test_insertIntoBST(t *testing.T) {
	type args struct {
		root *internel.TreeNode
		val  int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			args: args{
				root: internel.CreateTreeByIntLevel([]int{4,2,7,1,3}),
				val:  5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertIntoBST(tt.args.root, tt.args.val); !got.ValidBST() || !got.ContainVal(tt.args.val) {
				t.Errorf("insertIntoBST() not a BST  = %v or not contain val = %d", got.TreeToInt(), tt.args.val)
			}
		})
	}
}
