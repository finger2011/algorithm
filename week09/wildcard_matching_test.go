package week09

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test04",
			args: args{
				s: "abcde",
				p: "*a*b",
			},
			want: false,
		},
		{
			name: "test03",
			args: args{
				s: "aab",
				p: "c*",
			},
			want: false,
		},
		{
			name: "test01",
			args: args{
				s: "aa",
				p: "a",
			},
			want: false,
		},
		{
			name: "test02",
			args: args{
				s: "aa",
				p: "*",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
