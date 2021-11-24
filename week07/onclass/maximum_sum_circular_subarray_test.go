package week07onclass

import (
	"testing"
)

func Test_maxSubarraySumCircular(t *testing.T) {
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
				nums: []int{1, -2, 3, -2},
			},
			want: 3,
		},
		{
			name: "test02",
			args: args{
				nums: []int{5, -3, 5},
			},
			want: 10,
		},
		{
			name: "test01",
			args: args{
				nums: []int{3, -1, 2, -1},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubarraySumCircular(tt.args.nums); got != tt.want {
				t.Errorf("maxSubarraySumCircular() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxSubarraySumCircular2(t *testing.T) {
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
				nums: []int{1, -2, 3, -2},
			},
			want: 3,
		},
		{
			name: "test02",
			args: args{
				nums: []int{5, -3, 5},
			},
			want: 10,
		},
		{
			name: "test01",
			args: args{
				nums: []int{3, -1, 2, -1},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubarraySumCircular2(tt.args.nums); got != tt.want {
				t.Errorf("maxSubarraySumCircular2() = %v, want %v", got, tt.want)
			}
		})
	}
}
