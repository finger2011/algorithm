package week09

import "testing"

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			args: args{
				s: []byte("hello"),
			},
			want: []byte("olleh"),
		},
		{
			name: "test02",
			args: args{
				s: []byte("Hannah"),
			},
			want: []byte("hannaH"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.args.s)
			if string(tt.want) != string(tt.args.s) {
				t.Errorf("reverseString() = %v, want %v", tt.args.s, tt.want)
			}
		})
	}
}
