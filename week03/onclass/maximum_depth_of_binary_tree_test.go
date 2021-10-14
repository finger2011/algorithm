package week03onclass

import (
	"finger2011/algggorithm/internel"
	"testing"
)

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *internel.TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			args: args{
				root: internel.CreateTreeByIntLevel([]int{3, 9, 20, internel.MinInt, internel.MinInt, 15, 7}),
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxDepth2(t *testing.T) {
	type args struct {
		root *internel.TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			args: args{
				root: internel.CreateTreeByIntLevel([]int{3, 9, 20, internel.MinInt, internel.MinInt, 15, 7}),
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth2(tt.args.root); got != tt.want {
				t.Errorf("maxDepth2() = %v, want %v", got, tt.want)
			}
		})
	}
}
