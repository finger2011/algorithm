package week03onclass

import (
	"finger2011/algggorithm/internel"
	"reflect"
	"testing"
)

func Test_invertTree(t *testing.T) {
	type args struct {
		root *internel.TreeNode
	}
	tests := []struct {
		name string
		args args
		want *internel.TreeNode
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			args: args{
				root: internel.CreateTreeByIntLevel([]int{4, 2, 7, 1, 3, 6, 9}),
			},
			want: internel.CreateTreeByIntLevel([]int{4, 7, 2, 9, 6, 3, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invertTree(tt.args.root); !reflect.DeepEqual(got.TreeToInt(), tt.want.TreeToInt()) {
				t.Errorf("invertTree() = %v, want %v", got.TreeToInt(), tt.want.TreeToInt())
			}
		})
	}
}

func Test_invertTree2(t *testing.T) {
	type args struct {
		root *internel.TreeNode
	}
	tests := []struct {
		name string
		args args
		want *internel.TreeNode
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			args: args{
				root: internel.CreateTreeByIntLevel([]int{4, 2, 7, 1, 3, 6, 9}),
			},
			want: internel.CreateTreeByIntLevel([]int{4, 7, 2, 9, 6, 3, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invertTree2(tt.args.root); !reflect.DeepEqual(got.TreeToInt(), tt.want.TreeToInt()) {
				t.Errorf("invertTree2() = %v, want %v", got.TreeToInt(), tt.want.TreeToInt())
			}
		})
	}
}
