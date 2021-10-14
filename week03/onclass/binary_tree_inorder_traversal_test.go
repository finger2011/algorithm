package week03onclass

import (
	"finger2011/algggorithm/internel"
	"reflect"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	type args struct {
		root *internel.TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			args: args{root: internel.CreateTreeByIntLevel([]int{})},
			want: []int{},
		},
		{
			name: "test02",
			args: args{root: internel.CreateTreeByIntLevel([]int{1, internel.MinInt, 2, 3})},
			want: []int{1, 3, 2},
		},
		{
			name: "test03",
			args: args{root: internel.CreateTreeByIntLevel([]int{1})},
			want: []int{1},
		},
		{
			name: "test04",
			args: args{root: internel.CreateTreeByIntLevel([]int{1, 2})},
			want: []int{2, 1},
		},
		{
			name: "test05",
			args: args{root: internel.CreateTreeByIntLevel([]int{1, internel.MinInt, 2})},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
