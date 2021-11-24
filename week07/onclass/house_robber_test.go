package week07onclass

import (
	"testing"
)

func Test_rob1(t *testing.T) {
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
				nums: []int{1, 2, 3, 1},
			},
			want: 4,
		},
		{
			name: "test02",
			args: args{
				nums: []int{2, 7, 9, 3, 1},
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob1(tt.args.nums); got != tt.want {
				t.Errorf("rob1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rob1_2(t *testing.T) {
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
				nums: []int{1, 2, 3, 1},
			},
			want: 4,
		},
		{
			name: "test02",
			args: args{
				nums: []int{2, 7, 9, 3, 1},
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob1_2(tt.args.nums); got != tt.want {
				t.Errorf("rob1_2() = %v, want %v", got, tt.want)
			}
		})
	}
}
