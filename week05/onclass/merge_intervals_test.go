package week05onclass

import (
	"reflect"
	"testing"
)

func Test_mergeIntervals(t *testing.T) {
	type args struct {
		intervals [][]int
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
				intervals: [][]int{
					{1, 3},
					{2, 6},
					{8, 10},
					{14, 18},
				},
			},
			want: [][]int{
				{1, 6},
				{8, 10},
				{14, 18},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeIntervals(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeIntervals2(t *testing.T) {
	type args struct {
		intervals [][]int
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
				intervals: [][]int{
					{1, 3},
					{2, 6},
					{8, 10},
					{14, 18},
				},
			},
			want: [][]int{
				{1, 6},
				{8, 10},
				{14, 18},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeIntervals2(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeIntervals2() = %v, want %v", got, tt.want)
			}
		})
	}
}
