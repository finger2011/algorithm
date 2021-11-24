package week07onclass

import (
	"testing"
)

func Test_maxProfit6(t *testing.T) {
	type args struct {
		prices []int
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
				prices: []int{1, 2, 3, 0, 2},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit6(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit6() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProfit6_2(t *testing.T) {
	type args struct {
		prices []int
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
				prices: []int{1, 2, 3, 0, 2},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit6_2(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit6_2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProfit6_3(t *testing.T) {
	type args struct {
		prices []int
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
				prices: []int{1, 2, 3, 0, 2},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit6_3(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit6_3() = %v, want %v", got, tt.want)
			}
		})
	}
}
