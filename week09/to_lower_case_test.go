package week09

import "testing"

func Test_toLowerCase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			args: args{
				s: "Hello",
			},
			want: "hello",
		},
		{
			name: "test02",
			args: args{
				s: "here",
			},
			want: "here",
		},
		{
			name: "test03",
			args: args{
				s: "LOVELY",
			},
			want: "lovely",
		},
		{
			name: "test04",
			args: args{
				s: "al&phaBET",
			},
			want: "al&phabet",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toLowerCase(tt.args.s); got != tt.want {
				t.Errorf("toLowerCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
