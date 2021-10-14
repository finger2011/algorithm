package week03onclass

import "testing"

func Test_myPow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			args: args{
				x: 2.00000,
				n: 10,
			},
			want: 1024.00000,
		},
		{
			name: "test02",
			args: args{
				x: 2.10000,
				n: 3,
			},
			want: 9.26100,
		},
		{
			name: "test03",
			args: args{
				x: 2.00000,
				n: -2,
			},
			want: 0.25000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myPow(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("myPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
