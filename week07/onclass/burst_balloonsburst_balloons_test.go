package week07onclass

import (
	"testing"
)

func Test_maxCoins(t *testing.T) {
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
				nums: []int{3, 1, 5, 8},
			},
			want: 167,
		},
		{
			name: "test02",
			args: args{
				nums: []int{1, 5},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCoins(tt.args.nums); got != tt.want {
				t.Errorf("maxCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxCoins2(t *testing.T) {
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
				nums: []int{3, 1, 5, 8},
			},
			want: 167,
		},
		{
			name: "test02",
			args: args{
				nums: []int{1, 5},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCoins2(tt.args.nums); got != tt.want {
				t.Errorf("maxCoins2() = %v, want %v", got, tt.want)
			}
		})
	}
}
