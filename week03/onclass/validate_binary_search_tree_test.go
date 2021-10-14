package week03onclass

import (
	"finger2011/algggorithm/internel"
	"testing"
)

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *internel.TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			args: args{root: internel.CreateTreeByIntLevel([]int{2, 1, 3})},
			want: true,
		},
		{
			name: "test02",
			args: args{root: internel.CreateTreeByIntLevel([]int{5, 1, 4, internel.MinInt, internel.MinInt, 3, 6})},
			want: false,
		},
		{
			name: "test03",
			args: args{root: internel.CreateTreeByIntLevel([]int{2, 2, 2})},
			want: false,
		},
		{
			name: "test04",
			args: args{root: internel.CreateTreeByIntLevel([]int{0, -1})},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
