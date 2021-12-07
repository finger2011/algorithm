package week09onclass

import "testing"

func Test_eight(t *testing.T) {
	type args struct {
		start string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test02",
			args: args{start: "23415x768"},
			want: 19,
		},
		{
			name: "test01",
			args: args{start: "123x46758"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eight(tt.args.start); got != tt.want {
				t.Errorf("eight() = %v, want %v", got, tt.want)
			}
		})
	}
}
