package week03onclass

import (
	"finger2011/algggorithm/internel"
	"testing"
)

func Test_subsets(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {1, 2, 3}, {2, 3}},
		},
		{
			name: "test02",
			args: args{
				nums: []int{1, 2},
			},
			want: [][]int{{}, {1}, {2}, {1, 2}},
		},
		{
			name: "test03",
			args: args{
				nums: []int{0},
			},
			want: [][]int{{}, {0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsets(tt.args.nums); !internel.IntsEqual(got, tt.want) {
				t.Errorf("subsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
