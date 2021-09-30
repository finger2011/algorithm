package week01onclass

import (
	"reflect"
	"testing"
)

func Test_maxSlidingWindow(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			args: args{
				nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
				k:    3,
			},
			want: []int{3, 3, 5, 5, 6, 7},
		},
		{
			name: "test02",
			args: args{
				nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
				k:    1,
			},
			want: []int{1, 3, -1, -3, 5, 3, 6, 7},
		},
		{
			name: "test03",
			args: args{
				nums: []int{9, 11},
				k:    2,
			},
			want: []int{11},
		},
		{
			name: "test04",
			args: args{
				nums: []int{4, -2},
				k:    2,
			},
			want: []int{4},
		},
		{
			name: "test05",
			args: args{
				nums: []int{7, 2, 4},
				k:    2,
			},
			want: []int{7, 4},
		},
		//[-7,-8,7,5,7,1,6,0] 4
		{
			name: "test06",
			args: args{
				nums: []int{-7, -8, 7, 5, 7, 1, 6, 0},
				k:    4,
			},
			want: []int{7, 7, 7, 7, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSlidingWindow(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxSlidingWindow2(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			args: args{
				nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
				k:    3,
			},
			want: []int{3, 3, 5, 5, 6, 7},
		},
		{
			name: "test02",
			args: args{
				nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
				k:    1,
			},
			want: []int{1, 3, -1, -3, 5, 3, 6, 7},
		},
		{
			name: "test03",
			args: args{
				nums: []int{9, 11},
				k:    2,
			},
			want: []int{11},
		},
		{
			name: "test04",
			args: args{
				nums: []int{4, -2},
				k:    2,
			},
			want: []int{4},
		},
		{
			name: "test05",
			args: args{
				nums: []int{7, 2, 4},
				k:    2,
			},
			want: []int{7, 4},
		},
		//[-7,-8,7,5,7,1,6,0] 4
		{
			name: "test06",
			args: args{
				nums: []int{-7, -8, 7, 5, 7, 1, 6, 0},
				k:    4,
			},
			want: []int{7, 7, 7, 7, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSlidingWindow2(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSlidingWindow2() = %v, want %v", got, tt.want)
			}
		})
	}
}
