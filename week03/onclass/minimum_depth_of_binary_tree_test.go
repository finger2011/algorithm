package week03onclass

import (
	"finger2011/algggorithm/internel"
	"testing"
)

func Test_minDepth(t *testing.T) {
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
			want: 2,
		},
		{
			name: "test02",
			args: args{
				root: internel.CreateTreeByIntLevel([]int{2, internel.MinInt, 3, internel.MinInt, 4, internel.MinInt, 5, internel.MinInt, 6}),
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDepth(tt.args.root); got != tt.want {
				t.Errorf("minDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minDepth2(t *testing.T) {
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
			want: 2,
		},
		{
			name: "test02",
			args: args{
				root: internel.CreateTreeByIntLevel([]int{2, internel.MinInt, 3, internel.MinInt, 4, internel.MinInt, 5, internel.MinInt, 6}),
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDepth2(tt.args.root); got != tt.want {
				t.Errorf("minDepth2() = %v, want %v", got, tt.want)
			}
		})
	}
}
