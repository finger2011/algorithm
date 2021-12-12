package week10

import (
	"reflect"
	"testing"
)

func Test_lookup(t *testing.T) {
	type args struct {
		k    int
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "test03",
			args: args{
				k:    10,
				nums: []int{4, 5, 6, 1, 2, 3, 7, 8, 9, 10},
			},
			want: [][]int{
				{1, 1},
				{1, 2},
				{3, 1},
				{1, 4},
				{1, 5},
				{1, 3},
				{1, 7},
				{1, 8},
				{1, 9},
			},
		},
		{
			name: "test01",
			args: args{
				k:    3,
				nums: []int{1, 5, 3},
			},
			want: [][]int{
				{4, 1},
				{2, 1},
			},
		},
		{
			name: "test02",
			args: args{
				k:    4,
				nums: []int{1, 5, 3, 4},
			},
			want: [][]int{
				{4, 1},
				{2, 1},
				{1, 3},
			},
		},
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lookup(tt.args.k, tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lookup() = %v, want %v", got, tt.want)
			}
		})
	}
}
