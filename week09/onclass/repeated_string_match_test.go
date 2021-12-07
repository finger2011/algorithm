package week09onclass

import "testing"

func Test_repeatedStringMatch(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test02",
			args: args{
				a: "abc",
				b: "cabcabca",
			},
			want: 4,
		},
		{
			name: "test01",
			args: args{
				a: "abcd",
				b: "cdabcdab",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedStringMatch(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("repeatedStringMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
