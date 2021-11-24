package week07onclass

import "testing"

func Test_findMaxValueOfEquation(t *testing.T) {
	type args struct {
		points [][]int
		k      int
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
				points: [][]int{{1, 3}, {2, 0}, {5, 10}, {6, -10}},
				k:      1,
			},
			want: 4,
		},
		{
			name: "test02",
			args: args{
				points: [][]int{{0, 0}, {3, 0}, {9, 2}},
				k:      3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxValueOfEquation(tt.args.points, tt.args.k); got != tt.want {
				t.Errorf("findMaxValueOfEquation() = %v, want %v", got, tt.want)
			}
		})
	}
}
