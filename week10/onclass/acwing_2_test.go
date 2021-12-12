package week10onclass

import (
	"testing"
)

type arg struct {
	cal  string
	n1   int
	n2   int
	n3   int
	want int
}

func Test_SegmentTree(t *testing.T) {
	nums := []int{-79, -37, 39, 10, 57, -48, 77, -46, 72, -17, 11, 60, 41, -55, -6, -53, 56, -18, 53, -15, -18, -46, -20, -73, -32, 79, 47, -21, 29, 49, 76, -46, -16, -25, -27, 37, 59, 15, -64, -51}
	tree := newSegmentTree(nums)
	args := []arg{
		{cal: "Q", n1: 30, n2: 30, want: 49},
		{cal: "C", n1: 29, n2: 38, n3: -27},
		{cal: "C", n1: 5, n2: 40, n3: -33},
		{cal: "C", n1: 15, n2: 39, n3: 33},
		{cal: "Q", n1: 30, n2: 39, want: -185},
	}
	for _, arg := range args {
		if arg.cal == "Q" {
			ans := tree.getRank(1, arg.n1, arg.n2)
			if ans != arg.want {
				t.Errorf("getRank return %d, want:%d", ans, arg.want)
			}
		} else {
			tree.change(1, arg.n1, arg.n2, arg.n3)
		}
	}
}
