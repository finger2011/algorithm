package week01onclass

import "testing"

func Test_calculateWithParenthess(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			args: args{s: "(1+(4+5+2)-3)+(6+8)"},
			want: 23,
		},
		{
			name: "test02",
			args: args{s: "1+(4+5)*2"},
			want: 19,
		},
		{
			name: "test03",
			args: args{s: "-2+ 1"},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateWithParenthess(tt.args.s); got != tt.want {
				t.Errorf("calculateWithParenthess() = %v, want %v", got, tt.want)
			}
		})
	}
}
