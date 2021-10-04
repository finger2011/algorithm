package week02

import (
	"testing"
)

func Test_findShortestSubArray(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{1, 2, 2, 3, 1},
			},
			want: 2,
		},
		{
			name: "test02",
			args: args{
				nums: []int{1, 2, 2, 3, 1, 4, 2},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findShortestSubArray(tt.args.nums); got != tt.want {
				t.Errorf("findShortestSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findShortestSubArray2(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{1, 2, 2, 3, 1},
			},
			want: 2,
		},
		{
			name: "test02",
			args: args{
				nums: []int{1, 2, 2, 3, 1, 4, 2},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findShortestSubArray2(tt.args.nums); got != tt.want {
				t.Errorf("findShortestSubArray2() = %v, want %v", got, tt.want)
			}
		})
	}
}
