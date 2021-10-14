package week03

import (
	"reflect"
	"testing"
)

func Test_findRedundantDirectedConnection(t *testing.T) {
	type args struct {
		edges [][]int
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
				edges: [][]int{{1, 2}, {1, 3}, {2, 3}},
			},
			want: []int{2, 3},
		},
		{
			name: "test02",
			args: args{
				// [[1,2],[2,3],[3,4],[4,1],[1,5]]
				edges: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 5}},
			},
			want: []int{4, 1},
		},
		{
			name: "test03",
			args: args{
				//[[2,1],[3,1],[4,2],[1,4]]
				edges: [][]int{{2, 1}, {3, 1}, {4, 2}, {1, 4}},
			},
			want: []int{2, 1},
		},
		{
			name: "test04",
			args: args{
				//[[4,2],[1,5],[5,2],[5,3],[2,4]]
				edges: [][]int{{4, 2}, {1, 5}, {5, 2}, {5, 3}, {2, 4}},
			},
			want: []int{4, 2},
		},
		{
			name: "test05",
			args: args{
				//[[1,2],[2,3],[3,1],[4,1]]
				edges: [][]int{{1, 2}, {2, 3}, {3, 1},{4,1}},
			},
			want: []int{3, 1},
		},
		{
			name: "test06",
			args: args{
				//[[5,2],[5,1],[3,1],[3,4],[3,5]]
				edges: [][]int{{5, 2}, {5, 1}, {3, 1}, {3, 4}, {3, 5}},
			},
			want: []int{3, 1},
		},
		{
			name: "test07",
			args: args{
				//[[2,3],[3,1],[3,4],[4,2]]
				edges: [][]int{{2, 3}, {3, 1}, {3, 4}, {4, 2}},
			},
			want: []int{4, 2},
		},
		{
			name: "test08",
			args: args{
				//[[1,2],[2,1],[2,3],[3,4]]
				edges: [][]int{{1, 2}, {2, 1}, {2, 3}, {3, 4}},
			},
			want: []int{2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRedundantDirectedConnection(tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRedundantDirectedConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}
