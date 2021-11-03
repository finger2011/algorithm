package week05onclass

import (
	"reflect"
	"testing"
)

func Test_selectionSort(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{5, 2, 3, 1},
			},
			want: []int{1, 2, 3, 5},
		},
		{
			name: "test02",
			args: args{
				nums: []int{5, 1, 1, 2, 0, 0},
			},
			want: []int{0, 0, 1, 1, 2, 5},
		},
		{
			name: "test03",
			args: args{
				nums: []int{-4, 0, 7, 4, 9, -5, -1, 0, -7, -1},
			},
			want: []int{-7, -5, -4, -1, -1, 0, 0, 4, 7, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := selectionSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("selectionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_heapSort(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{5, 2, 3, 1},
			},
			want: []int{1, 2, 3, 5},
		},
		{
			name: "test02",
			args: args{
				nums: []int{5, 1, 1, 2, 0, 0},
			},
			want: []int{0, 0, 1, 1, 2, 5},
		},
		{
			name: "test03",
			args: args{
				nums: []int{-4, 0, 7, 4, 9, -5, -1, 0, -7, -1},
			},
			want: []int{-7, -5, -4, -1, -1, 0, 0, 4, 7, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := heapSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("heapSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortArray(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{5, 2, 3, 1},
			},
			want: []int{1, 2, 3, 5},
		},
		{
			name: "test02",
			args: args{
				nums: []int{5, 1, 1, 2, 0, 0},
			},
			want: []int{0, 0, 1, 1, 2, 5},
		},
		{
			name: "test03",
			args: args{
				nums: []int{-4, 0, 7, 4, 9, -5, -1, 0, -7, -1},
			},
			want: []int{-7, -5, -4, -1, -1, 0, 0, 4, 7, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_insertionSort(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{5, 2, 3, 1},
			},
			want: []int{1, 2, 3, 5},
		},
		{
			name: "test02",
			args: args{
				nums: []int{5, 1, 1, 2, 0, 0},
			},
			want: []int{0, 0, 1, 1, 2, 5},
		},
		{
			name: "test03",
			args: args{
				nums: []int{-4, 0, 7, 4, 9, -5, -1, 0, -7, -1},
			},
			want: []int{-7, -5, -4, -1, -1, 0, 0, 4, 7, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertionSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bubbleSort(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{5, 2, 3, 1},
			},
			want: []int{1, 2, 3, 5},
		},
		{
			name: "test02",
			args: args{
				nums: []int{5, 1, 1, 2, 0, 0},
			},
			want: []int{0, 0, 1, 1, 2, 5},
		},
		{
			name: "test03",
			args: args{
				nums: []int{-4, 0, 7, 4, 9, -5, -1, 0, -7, -1},
			},
			want: []int{-7, -5, -4, -1, -1, 0, 0, 4, 7, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bubbleSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
