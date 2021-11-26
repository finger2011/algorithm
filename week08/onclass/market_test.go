package week08onclass

import "testing"

func Test_market(t *testing.T) {
	type args struct {
		products [][]int
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
				products: [][]int{
					{50, 2},
					{10, 1},
					{20, 2},
					{30, 1},
				},
			},
			want: 80,
		},
		{
			name: "test02",
			args: args{
				products: [][]int{
					{20, 1},
					{2, 1},
					{10, 3},
					{100, 2},
					{8, 2},
					{5, 20},
					{50, 10},
				},
			},
			want: 185,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := market(tt.args.products); got != tt.want {
				t.Errorf("market() = %v, want %v", got, tt.want)
			}
		})
	}
}
