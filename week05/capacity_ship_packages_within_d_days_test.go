package week05

import "testing"

func Test_shipWithinDays(t *testing.T) {
	type args struct {
		weights []int
		days    int
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
				weights: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				days:    5,
			},
			want: 15,
		},
		{
			name: "test02",
			args: args{
				weights: []int{3, 2, 2, 4, 1, 4},
				days:    3,
			},
			want: 6,
		},
		{
			name: "test03",
			args: args{
				weights: []int{1, 2, 3, 1, 1},
				days:    4,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shipWithinDays(tt.args.weights, tt.args.days); got != tt.want {
				t.Errorf("shipWithinDays() = %v, want %v", got, tt.want)
			}
		})
	}
}
