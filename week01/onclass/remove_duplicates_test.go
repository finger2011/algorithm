package week01onclass

import (
	"reflect"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 int
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			args: args{
				nums: []int{1, 1, 2},
			},
			want:  []int{1, 2},
			want1: 2,
		},
		{
			name: "test02",
			args: args{
				nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			},
			want:  []int{0, 1, 2, 3, 4},
			want1: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := removeDuplicates(tt.args.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeDuplicates() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("removeDuplicates() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_removeDuplicatesByK(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 int
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			args: args{
				nums: []int{1, 1, 2},
				k:    3,
			},
			want:  []int{1, 1, 2},
			want1: 3,
		},
		{
			name: "test02",
			args: args{
				nums: []int{1, 1, 1, 2, 2, 3},
				k:    2,
			},
			want:  []int{1, 1, 2, 2, 3},
			want1: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := removeDuplicatesByK(tt.args.nums, tt.args.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeDuplicatesByK() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("removeDuplicatesByK() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
