package week02onclass

import (
	"testing"
)

func Test_maxSubArray(t *testing.T) {
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
				nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			},
			want: 6,
		},
		{
			name: "test02",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "test03",
			args: args{
				nums: []int{0},
			},
			want: 0,
		},
		{
			name: "test04",
			args: args{
				nums: []int{-1},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxSubArray2(t *testing.T) {
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
				nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			},
			want: 6,
		},
		{
			name: "test02",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "test03",
			args: args{
				nums: []int{0},
			},
			want: 0,
		},
		{
			name: "test04",
			args: args{
				nums: []int{-1},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray2(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxSubArray3(t *testing.T) {
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
				nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			},
			want: 6,
		},
		{
			name: "test02",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "test03",
			args: args{
				nums: []int{0},
			},
			want: 0,
		},
		{
			name: "test04",
			args: args{
				nums: []int{-1},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray3(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray3() = %v, want %v", got, tt.want)
			}
		})
	}
}
